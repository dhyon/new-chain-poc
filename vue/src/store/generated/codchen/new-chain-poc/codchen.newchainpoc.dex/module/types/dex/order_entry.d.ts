import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "codchen.newchainpoc.dex";
export interface OrderEntry {
    price: number;
    quantity: number;
    allocationCreator: string[];
    allocation: number[];
}
export declare const OrderEntry: {
    encode(message: OrderEntry, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): OrderEntry;
    fromJSON(object: any): OrderEntry;
    toJSON(message: OrderEntry): unknown;
    fromPartial(object: DeepPartial<OrderEntry>): OrderEntry;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
