import { WireType } from "@protobuf-ts/runtime";
import { UnknownFieldHandler } from "@protobuf-ts/runtime";
import { reflectionMergePartial } from "@protobuf-ts/runtime";
import { MESSAGE_TYPE } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
import { Audience } from "./sp_common.js";
import { AbortCondition } from "./sp_pipeline.js";
/**
 * @generated from protobuf enum protos.ExecStatus
 */
export var ExecStatus;
(function (ExecStatus) {
    /**
     * Unset status. This should never be returned by the SDK. If it does, it is
     * probably a bug (and you should file an issue)
     *
     * @generated from protobuf enum value: EXEC_STATUS_UNSET = 0;
     */
    ExecStatus[ExecStatus["UNSET"] = 0] = "UNSET";
    /**
     * Indicates that the step execution evaluated to "true"
     *
     * @generated from protobuf enum value: EXEC_STATUS_TRUE = 1;
     */
    ExecStatus[ExecStatus["TRUE"] = 1] = "TRUE";
    /**
     * Indicates that the step execution evaluated to "false"
     *
     * @generated from protobuf enum value: EXEC_STATUS_FALSE = 2;
     */
    ExecStatus[ExecStatus["FALSE"] = 2] = "FALSE";
    /**
     * Indicates that the SDK encountered an error while trying to process the
     * request. Example error cases: SDK can't find the appropriate Wasm module,
     * Wasm function cannot alloc or dealloc memory, etc.
     *
     * @generated from protobuf enum value: EXEC_STATUS_ERROR = 3;
     */
    ExecStatus[ExecStatus["ERROR"] = 3] = "ERROR";
})(ExecStatus || (ExecStatus = {}));
/**
 * ShimErrorMode is used to alter the error behavior of a shim library
 * instrumented with the Streamdal SDK at runtime.
 *
 * NOTE: This structure is usually used when the SDK is used via a shim/wrapper
 * library where you have less control over SDK behavior. Read more about shims
 * here: https://docs.streamdal.com/en/core-components/libraries-shims/
 *
 * @generated from protobuf enum protos.ShimErrorMode
 */
export var ShimErrorMode;
(function (ShimErrorMode) {
    /**
     * This instructs the shim to IGNORE any non-recoverable errors that the SDK
     * might run into. If the SDK runs into an error, the shim will NOT pass the
     * error back to the user - it will instead return the whatever the upstream
     * library would normally return to the user.
     *
     * *** This is the default behavior ***
     *
     * Example with Redis Shim
     * ------------------------
     * Under normal conditions, a Redis shim would work in the following way when
     * user is performing a read operation:
     *
     * 1. The shim would call the upstream Redis library to perform the read operation
     * 2. Upstream library returns results to the shim
     * 3. Shim passes the result to the integrated Streamdal SDK for processing
     * 4. SDK returns (potentially) modified data to the shim
     * 5. Shim returns the modified data to the user
     *
     * This setting tells the shim that IF it runs into a non-recoverable error
     * while calling the SDK (step 3), it will side-step steps 4 and 5 and instead
     * return the _original_ payload (read during step 1) to the user.
     *
     * @generated from protobuf enum value: SHIM_ERROR_MODE_UNSET = 0;
     */
    ShimErrorMode[ShimErrorMode["UNSET"] = 0] = "UNSET";
    /**
     * This instructs the shim to ABORT execution if the SDK runs into any
     * non-recoverable errors. Upon aborting, the shim will return the error that
     * the SDK ran into and the error will be passed all the way back to the user.
     *
     * @generated from protobuf enum value: SHIM_ERROR_MODE_STRICT = 1;
     */
    ShimErrorMode[ShimErrorMode["STRICT"] = 1] = "STRICT";
})(ShimErrorMode || (ShimErrorMode = {}));
// @generated message type with reflection information, may provide speed optimized methods
class SDKResponse$Type extends MessageType {
    constructor() {
        super("protos.SDKResponse", [
            { no: 1, name: "data", kind: "scalar", T: 12 /*ScalarType.BYTES*/ },
            { no: 2, name: "status", kind: "enum", T: () => ["protos.ExecStatus", ExecStatus, "EXEC_STATUS_"] },
            { no: 3, name: "status_message", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "pipeline_status", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => PipelineStatus },
            { no: 5, name: "metadata", kind: "map", K: 9 /*ScalarType.STRING*/, V: { kind: "scalar", T: 9 /*ScalarType.STRING*/ } }
        ]);
    }
    create(value) {
        const message = { data: new Uint8Array(0), status: 0, pipelineStatus: [], metadata: {} };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial(this, message, value);
        return message;
    }
    internalBinaryRead(reader, length, options, target) {
        let message = target !== null && target !== void 0 ? target : this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* bytes data */ 1:
                    message.data = reader.bytes();
                    break;
                case /* protos.ExecStatus status */ 2:
                    message.status = reader.int32();
                    break;
                case /* optional string status_message */ 3:
                    message.statusMessage = reader.string();
                    break;
                case /* repeated protos.PipelineStatus pipeline_status */ 4:
                    message.pipelineStatus.push(PipelineStatus.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                case /* map<string, string> metadata */ 5:
                    this.binaryReadMap5(message.metadata, reader, options);
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    binaryReadMap5(map, reader, options) {
        let len = reader.uint32(), end = reader.pos + len, key, val;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case 1:
                    key = reader.string();
                    break;
                case 2:
                    val = reader.string();
                    break;
                default: throw new globalThis.Error("unknown map entry field for field protos.SDKResponse.metadata");
            }
        }
        map[key !== null && key !== void 0 ? key : ""] = val !== null && val !== void 0 ? val : "";
    }
    internalBinaryWrite(message, writer, options) {
        /* bytes data = 1; */
        if (message.data.length)
            writer.tag(1, WireType.LengthDelimited).bytes(message.data);
        /* protos.ExecStatus status = 2; */
        if (message.status !== 0)
            writer.tag(2, WireType.Varint).int32(message.status);
        /* optional string status_message = 3; */
        if (message.statusMessage !== undefined)
            writer.tag(3, WireType.LengthDelimited).string(message.statusMessage);
        /* repeated protos.PipelineStatus pipeline_status = 4; */
        for (let i = 0; i < message.pipelineStatus.length; i++)
            PipelineStatus.internalBinaryWrite(message.pipelineStatus[i], writer.tag(4, WireType.LengthDelimited).fork(), options).join();
        /* map<string, string> metadata = 5; */
        for (let k of Object.keys(message.metadata))
            writer.tag(5, WireType.LengthDelimited).fork().tag(1, WireType.LengthDelimited).string(k).tag(2, WireType.LengthDelimited).string(message.metadata[k]).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message protos.SDKResponse
 */
export const SDKResponse = new SDKResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class PipelineStatus$Type extends MessageType {
    constructor() {
        super("protos.PipelineStatus", [
            { no: 1, name: "id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "name", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "step_status", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => StepStatus }
        ]);
    }
    create(value) {
        const message = { id: "", name: "", stepStatus: [] };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial(this, message, value);
        return message;
    }
    internalBinaryRead(reader, length, options, target) {
        let message = target !== null && target !== void 0 ? target : this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string id */ 1:
                    message.id = reader.string();
                    break;
                case /* string name */ 2:
                    message.name = reader.string();
                    break;
                case /* repeated protos.StepStatus step_status */ 3:
                    message.stepStatus.push(StepStatus.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message, writer, options) {
        /* string id = 1; */
        if (message.id !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.id);
        /* string name = 2; */
        if (message.name !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.name);
        /* repeated protos.StepStatus step_status = 3; */
        for (let i = 0; i < message.stepStatus.length; i++)
            StepStatus.internalBinaryWrite(message.stepStatus[i], writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message protos.PipelineStatus
 */
export const PipelineStatus = new PipelineStatus$Type();
// @generated message type with reflection information, may provide speed optimized methods
class StepStatus$Type extends MessageType {
    constructor() {
        super("protos.StepStatus", [
            { no: 1, name: "name", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "status", kind: "enum", T: () => ["protos.ExecStatus", ExecStatus, "EXEC_STATUS_"] },
            { no: 3, name: "status_message", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "abort_condition", kind: "enum", T: () => ["protos.AbortCondition", AbortCondition, "ABORT_CONDITION_"] }
        ]);
    }
    create(value) {
        const message = { name: "", status: 0, abortCondition: 0 };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial(this, message, value);
        return message;
    }
    internalBinaryRead(reader, length, options, target) {
        let message = target !== null && target !== void 0 ? target : this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string name */ 1:
                    message.name = reader.string();
                    break;
                case /* protos.ExecStatus status */ 2:
                    message.status = reader.int32();
                    break;
                case /* optional string status_message */ 3:
                    message.statusMessage = reader.string();
                    break;
                case /* protos.AbortCondition abort_condition */ 4:
                    message.abortCondition = reader.int32();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message, writer, options) {
        /* string name = 1; */
        if (message.name !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.name);
        /* protos.ExecStatus status = 2; */
        if (message.status !== 0)
            writer.tag(2, WireType.Varint).int32(message.status);
        /* optional string status_message = 3; */
        if (message.statusMessage !== undefined)
            writer.tag(3, WireType.LengthDelimited).string(message.statusMessage);
        /* protos.AbortCondition abort_condition = 4; */
        if (message.abortCondition !== 0)
            writer.tag(4, WireType.Varint).int32(message.abortCondition);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message protos.StepStatus
 */
export const StepStatus = new StepStatus$Type();
// @generated message type with reflection information, may provide speed optimized methods
class SDKStartupConfig$Type extends MessageType {
    constructor() {
        super("protos.SDKStartupConfig", [
            { no: 1, name: "url", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "token", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "service_name", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "pipeline_timeout_seconds", kind: "scalar", opt: true, T: 5 /*ScalarType.INT32*/ },
            { no: 5, name: "step_timeout_seconds", kind: "scalar", opt: true, T: 5 /*ScalarType.INT32*/ },
            { no: 6, name: "dry_run", kind: "scalar", opt: true, T: 8 /*ScalarType.BOOL*/ },
            { no: 1000, name: "shim_require_runtime_config", kind: "scalar", opt: true, T: 8 /*ScalarType.BOOL*/ },
            { no: 1001, name: "shim_error_mode", kind: "enum", opt: true, T: () => ["protos.ShimErrorMode", ShimErrorMode, "SHIM_ERROR_MODE_"] }
        ]);
    }
    create(value) {
        const message = { url: "", token: "", serviceName: "" };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial(this, message, value);
        return message;
    }
    internalBinaryRead(reader, length, options, target) {
        let message = target !== null && target !== void 0 ? target : this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string url */ 1:
                    message.url = reader.string();
                    break;
                case /* string token */ 2:
                    message.token = reader.string();
                    break;
                case /* string service_name */ 3:
                    message.serviceName = reader.string();
                    break;
                case /* optional int32 pipeline_timeout_seconds */ 4:
                    message.pipelineTimeoutSeconds = reader.int32();
                    break;
                case /* optional int32 step_timeout_seconds */ 5:
                    message.stepTimeoutSeconds = reader.int32();
                    break;
                case /* optional bool dry_run */ 6:
                    message.dryRun = reader.bool();
                    break;
                case /* optional bool shim_require_runtime_config */ 1000:
                    message.shimRequireRuntimeConfig = reader.bool();
                    break;
                case /* optional protos.ShimErrorMode shim_error_mode */ 1001:
                    message.shimErrorMode = reader.int32();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message, writer, options) {
        /* string url = 1; */
        if (message.url !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.url);
        /* string token = 2; */
        if (message.token !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.token);
        /* string service_name = 3; */
        if (message.serviceName !== "")
            writer.tag(3, WireType.LengthDelimited).string(message.serviceName);
        /* optional int32 pipeline_timeout_seconds = 4; */
        if (message.pipelineTimeoutSeconds !== undefined)
            writer.tag(4, WireType.Varint).int32(message.pipelineTimeoutSeconds);
        /* optional int32 step_timeout_seconds = 5; */
        if (message.stepTimeoutSeconds !== undefined)
            writer.tag(5, WireType.Varint).int32(message.stepTimeoutSeconds);
        /* optional bool dry_run = 6; */
        if (message.dryRun !== undefined)
            writer.tag(6, WireType.Varint).bool(message.dryRun);
        /* optional bool shim_require_runtime_config = 1000; */
        if (message.shimRequireRuntimeConfig !== undefined)
            writer.tag(1000, WireType.Varint).bool(message.shimRequireRuntimeConfig);
        /* optional protos.ShimErrorMode shim_error_mode = 1001; */
        if (message.shimErrorMode !== undefined)
            writer.tag(1001, WireType.Varint).int32(message.shimErrorMode);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message protos.SDKStartupConfig
 */
export const SDKStartupConfig = new SDKStartupConfig$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ShimRuntimeConfig$Type extends MessageType {
    constructor() {
        super("protos.ShimRuntimeConfig", [
            { no: 1, name: "error_mode", kind: "enum", opt: true, T: () => ["protos.ShimErrorMode", ShimErrorMode, "SHIM_ERROR_MODE_"] },
            { no: 2, name: "audience", kind: "message", T: () => Audience }
        ]);
    }
    create(value) {
        const message = {};
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial(this, message, value);
        return message;
    }
    internalBinaryRead(reader, length, options, target) {
        let message = target !== null && target !== void 0 ? target : this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* optional protos.ShimErrorMode error_mode */ 1:
                    message.errorMode = reader.int32();
                    break;
                case /* optional protos.Audience audience */ 2:
                    message.audience = Audience.internalBinaryRead(reader, reader.uint32(), options, message.audience);
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message, writer, options) {
        /* optional protos.ShimErrorMode error_mode = 1; */
        if (message.errorMode !== undefined)
            writer.tag(1, WireType.Varint).int32(message.errorMode);
        /* optional protos.Audience audience = 2; */
        if (message.audience)
            Audience.internalBinaryWrite(message.audience, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message protos.ShimRuntimeConfig
 */
export const ShimRuntimeConfig = new ShimRuntimeConfig$Type();
