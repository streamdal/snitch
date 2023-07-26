// @generated by protobuf-ts 2.9.0 with parameter optimize_code_size
// @generated from protobuf file "bus.proto" (package "protos", syntax proto3)
// tslint:disable
import { MessageType } from "@protobuf-ts/runtime";
import { HeartbeatRequest } from "./internal.js";
import { DeregisterRequest } from "./internal.js";
import { RegisterRequest } from "./internal.js";
import { Command } from "./command.js";
/**
 * Type used by `snitch-server` for sending messages on its local bus.
 *
 * @generated from protobuf message protos.BusEvent
 */
export interface BusEvent {
    /**
     * @generated from protobuf field: string source = 1;
     */
    source: string;
    /**
     * @generated from protobuf oneof: event
     */
    event: {
        oneofKind: "commandResponse";
        /**
         * @generated from protobuf field: protos.Command command_response = 100;
         */
        commandResponse: Command;
    } | {
        oneofKind: "registerRequest";
        /**
         * @generated from protobuf field: protos.RegisterRequest register_request = 101;
         */
        registerRequest: RegisterRequest;
    } | {
        oneofKind: "deregisterRequest";
        /**
         * @generated from protobuf field: protos.DeregisterRequest deregister_request = 102;
         */
        deregisterRequest: DeregisterRequest;
    } | {
        oneofKind: "heartbeatRequest";
        /**
         * @generated from protobuf field: protos.HeartbeatRequest heartbeat_request = 103;
         */
        heartbeatRequest: HeartbeatRequest;
    } | {
        oneofKind: undefined;
    };
    /**
     * All gRPC metadata is stored in ctx; when request goes outside of gRPC
     * bounds, we will translate ctx metadata into this field.
     *
     * Example:
     * 1. Request comes into snitch-server via external gRPC to set new pipeline
     * 2. snitch-server has to send SetPipeline cmd to SDK via gRPC - it passes
     *    on original metadata in request.
     * 3. snitch-server has to broadcast SetPipeline cmd to other services via bus
     * 4. Since this is not a gRPC call, snitch-server translates ctx metadata to
     *    this field and includes it in the bus event.
     *
     * @generated from protobuf field: map<string, string> _metadata = 1000;
     */
    Metadata: {
        [key: string]: string;
    }; // protolint:disable:this FIELD_NAMES_LOWER_SNAKE_CASE
}
// @generated message type with reflection information, may provide speed optimized methods
class BusEvent$Type extends MessageType<BusEvent> {
    constructor() {
        super("protos.BusEvent", [
            { no: 1, name: "source", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 100, name: "command_response", kind: "message", oneof: "event", T: () => Command },
            { no: 101, name: "register_request", kind: "message", oneof: "event", T: () => RegisterRequest },
            { no: 102, name: "deregister_request", kind: "message", oneof: "event", T: () => DeregisterRequest },
            { no: 103, name: "heartbeat_request", kind: "message", oneof: "event", T: () => HeartbeatRequest },
            { no: 1000, name: "_metadata", kind: "map", K: 9 /*ScalarType.STRING*/, V: { kind: "scalar", T: 9 /*ScalarType.STRING*/ } }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.BusEvent
 */
export const BusEvent = new BusEvent$Type();
