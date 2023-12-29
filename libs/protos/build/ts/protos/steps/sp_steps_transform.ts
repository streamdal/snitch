// @generated by protobuf-ts 2.9.0 with parameter long_type_string
// @generated from protobuf file "steps/sp_steps_transform.proto" (package "protos.steps", syntax proto3)
// tslint:disable
import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import { WireType } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import { UnknownFieldHandler } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { reflectionMergePartial } from "@protobuf-ts/runtime";
import { MESSAGE_TYPE } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
/**
 * @generated from protobuf message protos.steps.TransformStep
 */
export interface TransformStep {
    /**
     * @deprecated
     * @generated from protobuf field: string path = 1 [deprecated = true];
     */
    path: string;
    /**
     * @deprecated
     * @generated from protobuf field: string value = 2 [deprecated = true];
     */
    value: string; // Should this be bytes? ~DS
    /**
     * @generated from protobuf field: protos.steps.TransformType type = 3;
     */
    type: TransformType;
    /**
     * @generated from protobuf oneof: options
     */
    options: {
        oneofKind: "replaceValueOptions";
        /**
         * @generated from protobuf field: protos.steps.TransformReplaceValueStep replace_value_options = 101;
         */
        replaceValueOptions: TransformReplaceValueStep;
    } | {
        oneofKind: "deleteFieldOptions";
        /**
         * @generated from protobuf field: protos.steps.TransformDeleteFieldStep delete_field_options = 102;
         */
        deleteFieldOptions: TransformDeleteFieldStep;
    } | {
        oneofKind: "obfuscateOptions";
        /**
         * @generated from protobuf field: protos.steps.TransformObfuscateOptions obfuscate_options = 103;
         */
        obfuscateOptions: TransformObfuscateOptions;
    } | {
        oneofKind: "maskOptions";
        /**
         * @generated from protobuf field: protos.steps.TransformMaskOptions mask_options = 104;
         */
        maskOptions: TransformMaskOptions;
    } | {
        oneofKind: "truncateOptions";
        /**
         * @generated from protobuf field: protos.steps.TransformTruncateOptions truncate_options = 105;
         */
        truncateOptions: TransformTruncateOptions;
    } | {
        oneofKind: undefined;
    };
}
/**
 * @generated from protobuf message protos.steps.TransformTruncateOptions
 */
export interface TransformTruncateOptions {
    /**
     * @generated from protobuf field: protos.steps.TransformTruncateType type = 1;
     */
    type: TransformTruncateType;
    /**
     * Truncate after this many bytes or this percentage of the original value
     *
     * @generated from protobuf field: int32 value = 2;
     */
    value: number;
}
/**
 * @generated from protobuf message protos.steps.TransformDeleteFieldStep
 */
export interface TransformDeleteFieldStep {
    /**
     * @generated from protobuf field: string path = 1;
     */
    path: string;
}
/**
 * @generated from protobuf message protos.steps.TransformReplaceValueStep
 */
export interface TransformReplaceValueStep {
    /**
     * @generated from protobuf field: string path = 1;
     */
    path: string;
    /**
     * @generated from protobuf field: string value = 2;
     */
    value: string;
}
/**
 * @generated from protobuf message protos.steps.TransformObfuscateOptions
 */
export interface TransformObfuscateOptions {
    /**
     * @generated from protobuf field: string path = 1;
     */
    path: string;
    /**
     * @generated from protobuf field: string value = 2;
     */
    value: string;
}
/**
 * @generated from protobuf message protos.steps.TransformMaskOptions
 */
export interface TransformMaskOptions {
    /**
     * @generated from protobuf field: string path = 1;
     */
    path: string;
    /**
     * @generated from protobuf field: string value = 2;
     */
    value: string;
    /**
     * @generated from protobuf field: string mask = 3;
     */
    mask: string;
}
/**
 * @generated from protobuf enum protos.steps.TransformType
 */
export enum TransformType {
    /**
     * @generated from protobuf enum value: TRANSFORM_TYPE_UNKNOWN = 0;
     */
    UNKNOWN = 0,
    /**
     * @generated from protobuf enum value: TRANSFORM_TYPE_REPLACE_VALUE = 1;
     */
    REPLACE_VALUE = 1,
    /**
     * @generated from protobuf enum value: TRANSFORM_TYPE_DELETE_FIELD = 2;
     */
    DELETE_FIELD = 2,
    /**
     * @generated from protobuf enum value: TRANSFORM_TYPE_OBFUSCATE_VALUE = 3;
     */
    OBFUSCATE_VALUE = 3,
    /**
     * @generated from protobuf enum value: TRANSFORM_TYPE_MASK_VALUE = 4;
     */
    MASK_VALUE = 4,
    /**
     * TODO: type for delete all keys except specified ones
     *
     * @generated from protobuf enum value: TRANSFORM_TYPE_TRUNCATE_VALUE = 5;
     */
    TRUNCATE_VALUE = 5
}
/**
 * @generated from protobuf enum protos.steps.TransformTruncateType
 */
export enum TransformTruncateType {
    /**
     * @generated from protobuf enum value: TRANSFORM_TRUNCATE_TYPE_UNKNOWN = 0;
     */
    UNKNOWN = 0,
    /**
     * @generated from protobuf enum value: TRANSFORM_TRUNCATE_TYPE_LENGTH = 1;
     */
    LENGTH = 1,
    /**
     * @generated from protobuf enum value: TRANSFORM_TRUNCATE_TYPE_PERCENTAGE = 2;
     */
    PERCENTAGE = 2
}
// @generated message type with reflection information, may provide speed optimized methods
class TransformStep$Type extends MessageType<TransformStep> {
    constructor() {
        super("protos.steps.TransformStep", [
            { no: 1, name: "path", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "value", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "type", kind: "enum", T: () => ["protos.steps.TransformType", TransformType, "TRANSFORM_TYPE_"] },
            { no: 101, name: "replace_value_options", kind: "message", oneof: "options", T: () => TransformReplaceValueStep },
            { no: 102, name: "delete_field_options", kind: "message", oneof: "options", T: () => TransformDeleteFieldStep },
            { no: 103, name: "obfuscate_options", kind: "message", oneof: "options", T: () => TransformObfuscateOptions },
            { no: 104, name: "mask_options", kind: "message", oneof: "options", T: () => TransformMaskOptions },
            { no: 105, name: "truncate_options", kind: "message", oneof: "options", T: () => TransformTruncateOptions }
        ]);
    }
    create(value?: PartialMessage<TransformStep>): TransformStep {
        const message = { path: "", value: "", type: 0, options: { oneofKind: undefined } };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<TransformStep>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: TransformStep): TransformStep {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string path = 1 [deprecated = true];*/ 1:
                    message.path = reader.string();
                    break;
                case /* string value = 2 [deprecated = true];*/ 2:
                    message.value = reader.string();
                    break;
                case /* protos.steps.TransformType type */ 3:
                    message.type = reader.int32();
                    break;
                case /* protos.steps.TransformReplaceValueStep replace_value_options */ 101:
                    message.options = {
                        oneofKind: "replaceValueOptions",
                        replaceValueOptions: TransformReplaceValueStep.internalBinaryRead(reader, reader.uint32(), options, (message.options as any).replaceValueOptions)
                    };
                    break;
                case /* protos.steps.TransformDeleteFieldStep delete_field_options */ 102:
                    message.options = {
                        oneofKind: "deleteFieldOptions",
                        deleteFieldOptions: TransformDeleteFieldStep.internalBinaryRead(reader, reader.uint32(), options, (message.options as any).deleteFieldOptions)
                    };
                    break;
                case /* protos.steps.TransformObfuscateOptions obfuscate_options */ 103:
                    message.options = {
                        oneofKind: "obfuscateOptions",
                        obfuscateOptions: TransformObfuscateOptions.internalBinaryRead(reader, reader.uint32(), options, (message.options as any).obfuscateOptions)
                    };
                    break;
                case /* protos.steps.TransformMaskOptions mask_options */ 104:
                    message.options = {
                        oneofKind: "maskOptions",
                        maskOptions: TransformMaskOptions.internalBinaryRead(reader, reader.uint32(), options, (message.options as any).maskOptions)
                    };
                    break;
                case /* protos.steps.TransformTruncateOptions truncate_options */ 105:
                    message.options = {
                        oneofKind: "truncateOptions",
                        truncateOptions: TransformTruncateOptions.internalBinaryRead(reader, reader.uint32(), options, (message.options as any).truncateOptions)
                    };
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
    internalBinaryWrite(message: TransformStep, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string path = 1 [deprecated = true]; */
        if (message.path !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.path);
        /* string value = 2 [deprecated = true]; */
        if (message.value !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.value);
        /* protos.steps.TransformType type = 3; */
        if (message.type !== 0)
            writer.tag(3, WireType.Varint).int32(message.type);
        /* protos.steps.TransformReplaceValueStep replace_value_options = 101; */
        if (message.options.oneofKind === "replaceValueOptions")
            TransformReplaceValueStep.internalBinaryWrite(message.options.replaceValueOptions, writer.tag(101, WireType.LengthDelimited).fork(), options).join();
        /* protos.steps.TransformDeleteFieldStep delete_field_options = 102; */
        if (message.options.oneofKind === "deleteFieldOptions")
            TransformDeleteFieldStep.internalBinaryWrite(message.options.deleteFieldOptions, writer.tag(102, WireType.LengthDelimited).fork(), options).join();
        /* protos.steps.TransformObfuscateOptions obfuscate_options = 103; */
        if (message.options.oneofKind === "obfuscateOptions")
            TransformObfuscateOptions.internalBinaryWrite(message.options.obfuscateOptions, writer.tag(103, WireType.LengthDelimited).fork(), options).join();
        /* protos.steps.TransformMaskOptions mask_options = 104; */
        if (message.options.oneofKind === "maskOptions")
            TransformMaskOptions.internalBinaryWrite(message.options.maskOptions, writer.tag(104, WireType.LengthDelimited).fork(), options).join();
        /* protos.steps.TransformTruncateOptions truncate_options = 105; */
        if (message.options.oneofKind === "truncateOptions")
            TransformTruncateOptions.internalBinaryWrite(message.options.truncateOptions, writer.tag(105, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message protos.steps.TransformStep
 */
export const TransformStep = new TransformStep$Type();
// @generated message type with reflection information, may provide speed optimized methods
class TransformTruncateOptions$Type extends MessageType<TransformTruncateOptions> {
    constructor() {
        super("protos.steps.TransformTruncateOptions", [
            { no: 1, name: "type", kind: "enum", T: () => ["protos.steps.TransformTruncateType", TransformTruncateType, "TRANSFORM_TRUNCATE_TYPE_"] },
            { no: 2, name: "value", kind: "scalar", T: 5 /*ScalarType.INT32*/ }
        ]);
    }
    create(value?: PartialMessage<TransformTruncateOptions>): TransformTruncateOptions {
        const message = { type: 0, value: 0 };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<TransformTruncateOptions>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: TransformTruncateOptions): TransformTruncateOptions {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* protos.steps.TransformTruncateType type */ 1:
                    message.type = reader.int32();
                    break;
                case /* int32 value */ 2:
                    message.value = reader.int32();
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
    internalBinaryWrite(message: TransformTruncateOptions, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* protos.steps.TransformTruncateType type = 1; */
        if (message.type !== 0)
            writer.tag(1, WireType.Varint).int32(message.type);
        /* int32 value = 2; */
        if (message.value !== 0)
            writer.tag(2, WireType.Varint).int32(message.value);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message protos.steps.TransformTruncateOptions
 */
export const TransformTruncateOptions = new TransformTruncateOptions$Type();
// @generated message type with reflection information, may provide speed optimized methods
class TransformDeleteFieldStep$Type extends MessageType<TransformDeleteFieldStep> {
    constructor() {
        super("protos.steps.TransformDeleteFieldStep", [
            { no: 1, name: "path", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<TransformDeleteFieldStep>): TransformDeleteFieldStep {
        const message = { path: "" };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<TransformDeleteFieldStep>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: TransformDeleteFieldStep): TransformDeleteFieldStep {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string path */ 1:
                    message.path = reader.string();
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
    internalBinaryWrite(message: TransformDeleteFieldStep, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string path = 1; */
        if (message.path !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.path);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message protos.steps.TransformDeleteFieldStep
 */
export const TransformDeleteFieldStep = new TransformDeleteFieldStep$Type();
// @generated message type with reflection information, may provide speed optimized methods
class TransformReplaceValueStep$Type extends MessageType<TransformReplaceValueStep> {
    constructor() {
        super("protos.steps.TransformReplaceValueStep", [
            { no: 1, name: "path", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "value", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<TransformReplaceValueStep>): TransformReplaceValueStep {
        const message = { path: "", value: "" };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<TransformReplaceValueStep>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: TransformReplaceValueStep): TransformReplaceValueStep {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string path */ 1:
                    message.path = reader.string();
                    break;
                case /* string value */ 2:
                    message.value = reader.string();
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
    internalBinaryWrite(message: TransformReplaceValueStep, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string path = 1; */
        if (message.path !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.path);
        /* string value = 2; */
        if (message.value !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.value);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message protos.steps.TransformReplaceValueStep
 */
export const TransformReplaceValueStep = new TransformReplaceValueStep$Type();
// @generated message type with reflection information, may provide speed optimized methods
class TransformObfuscateOptions$Type extends MessageType<TransformObfuscateOptions> {
    constructor() {
        super("protos.steps.TransformObfuscateOptions", [
            { no: 1, name: "path", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "value", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<TransformObfuscateOptions>): TransformObfuscateOptions {
        const message = { path: "", value: "" };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<TransformObfuscateOptions>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: TransformObfuscateOptions): TransformObfuscateOptions {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string path */ 1:
                    message.path = reader.string();
                    break;
                case /* string value */ 2:
                    message.value = reader.string();
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
    internalBinaryWrite(message: TransformObfuscateOptions, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string path = 1; */
        if (message.path !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.path);
        /* string value = 2; */
        if (message.value !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.value);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message protos.steps.TransformObfuscateOptions
 */
export const TransformObfuscateOptions = new TransformObfuscateOptions$Type();
// @generated message type with reflection information, may provide speed optimized methods
class TransformMaskOptions$Type extends MessageType<TransformMaskOptions> {
    constructor() {
        super("protos.steps.TransformMaskOptions", [
            { no: 1, name: "path", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "value", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "mask", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<TransformMaskOptions>): TransformMaskOptions {
        const message = { path: "", value: "", mask: "" };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<TransformMaskOptions>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: TransformMaskOptions): TransformMaskOptions {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string path */ 1:
                    message.path = reader.string();
                    break;
                case /* string value */ 2:
                    message.value = reader.string();
                    break;
                case /* string mask */ 3:
                    message.mask = reader.string();
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
    internalBinaryWrite(message: TransformMaskOptions, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string path = 1; */
        if (message.path !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.path);
        /* string value = 2; */
        if (message.value !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.value);
        /* string mask = 3; */
        if (message.mask !== "")
            writer.tag(3, WireType.LengthDelimited).string(message.mask);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message protos.steps.TransformMaskOptions
 */
export const TransformMaskOptions = new TransformMaskOptions$Type();
