// @generated by protobuf-ts 2.9.0 with parameter long_type_string
// @generated from protobuf file "sp_external.proto" (package "protos", syntax proto3)
// tslint:disable
import type { RpcTransport } from "@protobuf-ts/runtime-rpc";
import type { ServiceInfo } from "@protobuf-ts/runtime-rpc";
import { External } from "./sp_external.js";
import type { TestResponse } from "./sp_external.js";
import type { TestRequest } from "./sp_external.js";
import type { DeleteCustomWasmRequest } from "./sp_external.js";
import type { UpdateCustomWasmRequest } from "./sp_external.js";
import type { CreateCustomWasmRequest } from "./sp_external.js";
import type { GetAllCustomWasmResponse } from "./sp_external.js";
import type { GetAllCustomWasmRequest } from "./sp_external.js";
import type { GetCustomWasmResponse } from "./sp_external.js";
import type { GetCustomWasmRequest } from "./sp_external.js";
import type { AppRegisterRejectRequest } from "./sp_external.js";
import type { AppVerifyRegistrationRequest } from "./sp_external.js";
import type { AppRegistrationRequest } from "./sp_external.js";
import type { AppRegistrationStatusResponse } from "./sp_external.js";
import type { AppRegistrationStatusRequest } from "./sp_external.js";
import type { GetSchemaResponse } from "./sp_external.js";
import type { GetSchemaRequest } from "./sp_external.js";
import type { GetAudienceRatesResponse } from "./sp_external.js";
import type { GetAudienceRatesRequest } from "./sp_external.js";
import type { ResumeTailRequest } from "./sp_external.js";
import type { PauseTailRequest } from "./sp_external.js";
import type { TailResponse } from "./sp_common.js";
import type { TailRequest } from "./sp_common.js";
import type { GetMetricsResponse } from "./sp_external.js";
import type { GetMetricsRequest } from "./sp_external.js";
import type { DeleteServiceRequest } from "./sp_external.js";
import type { DeleteAudienceRequest } from "./sp_external.js";
import type { CreateAudienceRequest } from "./sp_external.js";
import type { DetachNotificationRequest } from "./sp_external.js";
import type { AttachNotificationRequest } from "./sp_external.js";
import type { GetNotificationResponse } from "./sp_external.js";
import type { GetNotificationRequest } from "./sp_external.js";
import type { GetNotificationsResponse } from "./sp_external.js";
import type { GetNotificationsRequest } from "./sp_external.js";
import type { DeleteNotificationRequest } from "./sp_external.js";
import type { UpdateNotificationRequest } from "./sp_external.js";
import type { CreateNotificationResponse } from "./sp_external.js";
import type { CreateNotificationRequest } from "./sp_external.js";
import type { ResumePipelineRequest } from "./sp_external.js";
import type { PausePipelineRequest } from "./sp_external.js";
import type { SetPipelinesRequest } from "./sp_external.js";
import type { DeletePipelineRequest } from "./sp_external.js";
import type { StandardResponse } from "./sp_common.js";
import type { UpdatePipelineRequest } from "./sp_external.js";
import type { CreatePipelineResponse } from "./sp_external.js";
import type { CreatePipelineRequest } from "./sp_external.js";
import type { GetPipelineResponse } from "./sp_external.js";
import type { GetPipelineRequest } from "./sp_external.js";
import type { GetPipelinesResponse } from "./sp_external.js";
import type { GetPipelinesRequest } from "./sp_external.js";
import type { ServerStreamingCall } from "@protobuf-ts/runtime-rpc";
import { stackIntercept } from "@protobuf-ts/runtime-rpc";
import type { GetAllResponse } from "./sp_external.js";
import type { GetAllRequest } from "./sp_external.js";
import type { UnaryCall } from "@protobuf-ts/runtime-rpc";
import type { RpcOptions } from "@protobuf-ts/runtime-rpc";
/**
 * @generated from protobuf service protos.External
 */
export interface IExternalClient {
    /**
     * Returns all data needed for UI; called on initial console load
     *
     * @generated from protobuf rpc: GetAll(protos.GetAllRequest) returns (protos.GetAllResponse);
     */
    getAll(input: GetAllRequest, options?: RpcOptions): UnaryCall<GetAllRequest, GetAllResponse>;
    /**
     * Used by console to stream updates to UI; called after initial GetAll()
     *
     * @generated from protobuf rpc: GetAllStream(protos.GetAllRequest) returns (stream protos.GetAllResponse);
     */
    getAllStream(input: GetAllRequest, options?: RpcOptions): ServerStreamingCall<GetAllRequest, GetAllResponse>;
    /**
     * Returns pipelines (_wasm_bytes field is stripped)
     *
     * @generated from protobuf rpc: GetPipelines(protos.GetPipelinesRequest) returns (protos.GetPipelinesResponse);
     */
    getPipelines(input: GetPipelinesRequest, options?: RpcOptions): UnaryCall<GetPipelinesRequest, GetPipelinesResponse>;
    /**
     * Returns a single pipeline (_wasm_bytes field is stripped)
     *
     * @generated from protobuf rpc: GetPipeline(protos.GetPipelineRequest) returns (protos.GetPipelineResponse);
     */
    getPipeline(input: GetPipelineRequest, options?: RpcOptions): UnaryCall<GetPipelineRequest, GetPipelineResponse>;
    /**
     * Create a new pipeline; id must be left empty on create
     *
     * @generated from protobuf rpc: CreatePipeline(protos.CreatePipelineRequest) returns (protos.CreatePipelineResponse);
     */
    createPipeline(input: CreatePipelineRequest, options?: RpcOptions): UnaryCall<CreatePipelineRequest, CreatePipelineResponse>;
    /**
     * Update an existing pipeline; id must be set
     *
     * @generated from protobuf rpc: UpdatePipeline(protos.UpdatePipelineRequest) returns (protos.StandardResponse);
     */
    updatePipeline(input: UpdatePipelineRequest, options?: RpcOptions): UnaryCall<UpdatePipelineRequest, StandardResponse>;
    /**
     * Delete a pipeline
     *
     * @generated from protobuf rpc: DeletePipeline(protos.DeletePipelineRequest) returns (protos.StandardResponse);
     */
    deletePipeline(input: DeletePipelineRequest, options?: RpcOptions): UnaryCall<DeletePipelineRequest, StandardResponse>;
    /**
     * @generated from protobuf rpc: SetPipelines(protos.SetPipelinesRequest) returns (protos.StandardResponse);
     */
    setPipelines(input: SetPipelinesRequest, options?: RpcOptions): UnaryCall<SetPipelinesRequest, StandardResponse>;
    /**
     * Pause a pipeline; noop if pipeline is already paused
     *
     * @generated from protobuf rpc: PausePipeline(protos.PausePipelineRequest) returns (protos.StandardResponse);
     */
    pausePipeline(input: PausePipelineRequest, options?: RpcOptions): UnaryCall<PausePipelineRequest, StandardResponse>;
    /**
     * Resume a pipeline; noop if pipeline is not paused
     *
     * @generated from protobuf rpc: ResumePipeline(protos.ResumePipelineRequest) returns (protos.StandardResponse);
     */
    resumePipeline(input: ResumePipelineRequest, options?: RpcOptions): UnaryCall<ResumePipelineRequest, StandardResponse>;
    /**
     * Create a new notification config
     *
     * @generated from protobuf rpc: CreateNotification(protos.CreateNotificationRequest) returns (protos.CreateNotificationResponse);
     */
    createNotification(input: CreateNotificationRequest, options?: RpcOptions): UnaryCall<CreateNotificationRequest, CreateNotificationResponse>;
    /**
     * Update an existing notification config
     *
     * @generated from protobuf rpc: UpdateNotification(protos.UpdateNotificationRequest) returns (protos.StandardResponse);
     */
    updateNotification(input: UpdateNotificationRequest, options?: RpcOptions): UnaryCall<UpdateNotificationRequest, StandardResponse>;
    /**
     * Delete a notification config
     *
     * @generated from protobuf rpc: DeleteNotification(protos.DeleteNotificationRequest) returns (protos.StandardResponse);
     */
    deleteNotification(input: DeleteNotificationRequest, options?: RpcOptions): UnaryCall<DeleteNotificationRequest, StandardResponse>;
    /**
     * Returns all notification configs
     *
     * @generated from protobuf rpc: GetNotifications(protos.GetNotificationsRequest) returns (protos.GetNotificationsResponse);
     */
    getNotifications(input: GetNotificationsRequest, options?: RpcOptions): UnaryCall<GetNotificationsRequest, GetNotificationsResponse>;
    /**
     * Returns a single notification config
     *
     * @generated from protobuf rpc: GetNotification(protos.GetNotificationRequest) returns (protos.GetNotificationResponse);
     */
    getNotification(input: GetNotificationRequest, options?: RpcOptions): UnaryCall<GetNotificationRequest, GetNotificationResponse>;
    /**
     * Attach a notification config to a pipeline
     *
     * @deprecated
     * @generated from protobuf rpc: AttachNotification(protos.AttachNotificationRequest) returns (protos.StandardResponse);
     */
    attachNotification(input: AttachNotificationRequest, options?: RpcOptions): UnaryCall<AttachNotificationRequest, StandardResponse>;
    /**
     * Detach a notification config from a pipeline
     *
     * @deprecated
     * @generated from protobuf rpc: DetachNotification(protos.DetachNotificationRequest) returns (protos.StandardResponse);
     */
    detachNotification(input: DetachNotificationRequest, options?: RpcOptions): UnaryCall<DetachNotificationRequest, StandardResponse>;
    /**
     * Create an audience. Used for terraform purposes
     *
     * @generated from protobuf rpc: CreateAudience(protos.CreateAudienceRequest) returns (protos.StandardResponse);
     */
    createAudience(input: CreateAudienceRequest, options?: RpcOptions): UnaryCall<CreateAudienceRequest, StandardResponse>;
    /**
     * Delete an audience
     *
     * @generated from protobuf rpc: DeleteAudience(protos.DeleteAudienceRequest) returns (protos.StandardResponse);
     */
    deleteAudience(input: DeleteAudienceRequest, options?: RpcOptions): UnaryCall<DeleteAudienceRequest, StandardResponse>;
    /**
     * Delete a service and all associated audiences
     *
     * @generated from protobuf rpc: DeleteService(protos.DeleteServiceRequest) returns (protos.StandardResponse);
     */
    deleteService(input: DeleteServiceRequest, options?: RpcOptions): UnaryCall<DeleteServiceRequest, StandardResponse>;
    /**
     * Returns all metric counters
     *
     * @generated from protobuf rpc: GetMetrics(protos.GetMetricsRequest) returns (stream protos.GetMetricsResponse);
     */
    getMetrics(input: GetMetricsRequest, options?: RpcOptions): ServerStreamingCall<GetMetricsRequest, GetMetricsResponse>;
    /**
     * @generated from protobuf rpc: Tail(protos.TailRequest) returns (stream protos.TailResponse);
     */
    tail(input: TailRequest, options?: RpcOptions): ServerStreamingCall<TailRequest, TailResponse>;
    /**
     * @generated from protobuf rpc: PauseTail(protos.PauseTailRequest) returns (protos.StandardResponse);
     */
    pauseTail(input: PauseTailRequest, options?: RpcOptions): UnaryCall<PauseTailRequest, StandardResponse>;
    /**
     * @generated from protobuf rpc: ResumeTail(protos.ResumeTailRequest) returns (protos.StandardResponse);
     */
    resumeTail(input: ResumeTailRequest, options?: RpcOptions): UnaryCall<ResumeTailRequest, StandardResponse>;
    /**
     * @generated from protobuf rpc: GetAudienceRates(protos.GetAudienceRatesRequest) returns (stream protos.GetAudienceRatesResponse);
     */
    getAudienceRates(input: GetAudienceRatesRequest, options?: RpcOptions): ServerStreamingCall<GetAudienceRatesRequest, GetAudienceRatesResponse>;
    /**
     * @generated from protobuf rpc: GetSchema(protos.GetSchemaRequest) returns (protos.GetSchemaResponse);
     */
    getSchema(input: GetSchemaRequest, options?: RpcOptions): UnaryCall<GetSchemaRequest, GetSchemaResponse>;
    /**
     * @generated from protobuf rpc: AppRegistrationStatus(protos.AppRegistrationStatusRequest) returns (protos.AppRegistrationStatusResponse);
     */
    appRegistrationStatus(input: AppRegistrationStatusRequest, options?: RpcOptions): UnaryCall<AppRegistrationStatusRequest, AppRegistrationStatusResponse>;
    /**
     * @generated from protobuf rpc: AppRegister(protos.AppRegistrationRequest) returns (protos.StandardResponse);
     */
    appRegister(input: AppRegistrationRequest, options?: RpcOptions): UnaryCall<AppRegistrationRequest, StandardResponse>;
    /**
     * @generated from protobuf rpc: AppVerifyRegistration(protos.AppVerifyRegistrationRequest) returns (protos.StandardResponse);
     */
    appVerifyRegistration(input: AppVerifyRegistrationRequest, options?: RpcOptions): UnaryCall<AppVerifyRegistrationRequest, StandardResponse>;
    /**
     * @generated from protobuf rpc: AppRegisterReject(protos.AppRegisterRejectRequest) returns (protos.StandardResponse);
     */
    appRegisterReject(input: AppRegisterRejectRequest, options?: RpcOptions): UnaryCall<AppRegisterRejectRequest, StandardResponse>;
    /**
     * BEGIN Custom Wasm methods
     *
     * @generated from protobuf rpc: GetCustomWasm(protos.GetCustomWasmRequest) returns (protos.GetCustomWasmResponse);
     */
    getCustomWasm(input: GetCustomWasmRequest, options?: RpcOptions): UnaryCall<GetCustomWasmRequest, GetCustomWasmResponse>;
    /**
     * @generated from protobuf rpc: GetAllCustomWasm(protos.GetAllCustomWasmRequest) returns (protos.GetAllCustomWasmResponse);
     */
    getAllCustomWasm(input: GetAllCustomWasmRequest, options?: RpcOptions): UnaryCall<GetAllCustomWasmRequest, GetAllCustomWasmResponse>;
    /**
     * @generated from protobuf rpc: CreateCustomWasm(protos.CreateCustomWasmRequest) returns (protos.StandardResponse);
     */
    createCustomWasm(input: CreateCustomWasmRequest, options?: RpcOptions): UnaryCall<CreateCustomWasmRequest, StandardResponse>;
    /**
     * @generated from protobuf rpc: UpdateCustomWasm(protos.UpdateCustomWasmRequest) returns (protos.StandardResponse);
     */
    updateCustomWasm(input: UpdateCustomWasmRequest, options?: RpcOptions): UnaryCall<UpdateCustomWasmRequest, StandardResponse>;
    /**
     * END Custom Wasm methods
     *
     * @generated from protobuf rpc: DeleteCustomWasm(protos.DeleteCustomWasmRequest) returns (protos.StandardResponse);
     */
    deleteCustomWasm(input: DeleteCustomWasmRequest, options?: RpcOptions): UnaryCall<DeleteCustomWasmRequest, StandardResponse>;
    /**
     * Test method
     *
     * @generated from protobuf rpc: Test(protos.TestRequest) returns (protos.TestResponse);
     */
    test(input: TestRequest, options?: RpcOptions): UnaryCall<TestRequest, TestResponse>;
}
/**
 * @generated from protobuf service protos.External
 */
export class ExternalClient implements IExternalClient, ServiceInfo {
    typeName = External.typeName;
    methods = External.methods;
    options = External.options;
    constructor(private readonly _transport: RpcTransport) {
    }
    /**
     * Returns all data needed for UI; called on initial console load
     *
     * @generated from protobuf rpc: GetAll(protos.GetAllRequest) returns (protos.GetAllResponse);
     */
    getAll(input: GetAllRequest, options?: RpcOptions): UnaryCall<GetAllRequest, GetAllResponse> {
        const method = this.methods[0], opt = this._transport.mergeOptions(options);
        return stackIntercept<GetAllRequest, GetAllResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * Used by console to stream updates to UI; called after initial GetAll()
     *
     * @generated from protobuf rpc: GetAllStream(protos.GetAllRequest) returns (stream protos.GetAllResponse);
     */
    getAllStream(input: GetAllRequest, options?: RpcOptions): ServerStreamingCall<GetAllRequest, GetAllResponse> {
        const method = this.methods[1], opt = this._transport.mergeOptions(options);
        return stackIntercept<GetAllRequest, GetAllResponse>("serverStreaming", this._transport, method, opt, input);
    }
    /**
     * Returns pipelines (_wasm_bytes field is stripped)
     *
     * @generated from protobuf rpc: GetPipelines(protos.GetPipelinesRequest) returns (protos.GetPipelinesResponse);
     */
    getPipelines(input: GetPipelinesRequest, options?: RpcOptions): UnaryCall<GetPipelinesRequest, GetPipelinesResponse> {
        const method = this.methods[2], opt = this._transport.mergeOptions(options);
        return stackIntercept<GetPipelinesRequest, GetPipelinesResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * Returns a single pipeline (_wasm_bytes field is stripped)
     *
     * @generated from protobuf rpc: GetPipeline(protos.GetPipelineRequest) returns (protos.GetPipelineResponse);
     */
    getPipeline(input: GetPipelineRequest, options?: RpcOptions): UnaryCall<GetPipelineRequest, GetPipelineResponse> {
        const method = this.methods[3], opt = this._transport.mergeOptions(options);
        return stackIntercept<GetPipelineRequest, GetPipelineResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * Create a new pipeline; id must be left empty on create
     *
     * @generated from protobuf rpc: CreatePipeline(protos.CreatePipelineRequest) returns (protos.CreatePipelineResponse);
     */
    createPipeline(input: CreatePipelineRequest, options?: RpcOptions): UnaryCall<CreatePipelineRequest, CreatePipelineResponse> {
        const method = this.methods[4], opt = this._transport.mergeOptions(options);
        return stackIntercept<CreatePipelineRequest, CreatePipelineResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * Update an existing pipeline; id must be set
     *
     * @generated from protobuf rpc: UpdatePipeline(protos.UpdatePipelineRequest) returns (protos.StandardResponse);
     */
    updatePipeline(input: UpdatePipelineRequest, options?: RpcOptions): UnaryCall<UpdatePipelineRequest, StandardResponse> {
        const method = this.methods[5], opt = this._transport.mergeOptions(options);
        return stackIntercept<UpdatePipelineRequest, StandardResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * Delete a pipeline
     *
     * @generated from protobuf rpc: DeletePipeline(protos.DeletePipelineRequest) returns (protos.StandardResponse);
     */
    deletePipeline(input: DeletePipelineRequest, options?: RpcOptions): UnaryCall<DeletePipelineRequest, StandardResponse> {
        const method = this.methods[6], opt = this._transport.mergeOptions(options);
        return stackIntercept<DeletePipelineRequest, StandardResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: SetPipelines(protos.SetPipelinesRequest) returns (protos.StandardResponse);
     */
    setPipelines(input: SetPipelinesRequest, options?: RpcOptions): UnaryCall<SetPipelinesRequest, StandardResponse> {
        const method = this.methods[7], opt = this._transport.mergeOptions(options);
        return stackIntercept<SetPipelinesRequest, StandardResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * Pause a pipeline; noop if pipeline is already paused
     *
     * @generated from protobuf rpc: PausePipeline(protos.PausePipelineRequest) returns (protos.StandardResponse);
     */
    pausePipeline(input: PausePipelineRequest, options?: RpcOptions): UnaryCall<PausePipelineRequest, StandardResponse> {
        const method = this.methods[8], opt = this._transport.mergeOptions(options);
        return stackIntercept<PausePipelineRequest, StandardResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * Resume a pipeline; noop if pipeline is not paused
     *
     * @generated from protobuf rpc: ResumePipeline(protos.ResumePipelineRequest) returns (protos.StandardResponse);
     */
    resumePipeline(input: ResumePipelineRequest, options?: RpcOptions): UnaryCall<ResumePipelineRequest, StandardResponse> {
        const method = this.methods[9], opt = this._transport.mergeOptions(options);
        return stackIntercept<ResumePipelineRequest, StandardResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * Create a new notification config
     *
     * @generated from protobuf rpc: CreateNotification(protos.CreateNotificationRequest) returns (protos.CreateNotificationResponse);
     */
    createNotification(input: CreateNotificationRequest, options?: RpcOptions): UnaryCall<CreateNotificationRequest, CreateNotificationResponse> {
        const method = this.methods[10], opt = this._transport.mergeOptions(options);
        return stackIntercept<CreateNotificationRequest, CreateNotificationResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * Update an existing notification config
     *
     * @generated from protobuf rpc: UpdateNotification(protos.UpdateNotificationRequest) returns (protos.StandardResponse);
     */
    updateNotification(input: UpdateNotificationRequest, options?: RpcOptions): UnaryCall<UpdateNotificationRequest, StandardResponse> {
        const method = this.methods[11], opt = this._transport.mergeOptions(options);
        return stackIntercept<UpdateNotificationRequest, StandardResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * Delete a notification config
     *
     * @generated from protobuf rpc: DeleteNotification(protos.DeleteNotificationRequest) returns (protos.StandardResponse);
     */
    deleteNotification(input: DeleteNotificationRequest, options?: RpcOptions): UnaryCall<DeleteNotificationRequest, StandardResponse> {
        const method = this.methods[12], opt = this._transport.mergeOptions(options);
        return stackIntercept<DeleteNotificationRequest, StandardResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * Returns all notification configs
     *
     * @generated from protobuf rpc: GetNotifications(protos.GetNotificationsRequest) returns (protos.GetNotificationsResponse);
     */
    getNotifications(input: GetNotificationsRequest, options?: RpcOptions): UnaryCall<GetNotificationsRequest, GetNotificationsResponse> {
        const method = this.methods[13], opt = this._transport.mergeOptions(options);
        return stackIntercept<GetNotificationsRequest, GetNotificationsResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * Returns a single notification config
     *
     * @generated from protobuf rpc: GetNotification(protos.GetNotificationRequest) returns (protos.GetNotificationResponse);
     */
    getNotification(input: GetNotificationRequest, options?: RpcOptions): UnaryCall<GetNotificationRequest, GetNotificationResponse> {
        const method = this.methods[14], opt = this._transport.mergeOptions(options);
        return stackIntercept<GetNotificationRequest, GetNotificationResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * Attach a notification config to a pipeline
     *
     * @deprecated
     * @generated from protobuf rpc: AttachNotification(protos.AttachNotificationRequest) returns (protos.StandardResponse);
     */
    attachNotification(input: AttachNotificationRequest, options?: RpcOptions): UnaryCall<AttachNotificationRequest, StandardResponse> {
        const method = this.methods[15], opt = this._transport.mergeOptions(options);
        return stackIntercept<AttachNotificationRequest, StandardResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * Detach a notification config from a pipeline
     *
     * @deprecated
     * @generated from protobuf rpc: DetachNotification(protos.DetachNotificationRequest) returns (protos.StandardResponse);
     */
    detachNotification(input: DetachNotificationRequest, options?: RpcOptions): UnaryCall<DetachNotificationRequest, StandardResponse> {
        const method = this.methods[16], opt = this._transport.mergeOptions(options);
        return stackIntercept<DetachNotificationRequest, StandardResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * Create an audience. Used for terraform purposes
     *
     * @generated from protobuf rpc: CreateAudience(protos.CreateAudienceRequest) returns (protos.StandardResponse);
     */
    createAudience(input: CreateAudienceRequest, options?: RpcOptions): UnaryCall<CreateAudienceRequest, StandardResponse> {
        const method = this.methods[17], opt = this._transport.mergeOptions(options);
        return stackIntercept<CreateAudienceRequest, StandardResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * Delete an audience
     *
     * @generated from protobuf rpc: DeleteAudience(protos.DeleteAudienceRequest) returns (protos.StandardResponse);
     */
    deleteAudience(input: DeleteAudienceRequest, options?: RpcOptions): UnaryCall<DeleteAudienceRequest, StandardResponse> {
        const method = this.methods[18], opt = this._transport.mergeOptions(options);
        return stackIntercept<DeleteAudienceRequest, StandardResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * Delete a service and all associated audiences
     *
     * @generated from protobuf rpc: DeleteService(protos.DeleteServiceRequest) returns (protos.StandardResponse);
     */
    deleteService(input: DeleteServiceRequest, options?: RpcOptions): UnaryCall<DeleteServiceRequest, StandardResponse> {
        const method = this.methods[19], opt = this._transport.mergeOptions(options);
        return stackIntercept<DeleteServiceRequest, StandardResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * Returns all metric counters
     *
     * @generated from protobuf rpc: GetMetrics(protos.GetMetricsRequest) returns (stream protos.GetMetricsResponse);
     */
    getMetrics(input: GetMetricsRequest, options?: RpcOptions): ServerStreamingCall<GetMetricsRequest, GetMetricsResponse> {
        const method = this.methods[20], opt = this._transport.mergeOptions(options);
        return stackIntercept<GetMetricsRequest, GetMetricsResponse>("serverStreaming", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: Tail(protos.TailRequest) returns (stream protos.TailResponse);
     */
    tail(input: TailRequest, options?: RpcOptions): ServerStreamingCall<TailRequest, TailResponse> {
        const method = this.methods[21], opt = this._transport.mergeOptions(options);
        return stackIntercept<TailRequest, TailResponse>("serverStreaming", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: PauseTail(protos.PauseTailRequest) returns (protos.StandardResponse);
     */
    pauseTail(input: PauseTailRequest, options?: RpcOptions): UnaryCall<PauseTailRequest, StandardResponse> {
        const method = this.methods[22], opt = this._transport.mergeOptions(options);
        return stackIntercept<PauseTailRequest, StandardResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: ResumeTail(protos.ResumeTailRequest) returns (protos.StandardResponse);
     */
    resumeTail(input: ResumeTailRequest, options?: RpcOptions): UnaryCall<ResumeTailRequest, StandardResponse> {
        const method = this.methods[23], opt = this._transport.mergeOptions(options);
        return stackIntercept<ResumeTailRequest, StandardResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: GetAudienceRates(protos.GetAudienceRatesRequest) returns (stream protos.GetAudienceRatesResponse);
     */
    getAudienceRates(input: GetAudienceRatesRequest, options?: RpcOptions): ServerStreamingCall<GetAudienceRatesRequest, GetAudienceRatesResponse> {
        const method = this.methods[24], opt = this._transport.mergeOptions(options);
        return stackIntercept<GetAudienceRatesRequest, GetAudienceRatesResponse>("serverStreaming", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: GetSchema(protos.GetSchemaRequest) returns (protos.GetSchemaResponse);
     */
    getSchema(input: GetSchemaRequest, options?: RpcOptions): UnaryCall<GetSchemaRequest, GetSchemaResponse> {
        const method = this.methods[25], opt = this._transport.mergeOptions(options);
        return stackIntercept<GetSchemaRequest, GetSchemaResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: AppRegistrationStatus(protos.AppRegistrationStatusRequest) returns (protos.AppRegistrationStatusResponse);
     */
    appRegistrationStatus(input: AppRegistrationStatusRequest, options?: RpcOptions): UnaryCall<AppRegistrationStatusRequest, AppRegistrationStatusResponse> {
        const method = this.methods[26], opt = this._transport.mergeOptions(options);
        return stackIntercept<AppRegistrationStatusRequest, AppRegistrationStatusResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: AppRegister(protos.AppRegistrationRequest) returns (protos.StandardResponse);
     */
    appRegister(input: AppRegistrationRequest, options?: RpcOptions): UnaryCall<AppRegistrationRequest, StandardResponse> {
        const method = this.methods[27], opt = this._transport.mergeOptions(options);
        return stackIntercept<AppRegistrationRequest, StandardResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: AppVerifyRegistration(protos.AppVerifyRegistrationRequest) returns (protos.StandardResponse);
     */
    appVerifyRegistration(input: AppVerifyRegistrationRequest, options?: RpcOptions): UnaryCall<AppVerifyRegistrationRequest, StandardResponse> {
        const method = this.methods[28], opt = this._transport.mergeOptions(options);
        return stackIntercept<AppVerifyRegistrationRequest, StandardResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: AppRegisterReject(protos.AppRegisterRejectRequest) returns (protos.StandardResponse);
     */
    appRegisterReject(input: AppRegisterRejectRequest, options?: RpcOptions): UnaryCall<AppRegisterRejectRequest, StandardResponse> {
        const method = this.methods[29], opt = this._transport.mergeOptions(options);
        return stackIntercept<AppRegisterRejectRequest, StandardResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * BEGIN Custom Wasm methods
     *
     * @generated from protobuf rpc: GetCustomWasm(protos.GetCustomWasmRequest) returns (protos.GetCustomWasmResponse);
     */
    getCustomWasm(input: GetCustomWasmRequest, options?: RpcOptions): UnaryCall<GetCustomWasmRequest, GetCustomWasmResponse> {
        const method = this.methods[30], opt = this._transport.mergeOptions(options);
        return stackIntercept<GetCustomWasmRequest, GetCustomWasmResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: GetAllCustomWasm(protos.GetAllCustomWasmRequest) returns (protos.GetAllCustomWasmResponse);
     */
    getAllCustomWasm(input: GetAllCustomWasmRequest, options?: RpcOptions): UnaryCall<GetAllCustomWasmRequest, GetAllCustomWasmResponse> {
        const method = this.methods[31], opt = this._transport.mergeOptions(options);
        return stackIntercept<GetAllCustomWasmRequest, GetAllCustomWasmResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: CreateCustomWasm(protos.CreateCustomWasmRequest) returns (protos.StandardResponse);
     */
    createCustomWasm(input: CreateCustomWasmRequest, options?: RpcOptions): UnaryCall<CreateCustomWasmRequest, StandardResponse> {
        const method = this.methods[32], opt = this._transport.mergeOptions(options);
        return stackIntercept<CreateCustomWasmRequest, StandardResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: UpdateCustomWasm(protos.UpdateCustomWasmRequest) returns (protos.StandardResponse);
     */
    updateCustomWasm(input: UpdateCustomWasmRequest, options?: RpcOptions): UnaryCall<UpdateCustomWasmRequest, StandardResponse> {
        const method = this.methods[33], opt = this._transport.mergeOptions(options);
        return stackIntercept<UpdateCustomWasmRequest, StandardResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * END Custom Wasm methods
     *
     * @generated from protobuf rpc: DeleteCustomWasm(protos.DeleteCustomWasmRequest) returns (protos.StandardResponse);
     */
    deleteCustomWasm(input: DeleteCustomWasmRequest, options?: RpcOptions): UnaryCall<DeleteCustomWasmRequest, StandardResponse> {
        const method = this.methods[34], opt = this._transport.mergeOptions(options);
        return stackIntercept<DeleteCustomWasmRequest, StandardResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * Test method
     *
     * @generated from protobuf rpc: Test(protos.TestRequest) returns (protos.TestResponse);
     */
    test(input: TestRequest, options?: RpcOptions): UnaryCall<TestRequest, TestResponse> {
        const method = this.methods[35], opt = this._transport.mergeOptions(options);
        return stackIntercept<TestRequest, TestResponse>("unary", this._transport, method, opt, input);
    }
}
