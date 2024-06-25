// @generated by protobuf-ts 2.9.0 with parameter long_type_string
// @generated from protobuf file "sp_command.proto" (package "protos", syntax proto3)
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
import { TailRequest } from "./sp_common.ts";
import { KVInstruction } from "./sp_kv.ts";
import { WasmModule } from "./shared/sp_shared.ts";
import { Pipeline } from "./sp_pipeline.ts";
import { Audience } from "./sp_common.ts";
/**
 * Command is used by streamdal server for sending commands to SDKs
 *
 * @generated from protobuf message protos.Command
 */
export interface Command {
    /**
     * Who is this command intended for?
     * NOTE: Some commands (such as KeepAliveCommand, KVCommand) do NOT use audience and will ignore it
     *
     * @generated from protobuf field: protos.Audience audience = 1;
     */
    audience?: Audience;
    /**
     * @generated from protobuf oneof: command
     */
    command: {
        oneofKind: "setPipelines";
        /**
         * Emitted by server when a user makes a pause, resume, delete or update
         * pipeline and set pipelines external grpc API call.
         * NOTE: This was introduced during ordered pipeline updates.
         *
         * @generated from protobuf field: protos.SetPipelinesCommand set_pipelines = 100;
         */
        setPipelines: SetPipelinesCommand;
    } | {
        oneofKind: "keepAlive";
        /**
         * Server sends this periodically to SDKs to keep the connection alive
         *
         * @generated from protobuf field: protos.KeepAliveCommand keep_alive = 101;
         */
        keepAlive: KeepAliveCommand;
    } | {
        oneofKind: "kv";
        /**
         * Server will emit this when a user makes changes to the KV store
         * via the KV HTTP API.
         *
         * @generated from protobuf field: protos.KVCommand kv = 102;
         */
        kv: KVCommand;
    } | {
        oneofKind: "tail";
        /**
         * Emitted by server when a user makes a Tail() call
         * Consumed by all server instances and by SDKs
         *
         * @generated from protobuf field: protos.TailCommand tail = 103;
         */
        tail: TailCommand;
    } | {
        oneofKind: "delete";
        /**
         * @generated from protobuf field: protos.DeleteAudienceCommand delete = 104;
         */
        delete: DeleteAudienceCommand;
    } | {
        oneofKind: undefined;
    };
}
/**
 * @generated from protobuf message protos.SetPipelinesCommand
 */
export interface SetPipelinesCommand {
    /**
     * @generated from protobuf field: repeated protos.Pipeline pipelines = 1;
     */
    pipelines: Pipeline[];
    /**
     * ID = wasm ID
     *
     * @generated from protobuf field: map<string, protos.shared.WasmModule> wasm_modules = 2;
     */
    wasmModules: {
        [key: string]: WasmModule;
    };
}
/**
 * Nothing needed in here, just a ping from server to SDK
 *
 * @generated from protobuf message protos.KeepAliveCommand
 */
export interface KeepAliveCommand {
}
/**
 * Sent by server on Register channel(s) to live SDKs
 *
 * @generated from protobuf message protos.KVCommand
 */
export interface KVCommand {
    /**
     * @generated from protobuf field: repeated protos.KVInstruction instructions = 1;
     */
    instructions: KVInstruction[];
    /**
     * Create & Update specific setting that will cause the Create or Update to
     * work as an upsert.
     *
     * @generated from protobuf field: bool overwrite = 2;
     */
    overwrite: boolean;
}
/**
 * @generated from protobuf message protos.TailCommand
 */
export interface TailCommand {
    /**
     * @generated from protobuf field: protos.TailRequest request = 2;
     */
    request?: TailRequest;
}
/**
 * @generated from protobuf message protos.DeleteAudienceCommand
 */
export interface DeleteAudienceCommand {
    /**
     * @generated from protobuf field: protos.Audience audience = 1;
     */
    audience?: Audience;
}
// @generated message type with reflection information, may provide speed optimized methods
class Command$Type extends MessageType<Command> {
    constructor() {
        super("protos.Command", [
            { no: 1, name: "audience", kind: "message", T: () => Audience },
            { no: 100, name: "set_pipelines", kind: "message", oneof: "command", T: () => SetPipelinesCommand },
            { no: 101, name: "keep_alive", kind: "message", oneof: "command", T: () => KeepAliveCommand },
            { no: 102, name: "kv", kind: "message", oneof: "command", T: () => KVCommand },
            { no: 103, name: "tail", kind: "message", oneof: "command", T: () => TailCommand },
            { no: 104, name: "delete", kind: "message", oneof: "command", T: () => DeleteAudienceCommand }
        ]);
    }
    create(value?: PartialMessage<Command>): Command {
        const message = { command: { oneofKind: undefined } };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<Command>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Command): Command {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* protos.Audience audience */ 1:
                    message.audience = Audience.internalBinaryRead(reader, reader.uint32(), options, message.audience);
                    break;
                case /* protos.SetPipelinesCommand set_pipelines */ 100:
                    message.command = {
                        oneofKind: "setPipelines",
                        setPipelines: SetPipelinesCommand.internalBinaryRead(reader, reader.uint32(), options, (message.command as any).setPipelines)
                    };
                    break;
                case /* protos.KeepAliveCommand keep_alive */ 101:
                    message.command = {
                        oneofKind: "keepAlive",
                        keepAlive: KeepAliveCommand.internalBinaryRead(reader, reader.uint32(), options, (message.command as any).keepAlive)
                    };
                    break;
                case /* protos.KVCommand kv */ 102:
                    message.command = {
                        oneofKind: "kv",
                        kv: KVCommand.internalBinaryRead(reader, reader.uint32(), options, (message.command as any).kv)
                    };
                    break;
                case /* protos.TailCommand tail */ 103:
                    message.command = {
                        oneofKind: "tail",
                        tail: TailCommand.internalBinaryRead(reader, reader.uint32(), options, (message.command as any).tail)
                    };
                    break;
                case /* protos.DeleteAudienceCommand delete */ 104:
                    message.command = {
                        oneofKind: "delete",
                        delete: DeleteAudienceCommand.internalBinaryRead(reader, reader.uint32(), options, (message.command as any).delete)
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
    internalBinaryWrite(message: Command, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* protos.Audience audience = 1; */
        if (message.audience)
            Audience.internalBinaryWrite(message.audience, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* protos.SetPipelinesCommand set_pipelines = 100; */
        if (message.command.oneofKind === "setPipelines")
            SetPipelinesCommand.internalBinaryWrite(message.command.setPipelines, writer.tag(100, WireType.LengthDelimited).fork(), options).join();
        /* protos.KeepAliveCommand keep_alive = 101; */
        if (message.command.oneofKind === "keepAlive")
            KeepAliveCommand.internalBinaryWrite(message.command.keepAlive, writer.tag(101, WireType.LengthDelimited).fork(), options).join();
        /* protos.KVCommand kv = 102; */
        if (message.command.oneofKind === "kv")
            KVCommand.internalBinaryWrite(message.command.kv, writer.tag(102, WireType.LengthDelimited).fork(), options).join();
        /* protos.TailCommand tail = 103; */
        if (message.command.oneofKind === "tail")
            TailCommand.internalBinaryWrite(message.command.tail, writer.tag(103, WireType.LengthDelimited).fork(), options).join();
        /* protos.DeleteAudienceCommand delete = 104; */
        if (message.command.oneofKind === "delete")
            DeleteAudienceCommand.internalBinaryWrite(message.command.delete, writer.tag(104, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message protos.Command
 */
export const Command = new Command$Type();
// @generated message type with reflection information, may provide speed optimized methods
class SetPipelinesCommand$Type extends MessageType<SetPipelinesCommand> {
    constructor() {
        super("protos.SetPipelinesCommand", [
            { no: 1, name: "pipelines", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Pipeline },
            { no: 2, name: "wasm_modules", kind: "map", K: 9 /*ScalarType.STRING*/, V: { kind: "message", T: () => WasmModule } }
        ]);
    }
    create(value?: PartialMessage<SetPipelinesCommand>): SetPipelinesCommand {
        const message = { pipelines: [], wasmModules: {} };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<SetPipelinesCommand>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: SetPipelinesCommand): SetPipelinesCommand {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* repeated protos.Pipeline pipelines */ 1:
                    message.pipelines.push(Pipeline.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                case /* map<string, protos.shared.WasmModule> wasm_modules */ 2:
                    this.binaryReadMap2(message.wasmModules, reader, options);
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
    private binaryReadMap2(map: SetPipelinesCommand["wasmModules"], reader: IBinaryReader, options: BinaryReadOptions): void {
        let len = reader.uint32(), end = reader.pos + len, key: keyof SetPipelinesCommand["wasmModules"] | undefined, val: SetPipelinesCommand["wasmModules"][any] | undefined;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case 1:
                    key = reader.string();
                    break;
                case 2:
                    val = WasmModule.internalBinaryRead(reader, reader.uint32(), options);
                    break;
                default: throw new globalThis.Error("unknown map entry field for field protos.SetPipelinesCommand.wasm_modules");
            }
        }
        map[key ?? ""] = val ?? WasmModule.create();
    }
    internalBinaryWrite(message: SetPipelinesCommand, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* repeated protos.Pipeline pipelines = 1; */
        for (let i = 0; i < message.pipelines.length; i++)
            Pipeline.internalBinaryWrite(message.pipelines[i], writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* map<string, protos.shared.WasmModule> wasm_modules = 2; */
        for (let k of Object.keys(message.wasmModules)) {
            writer.tag(2, WireType.LengthDelimited).fork().tag(1, WireType.LengthDelimited).string(k);
            writer.tag(2, WireType.LengthDelimited).fork();
            WasmModule.internalBinaryWrite(message.wasmModules[k], writer, options);
            writer.join().join();
        }
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message protos.SetPipelinesCommand
 */
export const SetPipelinesCommand = new SetPipelinesCommand$Type();
// @generated message type with reflection information, may provide speed optimized methods
class KeepAliveCommand$Type extends MessageType<KeepAliveCommand> {
    constructor() {
        super("protos.KeepAliveCommand", []);
    }
    create(value?: PartialMessage<KeepAliveCommand>): KeepAliveCommand {
        const message = {};
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<KeepAliveCommand>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: KeepAliveCommand): KeepAliveCommand {
        return target ?? this.create();
    }
    internalBinaryWrite(message: KeepAliveCommand, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message protos.KeepAliveCommand
 */
export const KeepAliveCommand = new KeepAliveCommand$Type();
// @generated message type with reflection information, may provide speed optimized methods
class KVCommand$Type extends MessageType<KVCommand> {
    constructor() {
        super("protos.KVCommand", [
            { no: 1, name: "instructions", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => KVInstruction },
            { no: 2, name: "overwrite", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
    create(value?: PartialMessage<KVCommand>): KVCommand {
        const message = { instructions: [], overwrite: false };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<KVCommand>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: KVCommand): KVCommand {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* repeated protos.KVInstruction instructions */ 1:
                    message.instructions.push(KVInstruction.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                case /* bool overwrite */ 2:
                    message.overwrite = reader.bool();
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
    internalBinaryWrite(message: KVCommand, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* repeated protos.KVInstruction instructions = 1; */
        for (let i = 0; i < message.instructions.length; i++)
            KVInstruction.internalBinaryWrite(message.instructions[i], writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* bool overwrite = 2; */
        if (message.overwrite !== false)
            writer.tag(2, WireType.Varint).bool(message.overwrite);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message protos.KVCommand
 */
export const KVCommand = new KVCommand$Type();
// @generated message type with reflection information, may provide speed optimized methods
class TailCommand$Type extends MessageType<TailCommand> {
    constructor() {
        super("protos.TailCommand", [
            { no: 2, name: "request", kind: "message", T: () => TailRequest }
        ]);
    }
    create(value?: PartialMessage<TailCommand>): TailCommand {
        const message = {};
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<TailCommand>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: TailCommand): TailCommand {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* protos.TailRequest request */ 2:
                    message.request = TailRequest.internalBinaryRead(reader, reader.uint32(), options, message.request);
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
    internalBinaryWrite(message: TailCommand, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* protos.TailRequest request = 2; */
        if (message.request)
            TailRequest.internalBinaryWrite(message.request, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message protos.TailCommand
 */
export const TailCommand = new TailCommand$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DeleteAudienceCommand$Type extends MessageType<DeleteAudienceCommand> {
    constructor() {
        super("protos.DeleteAudienceCommand", [
            { no: 1, name: "audience", kind: "message", T: () => Audience }
        ]);
    }
    create(value?: PartialMessage<DeleteAudienceCommand>): DeleteAudienceCommand {
        const message = {};
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<DeleteAudienceCommand>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: DeleteAudienceCommand): DeleteAudienceCommand {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* protos.Audience audience */ 1:
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
    internalBinaryWrite(message: DeleteAudienceCommand, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* protos.Audience audience = 1; */
        if (message.audience)
            Audience.internalBinaryWrite(message.audience, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message protos.DeleteAudienceCommand
 */
export const DeleteAudienceCommand = new DeleteAudienceCommand$Type();
