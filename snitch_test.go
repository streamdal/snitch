package snitch

import (
	"context"
	"net"
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"google.golang.org/grpc"

	"github.com/streamdal/snitch-protos/build/go/protos"
	"github.com/streamdal/snitch-protos/build/go/protos/steps"

	"github.com/streamdal/snitch-go-client/logger"
	"github.com/streamdal/snitch-go-client/logger/loggerfakes"
	"github.com/streamdal/snitch-go-client/metrics/metricsfakes"
	"github.com/streamdal/snitch-go-client/server/serverfakes"
)

func TestValidateConfig(t *testing.T) {
	t.Run("invalid config", func(t *testing.T) {
		err := validateConfig(nil)
		if !errors.Is(err, ErrEmptyConfig) {
			t.Error("expected error but got nil")
		}
	})

	t.Run("empty data source", func(t *testing.T) {
		cfg := &Config{
			ServiceName: "",
			ShutdownCtx: context.Background(),
			SnitchURL:   "http://localhost:9090",
			SnitchToken: "foo",
			DryRun:      false,
			StepTimeout: time.Second,
			Logger:      &logger.NoOpLogger{},
		}
		err := validateConfig(cfg)
		if !errors.Is(err, ErrEmptyServiceName) {
			t.Error("expected ErrEmptyServiceName")
		}
	})

	t.Run("empty context", func(t *testing.T) {
		cfg := &Config{
			ServiceName: "mysvc1",
			ShutdownCtx: nil,
			SnitchURL:   "http://localhost:9090",
			SnitchToken: "foo",
			DryRun:      false,
			StepTimeout: time.Second,
			Logger:      &logger.NoOpLogger{},
		}
		err := validateConfig(cfg)
		if !errors.Is(err, ErrMissingShutdownCtx) {
			t.Error("expected ErrMissingShutdownCtx")
		}
	})

	t.Run("invalid wasm timeout duration", func(t *testing.T) {
		_ = os.Setenv("SNITCH_STEP_TIMEOUT", "foo")
		cfg := &Config{
			ServiceName: "mysvc1",
			ShutdownCtx: context.Background(),
			SnitchURL:   "http://localhost:9090",
			SnitchToken: "foo",
			DryRun:      false,
			Logger:      &logger.NoOpLogger{},
		}
		err := validateConfig(cfg)
		if err == nil || !strings.Contains(err.Error(), "unable to parse SNITCH_STEP_TIMEOUT") {
			t.Error("expected time.ParseDuration error")
		}
		_ = os.Unsetenv("SNITCH_STEP_TIMEOUT")
	})
}

type InternalServer struct {
	// Must be implemented in order to satisfy the protos InternalServer interface
	protos.UnimplementedInternalServer
}

func (i *InternalServer) GetAttachCommandsByService(ctx context.Context, req *protos.GetAttachCommandsByServiceRequest) (*protos.GetAttachCommandsByServiceResponse, error) {
	return &protos.GetAttachCommandsByServiceResponse{}, nil
}

func TestNew(t *testing.T) {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		t.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	protos.RegisterInternalServer(srv, &InternalServer{})

	go func() {
		if err := srv.Serve(lis); err != nil {
			panic("failed to serve: " + err.Error())
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := &Config{
		ServiceName: "mysvc1",
		ShutdownCtx: ctx,
		SnitchURL:   "localhost:9090",
		SnitchToken: "foo",
		DryRun:      false,
		Logger:      &loggerfakes.FakeLogger{},
	}

	if _, err := New(cfg); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestGetPipelines(t *testing.T) {
	ctx := context.Background()

	fakeClient := &serverfakes.FakeIServerClient{}

	s := &Snitch{
		pipelinesMtx: &sync.RWMutex{},
		pipelines:    map[string]map[string]*protos.Command{},
		serverClient: fakeClient,
		audiencesMtx: &sync.RWMutex{},
		audiences:    map[string]struct{}{},
	}

	aud := &protos.Audience{
		ServiceName:   "mysvc1",
		ComponentName: "kafka",
		OperationType: protos.OperationType_OPERATION_TYPE_PRODUCER,
		OperationName: "mytopic",
	}

	t.Run("no pipelines", func(t *testing.T) {
		pipelines := s.getPipelines(ctx, aud)
		if len(pipelines) != 0 {
			t.Error("expected empty map")
		}

		// Allow time for goroutine to run
		time.Sleep(time.Millisecond * 500)

		// Audience should be created on the server
		if fakeClient.NewAudienceCallCount() != 1 {
			t.Error("expected NewAudience to be called")
		}
	})

	t.Run("single pipeline", func(t *testing.T) {
		s.pipelines[audToStr(aud)] = map[string]*protos.Command{
			uuid.New().String(): {},
		}

		if len(s.getPipelines(ctx, aud)) != 1 {
			t.Error("expected 1 pipeline")
		}
	})
}

func TestHandleConditions(t *testing.T) {
	fakeClient := &serverfakes.FakeIServerClient{}

	s := &Snitch{
		serverClient: fakeClient,
		metrics:      &metricsfakes.FakeIMetrics{},
		config: &Config{
			Logger: &loggerfakes.FakeLogger{},
			DryRun: false,
		},
	}

	aud := &protos.Audience{}
	pipeline := &protos.Pipeline{}
	step := &protos.PipelineStep{}
	req := &ProcessRequest{}

	t.Run("notify condition", func(t *testing.T) {
		conditions := []protos.PipelineStepCondition{protos.PipelineStepCondition_PIPELINE_STEP_CONDITION_NOTIFY}

		got := s.handleConditions(context.Background(), conditions, pipeline, step, aud, req)
		if got != true {
			t.Error("handleConditions() should return true")
		}
		if fakeClient.NotifyCallCount() != 1 {
			t.Error("expected Notify() to be called")
		}
	})

	t.Run("abort condition", func(t *testing.T) {
		conditions := []protos.PipelineStepCondition{protos.PipelineStepCondition_PIPELINE_STEP_CONDITION_ABORT}

		got := s.handleConditions(context.Background(), conditions, pipeline, step, aud, req)
		if got != false {
			t.Error("handleConditions() should return false")
		}
	})

}

func TestProcess_nil(t *testing.T) {
	s := &Snitch{}
	_, err := s.Process(context.Background(), nil)
	if err == nil || !strings.Contains(err.Error(), "request cannot be nil") {
		t.Error("expected error")
	}
}

func TestProcess_success(t *testing.T) {
	aud := &protos.Audience{
		ServiceName:   "mysvc1",
		ComponentName: "kafka",
		OperationType: protos.OperationType_OPERATION_TYPE_PRODUCER,
		OperationName: "mytopic",
	}

	wasmData, err := os.ReadFile("src/detective.wasm")
	if err != nil {
		t.Fatal(err)
	}

	pipeline := &protos.Pipeline{
		Id:   uuid.New().String(),
		Name: "Test Pipeline",
		Steps: []*protos.PipelineStep{
			{
				Name:          "Step 1",
				XWasmId:       stringPtr(uuid.New().String()),
				XWasmBytes:    wasmData,
				XWasmFunction: stringPtr("f"),
				OnSuccess:     make([]protos.PipelineStepCondition, 0),
				OnFailure:     []protos.PipelineStepCondition{protos.PipelineStepCondition_PIPELINE_STEP_CONDITION_ABORT},
				Step: &protos.PipelineStep_Detective{
					Detective: &steps.DetectiveStep{
						Path:   stringPtr("object.payload"),
						Args:   []string{"gmail.com"},
						Negate: boolPtr(false),
						Type:   steps.DetectiveType_DETECTIVE_TYPE_STRING_CONTAINS_ANY,
					},
				},
			},
		},
	}

	s := &Snitch{
		serverClient: &serverfakes.FakeIServerClient{},
		functionsMtx: &sync.RWMutex{},
		functions:    map[string]*function{},
		audiencesMtx: &sync.RWMutex{},
		audiences:    map[string]struct{}{},
		config: &Config{
			ServiceName:     "mysvc1",
			Logger:          &logger.NoOpLogger{},
			StepTimeout:     time.Millisecond * 10,
			PipelineTimeout: time.Millisecond * 100,
		},
		metrics:      &metricsfakes.FakeIMetrics{},
		pipelinesMtx: &sync.RWMutex{},
		pipelines: map[string]map[string]*protos.Command{
			audToStr(aud): {
				pipeline.Id: {
					Audience: aud,
					Command: &protos.Command_AttachPipeline{
						AttachPipeline: &protos.AttachPipelineCommand{
							Pipeline: pipeline,
						},
					},
				},
			},
		},
	}

	resp, err := s.Process(context.Background(), &ProcessRequest{
		ComponentName: aud.ComponentName,
		OperationType: OperationType(aud.OperationType),
		OperationName: aud.OperationName,
		Data:          []byte(`{"object":{"payload":"streamdal@gmail.com"}`),
	})
	if err != nil {
		t.Fatal(err)
	}

	if resp.Error {
		t.Error(resp.Message)
	}

	if resp.Message == "No pipelines, message ignored" {
		t.Error("no pipelines, message ignored")
	}
}

func TestProcess_matchfail_and_abort(t *testing.T) {
	aud := &protos.Audience{
		ServiceName:   "mysvc1",
		ComponentName: "kafka",
		OperationType: protos.OperationType_OPERATION_TYPE_PRODUCER,
		OperationName: "mytopic",
	}

	wasmData, err := os.ReadFile("src/detective.wasm")
	if err != nil {
		t.Fatal(err)
	}

	pipeline := &protos.Pipeline{
		Id:   uuid.New().String(),
		Name: "Test Pipeline",
		Steps: []*protos.PipelineStep{
			{
				Name:          "Step 1",
				XWasmId:       stringPtr(uuid.New().String()),
				XWasmBytes:    wasmData,
				XWasmFunction: stringPtr("f"),
				OnSuccess:     make([]protos.PipelineStepCondition, 0),
				OnFailure:     []protos.PipelineStepCondition{protos.PipelineStepCondition_PIPELINE_STEP_CONDITION_ABORT},
				Step: &protos.PipelineStep_Detective{
					Detective: &steps.DetectiveStep{
						Path:   stringPtr("object.payload"),
						Args:   []string{"gmail.com"},
						Negate: boolPtr(false),
						Type:   steps.DetectiveType_DETECTIVE_TYPE_STRING_CONTAINS_ANY,
					},
				},
			},
		},
	}

	s := &Snitch{
		serverClient: &serverfakes.FakeIServerClient{},
		functionsMtx: &sync.RWMutex{},
		functions:    map[string]*function{},
		audiencesMtx: &sync.RWMutex{},
		audiences:    map[string]struct{}{},
		config: &Config{
			ServiceName:     "mysvc1",
			Logger:          &logger.NoOpLogger{},
			StepTimeout:     time.Millisecond * 10,
			PipelineTimeout: time.Millisecond * 100,
		},
		metrics:      &metricsfakes.FakeIMetrics{},
		pipelinesMtx: &sync.RWMutex{},
		pipelines: map[string]map[string]*protos.Command{
			audToStr(aud): {
				pipeline.Id: {
					Audience: aud,
					Command: &protos.Command_AttachPipeline{
						AttachPipeline: &protos.AttachPipelineCommand{
							Pipeline: pipeline,
						},
					},
				},
			},
		},
	}

	resp, err := s.Process(context.Background(), &ProcessRequest{
		ComponentName: aud.ComponentName,
		OperationType: OperationType(aud.OperationType),
		OperationName: aud.OperationName,
		Data:          []byte(`{"object":{"payload":"streamdal@hotmail.com"}`),
	})
	if err != nil {
		t.Fatal(err)
	}

	if !resp.Error {
		t.Error("expected ProcessResponse.Error = true")
	}

	if resp.Message != "detective step failed" {
		t.Error("Expected ProcessResponse.Message = 'detective step failed'")
	}
}

func stringPtr(in string) *string {
	return &in
}

func boolPtr(in bool) *bool {
	return &in
}

func TestHttpRequest(t *testing.T) {
	wasmData, err := os.ReadFile("src/httprequest.wasm")
	if err != nil {
		t.Fatal(err)
	}

	req := &protos.WASMRequest{
		Step: &protos.PipelineStep{
			Step: &protos.PipelineStep_HttpRequest{
				HttpRequest: &steps.HttpRequestStep{
					Request: &steps.HttpRequest{
						Url:    "https://www.google.com",
						Method: steps.HttpRequestMethod_HTTP_REQUEST_METHOD_GET,
						Body:   []byte(``),
					},
				},
			},
			XWasmId:       stringPtr(uuid.New().String()),
			XWasmFunction: stringPtr("f"),
			XWasmBytes:    wasmData,
		},
		Input: []byte(``),
	}

	f, err := createFunction(req.Step)
	if err != nil {
		t.Fatal(err)
	}

	req.Step.XWasmBytes = nil

	data, err := proto.Marshal(req)
	if err != nil {
		t.Fatalf("Unable to marshal WASMRequest: %s", err)
	}

	res, err := f.Exec(context.Background(), data)
	if err != nil {
		t.Fatal(err)
	}

	wasmResp := &protos.WASMResponse{}

	if err := proto.Unmarshal(res, wasmResp); err != nil {
		t.Fatal("unable to unmarshal wasm response: " + err.Error())
	}

	if wasmResp.ExitCode != protos.WASMExitCode_WASM_EXIT_CODE_SUCCESS {
		t.Errorf("expected ExitCode = 0, got = %d", wasmResp.ExitCode)
	}
}

//func BenchmarkMatchSmallJSON(b *testing.B) {
//	matchBench("json-examples/small.json", b)
//}
//
//func BenchmarkMatchMediumJSON(b *testing.B) {
//	matchBench("json-examples/medium.json", b)
//}
//
//func BenchmarkMatchLargeJSON(b *testing.B) {
//	matchBench("json-examples/large.json", b)
//}
//
//func BenchmarkTransformSmallJSON(b *testing.B) {
//	transformBench("json-examples/small.json", b)
//}
//
//func BenchmarkTransformMediumJSON(b *testing.B) {
//	transformBench("json-examples/medium.json", b)
//}
//
//func BenchmarkTransformLargeJSON(b *testing.B) {
//	transformBench("json-examples/large.json", b)
//}

//func matchBench(fileName string, b *testing.B) {
//	jsonData, err := os.ReadFile(fileName)
//	if err != nil {
//		b.Error("unable to read json: " + err.Error())
//	}
//
//	d, err := setup(Match)
//	if err != nil {
//		b.Error(err)
//	}
//
//	b.ResetTimer()
//
//	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
//	defer cancel()
//
//	for i := 0; i < b.N; i++ {
//		cfg := &protos.RuleConfigMatch{
//			Path:     "firstname",
//			Type:     "string_contains_any",
//			Operator: protos.MatchOperator_MATCH_OPERATOR_ISMATCH,
//			Args:     []string{"Rani"},
//		}
//		_, err := d.runMatch(ctx, jsonData, cfg)
//		if err != nil {
//			cancel()
//			b.Fatal("error during runMatch: " + err.Error())
//		}
//		cancel()
//	}
//}
//
//func transformBench(fileName string, b *testing.B) {
//	jsonData, err := os.ReadFile(fileName)
//	if err != nil {
//		b.Error("unable to read json: " + err.Error())
//	}
//
//	d, err := setup(Transform)
//	if err != nil {
//		b.Error(err)
//	}
//
//	b.ResetTimer()
//
//	fm := &protos.FailureModeTransform{
//		Type:  protos.FailureModeTransform_TRANSFORM_TYPE_REPLACE,
//		Path:  "firstname",
//		Value: "Testing",
//	}
//
//	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
//	defer cancel()
//
//	for i := 0; i < b.N; i++ {
//		_, err := d.failTransform(ctx, jsonData, fm)
//		if err != nil {
//			b.Error("error during runTransform: " + err.Error())
//		}
//	}
//}
