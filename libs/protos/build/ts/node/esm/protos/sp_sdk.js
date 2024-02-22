import { WireType } from "@protobuf-ts/runtime";
import { UnknownFieldHandler } from "@protobuf-ts/runtime";
import { reflectionMergePartial } from "@protobuf-ts/runtime";
import { MESSAGE_TYPE } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
import { AbortCondition } from "./sp_pipeline.js";
import { Audience } from "./sp_common.js";
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
 * Indicates whether the SDK is being used directly or via a shim/wrapper library.
 * This is primarily intended to be used by shims so that the SDK can determine
 * if the ServerURL and ServerToken should be optional or required.
 * protolint:disable ENUM_FIELD_NAMES_PREFIX
 *
 * @generated from protobuf enum protos.SDKClientType
 */
export var SDKClientType;
(function (SDKClientType) {
    /**
     * The SDK is used directly as a standalone library
     *
     * @generated from protobuf enum value: SDK_CLIENT_TYPE_DIRECT = 0;
     */
    SDKClientType[SDKClientType["SDK_CLIENT_TYPE_DIRECT"] = 0] = "SDK_CLIENT_TYPE_DIRECT";
    /**
     * The SDK is used within a shim/wrapper library
     *
     * @generated from protobuf enum value: SDK_CLIENT_TYPE_SHIM = 1;
     */
    SDKClientType[SDKClientType["SDK_CLIENT_TYPE_SHIM"] = 1] = "SDK_CLIENT_TYPE_SHIM";
})(SDKClientType || (SDKClientType = {}));
// @generated message type with reflection information, may provide speed optimized methods
class SDKRequest$Type extends MessageType {
    constructor() {
        super("protos.SDKRequest", [
            { no: 1, name: "data", kind: "scalar", T: 12 /*ScalarType.BYTES*/ },
            { no: 2, name: "audience", kind: "message", T: () => Audience }
        ]);
    }
    create(value) {
        const message = { data: new Uint8Array(0) };
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
                case /* protos.Audience audience */ 2:
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
        /* bytes data = 1; */
        if (message.data.length)
            writer.tag(1, WireType.LengthDelimited).bytes(message.data);
        /* protos.Audience audience = 2; */
        if (message.audience)
            Audience.internalBinaryWrite(message.audience, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message protos.SDKRequest
 */
export const SDKRequest = new SDKRequest$Type();
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
class SDKStartupConfig$Type extends MessageType {
    constructor() {
        super("protos.SDKStartupConfig", [
            { no: 1, name: "server_url", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "auth_token", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "service_name", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "audiences", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Audience },
            { no: 5, name: "pipeline_timeout_seconds", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 6, name: "step_timeout_seconds", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 7, name: "dry_run", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 1000, name: "_internal_client_type", kind: "enum", opt: true, T: () => ["protos.SDKClientType", SDKClientType] },
            { no: 2000, name: "_internal_shim_require_runtime_config", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 2001, name: "_internal_shim_strict_error_handling", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
    create(value) {
        const message = { serverUrl: "", authToken: "", serviceName: "", audiences: [], pipelineTimeoutSeconds: 0, stepTimeoutSeconds: 0, dryRun: false, InternalShimRequireRuntimeConfig: false, InternalShimStrictErrorHandling: false };
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
                case /* string server_url */ 1:
                    message.serverUrl = reader.string();
                    break;
                case /* string auth_token */ 2:
                    message.authToken = reader.string();
                    break;
                case /* string service_name */ 3:
                    message.serviceName = reader.string();
                    break;
                case /* repeated protos.Audience audiences */ 4:
                    message.audiences.push(Audience.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                case /* int32 pipeline_timeout_seconds */ 5:
                    message.pipelineTimeoutSeconds = reader.int32();
                    break;
                case /* int32 step_timeout_seconds */ 6:
                    message.stepTimeoutSeconds = reader.int32();
                    break;
                case /* bool dry_run */ 7:
                    message.dryRun = reader.bool();
                    break;
                case /* optional protos.SDKClientType _internal_client_type */ 1000:
                    message.InternalClientType = reader.int32();
                    break;
                case /* bool _internal_shim_require_runtime_config */ 2000:
                    message.InternalShimRequireRuntimeConfig = reader.bool();
                    break;
                case /* bool _internal_shim_strict_error_handling */ 2001:
                    message.InternalShimStrictErrorHandling = reader.bool();
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
        /* string server_url = 1; */
        if (message.serverUrl !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.serverUrl);
        /* string auth_token = 2; */
        if (message.authToken !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.authToken);
        /* string service_name = 3; */
        if (message.serviceName !== "")
            writer.tag(3, WireType.LengthDelimited).string(message.serviceName);
        /* repeated protos.Audience audiences = 4; */
        for (let i = 0; i < message.audiences.length; i++)
            Audience.internalBinaryWrite(message.audiences[i], writer.tag(4, WireType.LengthDelimited).fork(), options).join();
        /* int32 pipeline_timeout_seconds = 5; */
        if (message.pipelineTimeoutSeconds !== 0)
            writer.tag(5, WireType.Varint).int32(message.pipelineTimeoutSeconds);
        /* int32 step_timeout_seconds = 6; */
        if (message.stepTimeoutSeconds !== 0)
            writer.tag(6, WireType.Varint).int32(message.stepTimeoutSeconds);
        /* bool dry_run = 7; */
        if (message.dryRun !== false)
            writer.tag(7, WireType.Varint).bool(message.dryRun);
        /* optional protos.SDKClientType _internal_client_type = 1000; */
        if (message.InternalClientType !== undefined)
            writer.tag(1000, WireType.Varint).int32(message.InternalClientType);
        /* bool _internal_shim_require_runtime_config = 2000; */
        if (message.InternalShimRequireRuntimeConfig !== false)
            writer.tag(2000, WireType.Varint).bool(message.InternalShimRequireRuntimeConfig);
        /* bool _internal_shim_strict_error_handling = 2001; */
        if (message.InternalShimStrictErrorHandling !== false)
            writer.tag(2001, WireType.Varint).bool(message.InternalShimStrictErrorHandling);
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
class SDKRuntimeConfig$Type extends MessageType {
    constructor() {
        super("protos.SDKRuntimeConfig", [
            { no: 1, name: "audience", kind: "message", T: () => Audience },
            { no: 2, name: "strict_error_handling", kind: "scalar", opt: true, T: 8 /*ScalarType.BOOL*/ }
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
                case /* protos.Audience audience */ 1:
                    message.audience = Audience.internalBinaryRead(reader, reader.uint32(), options, message.audience);
                    break;
                case /* optional bool strict_error_handling */ 2:
                    message.strictErrorHandling = reader.bool();
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
        /* protos.Audience audience = 1; */
        if (message.audience)
            Audience.internalBinaryWrite(message.audience, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* optional bool strict_error_handling = 2; */
        if (message.strictErrorHandling !== undefined)
            writer.tag(2, WireType.Varint).bool(message.strictErrorHandling);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message protos.SDKRuntimeConfig
 */
export const SDKRuntimeConfig = new SDKRuntimeConfig$Type();
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
