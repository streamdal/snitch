package util

import (
	"context"
	"crypto/sha256"
	"fmt"
	"strings"
	"time"

	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"

	"github.com/streamdal/streamdal/libs/protos/build/go/protos"
	"github.com/streamdal/streamdal/libs/protos/build/go/protos/shared"
	"github.com/streamdal/streamdal/libs/protos/build/go/protos/steps"

	"github.com/streamdal/streamdal/apps/server/wasm"
)

const (
	GRPCRequestIDMetadataKey = "request-id"
)

func GenerateUUID() string {
	v, err := uuid.NewV4()
	if err != nil {
		panic("unable to generate v4 uuid: " + err.Error())
	}

	return v.String()
}

// CtxMetadata will return a map[string]string of the metadata in the given context (if md exists).
// NOTE: Will use only first value in metadata value slice.
func CtxMetadata(ctx context.Context) map[string]string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return map[string]string{}
	}

	m := map[string]string{}

	for k, v := range md {
		if len(v) < 1 {
			continue
		}

		m[k] = v[0]
	}

	return m
}

func CtxStringValue(ctx context.Context, key string) string {
	if ctx == nil {
		return ""
	}

	values := metadata.ValueFromIncomingContext(ctx, key)
	if len(values) == 0 {
		return ""
	}

	return values[0]
}

func Pointer[T any](val T) *T {
	return &val
}

func CtxRequestId(ctx context.Context) string {
	return CtxStringValue(ctx, GRPCRequestIDMetadataKey)
}

func AudienceToStr(audience *protos.Audience) string {
	if audience == nil {
		return ""
	}

	str := strings.ToLower(fmt.Sprintf("%s:%s:%s:%s", audience.ServiceName, audience.OperationType, audience.OperationName, audience.ComponentName))

	return str
}

// AudienceFromStr will parse a string into an Audience. If the string is invalid,
// nil will be returned.
func AudienceFromStr(s string) *protos.Audience {
	if s == "" {
		return nil
	}

	parts := strings.Split(s, ":")
	if len(parts) != 4 {
		return nil
	}

	opType := protos.OperationType_OPERATION_TYPE_UNSET

	if parts[1] == strings.ToLower(protos.OperationType_OPERATION_TYPE_CONSUMER.String()) {
		opType = protos.OperationType_OPERATION_TYPE_CONSUMER
	} else {
		opType = protos.OperationType_OPERATION_TYPE_PRODUCER
	}

	return &protos.Audience{
		ServiceName:   strings.ToLower(parts[0]),
		OperationType: opType,
		OperationName: strings.ToLower(parts[2]),
		ComponentName: strings.ToLower(parts[3]),
	}
}

func StandardResponse(ctx context.Context, code protos.ResponseCode, msg string) *protos.StandardResponse {
	reqId := CtxRequestId(ctx)

	return &protos.StandardResponse{
		Id:      reqId,
		Code:    code,
		Message: msg,
	}
}

// PopulateWASMFields is used for populating WASM in *protos.Pipeline because
// the SDK may not have had audiences at startup and thus GetSetPipelinesByService()
// would not have returned any WASM data.
func PopulateWASMFields(pipeline *protos.Pipeline, prefix string) error {
	if pipeline == nil {
		return errors.New("pipeline cannot be nil")
	}

	for _, s := range pipeline.Steps {
		var (
			mapping *wasm.Mapping
			err     error
		)

		// We can do this dynamically later
		switch s.Step.(type) {
		case *protos.PipelineStep_Detective:
			mapping, err = wasm.Load("detective", prefix)
		case *protos.PipelineStep_Transform:
			mapping, err = wasm.Load("transform", prefix)
		case *protos.PipelineStep_Kv:
			mapping, err = wasm.Load("kv", prefix)
		case *protos.PipelineStep_HttpRequest:
			mapping, err = wasm.Load("httprequest", prefix)
		case *protos.PipelineStep_InferSchema:
			mapping, err = wasm.Load("inferschema", prefix)
		case *protos.PipelineStep_SchemaValidation:
			mapping, err = wasm.Load("schemavalidation", prefix)
		default:
			return errors.Errorf("unknown pipeline step type: %T", s.Step)
		}

		if err != nil {
			return errors.Wrapf(err, "error loading '%T' WASM mapping", s.Step)
		}

		s.XWasmFunction = &mapping.FuncName
		s.XWasmBytes = mapping.Contents
		s.XWasmId = &mapping.ID
	}

	return nil
}

// GenerateWasmMapping will generate a map of WASM modules from the given command(s).
// NOTE: This is primarily useful for commands that have Steps which contain
// Wasm fields (like SetPipelines command). For commands that do not have Steps
// w/ Wasm, this will do nothing.
// This is used _specifically_ for the initial GetSetPipelinesByService() that
// SDKs call on startup (but if the SDK does not have any audiences, this will
// be empty).
func GenerateWasmMapping(commands ...*protos.Command) map[string]*protos.WasmModule {
	wasmModules := make(map[string]*protos.WasmModule)

	for _, cmd := range commands {
		if cmd.GetSetPipelines() == nil {
			continue
		}

		if cmd.GetSetPipelines().Pipelines == nil {
			logrus.Warnf("bug? attach pipeline command has nil pipelines. Audience: %s CommandStr: %s",
				AudienceToStr(cmd.Audience), cmd.String())
			continue
		}

		// Inject WASM data into its own map and zero out the bytes in the steps
		// This is to prevent the WASM data from being duplicated in the response
		for _, pipeline := range cmd.GetSetPipelines().Pipelines {
			for _, step := range pipeline.Steps {
				if _, ok := wasmModules[step.GetXWasmId()]; ok {
					step.XWasmBytes = nil
					continue
				}

				wasmModules[step.GetXWasmId()] = &protos.WasmModule{
					Id:       step.GetXWasmId(),
					Bytes:    step.GetXWasmBytes(),
					Function: step.GetXWasmFunction(),
				}
				step.XWasmBytes = nil
			}
		}
	}

	return wasmModules
}

func StripWASMFields(pipeline *protos.Pipeline) {
	if pipeline == nil {
		return
	}

	for _, s := range pipeline.Steps {
		s.XWasmFunction = nil
		s.XWasmBytes = nil
	}
}

// ConvertPipelineConfigsAudToStr is used for converting a map of *protos.Audience:*protos.PipelineConfigs
// to map of string:*protos.PipelineConfigs.
func ConvertPipelineConfigsAudToStr(configs map[*protos.Audience]*protos.PipelineConfigs) map[string]*protos.PipelineConfigs {
	if configs == nil {
		return nil
	}

	m := map[string]*protos.PipelineConfigs{}

	for k, pipelineConfigs := range configs {
		m[AudienceToStr(k)] = pipelineConfigs
	}

	return m
}

func GenerateSchemaInferencePipeline(wasmDir string) (*protos.Pipeline, error) {
	pipeline := &protos.Pipeline{
		Id:   GenerateUUID(),
		Name: "Schema Inference (auto-generated pipeline)",
		Steps: []*protos.PipelineStep{
			{
				Name: "Infer Schema (auto-generated step)",
				Step: &protos.PipelineStep_InferSchema{
					InferSchema: &steps.InferSchemaStep{
						CurrentSchema: make([]byte, 0),
					},
				},
			},
		},
	}

	if err := PopulateWASMFields(pipeline, wasmDir); err != nil {
		return nil, errors.Wrap(err, "error populating WASM fields")
	}

	return pipeline, nil
}

// InjectSchemaInferenceForSetPipelinesCommands is a helper function for injecting
// a schema inference pipeline into a slice of SetPipelines commands. This is
// basically InjectSchemaInferenceForPipeline() but for commands.
func InjectSchemaInferenceForSetPipelinesCommands(
	cmds []*protos.Command,
	wasmDir string,
) (int, error) {
	if len(cmds) == 0 {
		return 0, nil
	}

	var numInjected int

	for _, cmd := range cmds {
		if cmd.GetSetPipelines() == nil {
			fmt.Printf("skipping injection for non-SetPipelines command for audience '%s'\n", AudienceToStr(cmd.Audience))
			continue
		}

		updatedPipelines, err := InjectSchemaInferenceForPipelines(cmd.GetSetPipelines().Pipelines, wasmDir)
		if err != nil {
			return 0, errors.Wrap(err, "error injecting schema inference pipeline")
		}

		cmd.GetSetPipelines().Pipelines = updatedPipelines

		numInjected += 1
	}

	return numInjected, nil
}

// InjectSchemaInferenceForPipelines will inject a schema inference pipeline into
// the given slice of pipelines. This is useful
func InjectSchemaInferenceForPipelines(pipelines []*protos.Pipeline, wasmDir string) ([]*protos.Pipeline, error) {
	schemaInferencePipeline, err := GenerateSchemaInferencePipeline(wasmDir)
	if err != nil {
		return nil, errors.Wrap(err, "unable to generate schema inference pipeline")
	}

	return append([]*protos.Pipeline{schemaInferencePipeline}, pipelines...), nil
}

func AudienceEquals(a, b *protos.Audience) bool {
	if a == nil || b == nil {
		return false
	}

	return AudienceToStr(a) == AudienceToStr(b)
}

// GenerateKVRequest is used for converting from plain KVObjects to instructions.
// This func is used to simplify transforming KV HTTP requests -> KV broadcast requests.
func GenerateKVRequest(action shared.KVAction, kvs []*protos.KVObject, overwrite bool) *protos.KVRequest {
	instructions := make([]*protos.KVInstruction, 0)

	for _, kv := range kvs {
		i := &protos.KVInstruction{
			Id:                       GenerateUUID(),
			Action:                   action,
			Object:                   kv,
			RequestedAtUnixTsNanoUtc: time.Now().UTC().UnixNano(),
		}

		instructions = append(instructions, i)
	}

	return &protos.KVRequest{
		Instructions: instructions,
		Overwrite:    overwrite,
	}
}

func CounterName(name string, labels map[string]string) string {
	vals := make([]string, 0)
	for k, v := range labels {
		vals = append(vals, fmt.Sprintf("%s-%s", k, v))
	}

	return fmt.Sprintf("%s-%s", name, strings.Join(vals, "-"))
}

// GrpcMethodCounterName turns a gRPC method name into a counter name for statsd
// This turns a string such as "/protos.External/GetSchema" into "grpc_method_external_get_schema_total"
func GrpcMethodCounterName(method string) string {
	parts := strings.Split(method, "/")
	if len(parts) != 3 {
		return ""
	}

	which := "external"
	if parts[1] == "protos.Internal" {
		which = "internal"
	}

	// Loop over each character in the method name and replace upper case characters with an underscore
	// and the lowercase version
	var op string
	for _, c := range parts[2] {
		if c >= 'A' && c <= 'Z' {
			op += fmt.Sprintf("_%c", c+32)
		} else {
			op += string(c)
		}
	}

	// Trim leading underscore since all gRPC methods start with a capital letter
	op = strings.TrimPrefix(op, "_")

	return fmt.Sprintf("grpc_method_%s_%s_total", which, op)
}

// GenerateNodeID will generate a uuid node ID from the given install ID and node name
// This ensures that the node ID is deterministic and consistent across restarts since
// we have no local storage.
func GenerateNodeID(installID, nodeName string) string {
	hash := sha256.Sum256([]byte(installID + nodeName))

	id, err := uuid.FromBytes(hash[16:])
	if err != nil {
		return GenerateUUID()
	}

	return id.String()
}

// GetStepType transforms a protobuf enum of a step type into a string for use with telemetry
func GetStepType(step *protos.PipelineStep) string {
	switch {
	case step.GetDetective() != nil:
		return "detective"
	case step.GetTransform() != nil:
		return "transform"
	case step.GetKv() != nil:
		return "kv"
	case step.GetHttpRequest() != nil:
		return "http"
	case step.GetEncode() != nil:
		return "encode"
	case step.GetDecode() != nil:
		return "decode"
	case step.GetValidJson() != nil:
		return "valid_json"
	default:
		return "unknown"
	}
}

// GetStepSubType transforms a protobuf enum into a string for use with telemetry
// This is used for detective, transform, and KV steps.
func GetStepSubType(step *protos.PipelineStep) string {
	var st string
	switch {
	case step.GetDetective() != nil:
		st = strings.Replace(step.GetDetective().Type.String(), "DETECTIVE_TYPE_", "", -1)
	case step.GetTransform() != nil:
		st = strings.Replace(step.GetTransform().Type.String(), "TRANSFORM_TYPE_", "", -1)
	case step.GetKv() != nil:
		st = strings.Replace(step.GetKv().Action.String(), "KVAction_KV_ACTION_", "", -1)
	default:
		st = ""
	}

	return strings.ToLower(st)
}

func AudienceInList(audience *protos.Audience, list []*protos.Audience) bool {
	for _, a := range list {
		if AudienceEquals(a, audience) {
			return true
		}
	}

	return false
}
