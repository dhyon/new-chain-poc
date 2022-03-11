import { Reader, Writer } from "protobufjs/minimal";
export declare const protobufPackage = "codchen.newchainpoc.dex";
export interface MsgLimitBuy {
    creator: string;
    price: number;
    quantity: number;
}
export interface MsgLimitBuyResponse {
}
export interface MsgLimitSell {
    creator: string;
    price: number;
    quantity: number;
}
export interface MsgLimitSellResponse {
}
export interface MsgMarketBuy {
    creator: string;
    quantity: number;
}
export interface MsgMarketBuyResponse {
}
export interface MsgMarketSell {
    creator: string;
    quantity: number;
}
export interface MsgMarketSellResponse {
}
export declare const MsgLimitBuy: {
    encode(message: MsgLimitBuy, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgLimitBuy;
    fromJSON(object: any): MsgLimitBuy;
    toJSON(message: MsgLimitBuy): unknown;
    fromPartial(object: DeepPartial<MsgLimitBuy>): MsgLimitBuy;
};
export declare const MsgLimitBuyResponse: {
    encode(_: MsgLimitBuyResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgLimitBuyResponse;
    fromJSON(_: any): MsgLimitBuyResponse;
    toJSON(_: MsgLimitBuyResponse): unknown;
    fromPartial(_: DeepPartial<MsgLimitBuyResponse>): MsgLimitBuyResponse;
};
export declare const MsgLimitSell: {
    encode(message: MsgLimitSell, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgLimitSell;
    fromJSON(object: any): MsgLimitSell;
    toJSON(message: MsgLimitSell): unknown;
    fromPartial(object: DeepPartial<MsgLimitSell>): MsgLimitSell;
};
export declare const MsgLimitSellResponse: {
    encode(_: MsgLimitSellResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgLimitSellResponse;
    fromJSON(_: any): MsgLimitSellResponse;
    toJSON(_: MsgLimitSellResponse): unknown;
    fromPartial(_: DeepPartial<MsgLimitSellResponse>): MsgLimitSellResponse;
};
export declare const MsgMarketBuy: {
    encode(message: MsgMarketBuy, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgMarketBuy;
    fromJSON(object: any): MsgMarketBuy;
    toJSON(message: MsgMarketBuy): unknown;
    fromPartial(object: DeepPartial<MsgMarketBuy>): MsgMarketBuy;
};
export declare const MsgMarketBuyResponse: {
    encode(_: MsgMarketBuyResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgMarketBuyResponse;
    fromJSON(_: any): MsgMarketBuyResponse;
    toJSON(_: MsgMarketBuyResponse): unknown;
    fromPartial(_: DeepPartial<MsgMarketBuyResponse>): MsgMarketBuyResponse;
};
export declare const MsgMarketSell: {
    encode(message: MsgMarketSell, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgMarketSell;
    fromJSON(object: any): MsgMarketSell;
    toJSON(message: MsgMarketSell): unknown;
    fromPartial(object: DeepPartial<MsgMarketSell>): MsgMarketSell;
};
export declare const MsgMarketSellResponse: {
    encode(_: MsgMarketSellResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgMarketSellResponse;
    fromJSON(_: any): MsgMarketSellResponse;
    toJSON(_: MsgMarketSellResponse): unknown;
    fromPartial(_: DeepPartial<MsgMarketSellResponse>): MsgMarketSellResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    LimitBuy(request: MsgLimitBuy): Promise<MsgLimitBuyResponse>;
    LimitSell(request: MsgLimitSell): Promise<MsgLimitSellResponse>;
    MarketBuy(request: MsgMarketBuy): Promise<MsgMarketBuyResponse>;
    /** this line is used by starport scaffolding # proto/tx/rpc */
    MarketSell(request: MsgMarketSell): Promise<MsgMarketSellResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    LimitBuy(request: MsgLimitBuy): Promise<MsgLimitBuyResponse>;
    LimitSell(request: MsgLimitSell): Promise<MsgLimitSellResponse>;
    MarketBuy(request: MsgMarketBuy): Promise<MsgMarketBuyResponse>;
    MarketSell(request: MsgMarketSell): Promise<MsgMarketSellResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
