// @generated by protobuf-ts 2.9.0 with parameter optimize_code_size,client_grpc1
// @generated from protobuf file "steps/transform.proto" (package "protos.steps", syntax proto3)
// tslint:disable
import { MessageType } from "@protobuf-ts/runtime";
/**
 * @generated from protobuf message protos.steps.TransformStep
 */
export interface TransformStep {
    /**
     * @generated from protobuf field: string path = 1;
     */
    path: string;
    /**
     * @generated from protobuf field: string value = 2;
     */
    value: string; // Should this be bytes? ~DS
    /**
     * @generated from protobuf field: protos.steps.TransformType type = 3;
     */
    type: TransformType;
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
    MASK_VALUE = 4
}
// @generated message type with reflection information, may provide speed optimized methods
class TransformStep$Type extends MessageType<TransformStep> {
    constructor() {
        super("protos.steps.TransformStep", [
            { no: 1, name: "path", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "value", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "type", kind: "enum", T: () => ["protos.steps.TransformType", TransformType, "TRANSFORM_TYPE_"] }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.steps.TransformStep
 */
export const TransformStep = new TransformStep$Type();
