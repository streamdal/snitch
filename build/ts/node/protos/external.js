// @generated by protobuf-ts 2.9.0 with parameter optimize_code_size
// @generated from protobuf file "external.proto" (package "protos", syntax proto3)
// tslint:disable
import { StandardResponse } from "./common.js";
import { ServiceType } from "@protobuf-ts/runtime-rpc";
import { MessageType } from "@protobuf-ts/runtime";
import { Audience } from "./common.js";
import { Pipeline } from "./pipeline.js";
import { ServiceInfo } from "./info.js";
// @generated message type with reflection information, may provide speed optimized methods
class GetServiceMapRequest$Type extends MessageType {
    constructor() {
        super("protos.GetServiceMapRequest", []);
    }
}
/**
 * @generated MessageType for protobuf message protos.GetServiceMapRequest
 */
export const GetServiceMapRequest = new GetServiceMapRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GetServiceMapResponse$Type extends MessageType {
    constructor() {
        super("protos.GetServiceMapResponse", [
            { no: 1, name: "service_map", kind: "map", K: 9 /*ScalarType.STRING*/, V: { kind: "message", T: () => ServiceInfo } }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.GetServiceMapResponse
 */
export const GetServiceMapResponse = new GetServiceMapResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GetPipelinesRequest$Type extends MessageType {
    constructor() {
        super("protos.GetPipelinesRequest", []);
    }
}
/**
 * @generated MessageType for protobuf message protos.GetPipelinesRequest
 */
export const GetPipelinesRequest = new GetPipelinesRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GetPipelinesResponse$Type extends MessageType {
    constructor() {
        super("protos.GetPipelinesResponse", [
            { no: 1, name: "pipelines", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Pipeline }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.GetPipelinesResponse
 */
export const GetPipelinesResponse = new GetPipelinesResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GetPipelineRequest$Type extends MessageType {
    constructor() {
        super("protos.GetPipelineRequest", [
            { no: 1, name: "pipeline_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.GetPipelineRequest
 */
export const GetPipelineRequest = new GetPipelineRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GetPipelineResponse$Type extends MessageType {
    constructor() {
        super("protos.GetPipelineResponse", [
            { no: 1, name: "pipeline", kind: "message", T: () => Pipeline }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.GetPipelineResponse
 */
export const GetPipelineResponse = new GetPipelineResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class CreatePipelineRequest$Type extends MessageType {
    constructor() {
        super("protos.CreatePipelineRequest", [
            { no: 1, name: "pipeline", kind: "message", T: () => Pipeline }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.CreatePipelineRequest
 */
export const CreatePipelineRequest = new CreatePipelineRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class UpdatePipelineRequest$Type extends MessageType {
    constructor() {
        super("protos.UpdatePipelineRequest", [
            { no: 1, name: "pipeline", kind: "message", T: () => Pipeline }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.UpdatePipelineRequest
 */
export const UpdatePipelineRequest = new UpdatePipelineRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DeletePipelineRequest$Type extends MessageType {
    constructor() {
        super("protos.DeletePipelineRequest", [
            { no: 1, name: "pipeline_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.DeletePipelineRequest
 */
export const DeletePipelineRequest = new DeletePipelineRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class AttachPipelineRequest$Type extends MessageType {
    constructor() {
        super("protos.AttachPipelineRequest", [
            { no: 1, name: "pipeline_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "audience", kind: "message", T: () => Audience }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.AttachPipelineRequest
 */
export const AttachPipelineRequest = new AttachPipelineRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DetachPipelineRequest$Type extends MessageType {
    constructor() {
        super("protos.DetachPipelineRequest", [
            { no: 1, name: "pipeline_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "audience", kind: "message", T: () => Audience }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.DetachPipelineRequest
 */
export const DetachPipelineRequest = new DetachPipelineRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class PausePipelineRequest$Type extends MessageType {
    constructor() {
        super("protos.PausePipelineRequest", [
            { no: 1, name: "pipeline_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.PausePipelineRequest
 */
export const PausePipelineRequest = new PausePipelineRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ResumePipelineRequest$Type extends MessageType {
    constructor() {
        super("protos.ResumePipelineRequest", [
            { no: 1, name: "pipeline_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.ResumePipelineRequest
 */
export const ResumePipelineRequest = new ResumePipelineRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class TestRequest$Type extends MessageType {
    constructor() {
        super("protos.TestRequest", [
            { no: 1, name: "input", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.TestRequest
 */
export const TestRequest = new TestRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class TestResponse$Type extends MessageType {
    constructor() {
        super("protos.TestResponse", [
            { no: 2, name: "output", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.TestResponse
 */
export const TestResponse = new TestResponse$Type();
/**
 * @generated ServiceType for protobuf service protos.External
 */
export const External = new ServiceType("protos.External", [
    { name: "GetServiceMap", options: {}, I: GetServiceMapRequest, O: GetServiceMapResponse },
    { name: "GetPipelines", options: {}, I: GetPipelinesRequest, O: GetPipelinesResponse },
    { name: "GetPipeline", options: {}, I: GetPipelineRequest, O: GetPipelineResponse },
    { name: "CreatePipeline", options: {}, I: CreatePipelineRequest, O: StandardResponse },
    { name: "UpdatePipeline", options: {}, I: UpdatePipelineRequest, O: StandardResponse },
    { name: "DeletePipeline", options: {}, I: DeletePipelineRequest, O: StandardResponse },
    { name: "AttachPipeline", options: {}, I: AttachPipelineRequest, O: StandardResponse },
    { name: "DetachPipeline", options: {}, I: DetachPipelineRequest, O: StandardResponse },
    { name: "PausePipeline", options: {}, I: PausePipelineRequest, O: StandardResponse },
    { name: "ResumePipeline", options: {}, I: ResumePipelineRequest, O: StandardResponse },
    { name: "Test", options: {}, I: TestRequest, O: TestResponse }
]);
//# sourceMappingURL=external.js.map