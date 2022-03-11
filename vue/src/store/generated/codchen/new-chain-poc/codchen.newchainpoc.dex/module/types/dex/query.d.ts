import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../dex/params";
import { LongBook } from "../dex/long_book";
import { PageRequest, PageResponse } from "../cosmos/base/query/v1beta1/pagination";
import { ShortBook } from "../dex/short_book";
export declare const protobufPackage = "codchen.newchainpoc.dex";
/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}
/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
    /** params holds all the parameters of this module. */
    params: Params | undefined;
}
export interface QueryGetLongBookRequest {
    id: number;
}
export interface QueryGetLongBookResponse {
    LongBook: LongBook | undefined;
}
export interface QueryAllLongBookRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllLongBookResponse {
    LongBook: LongBook[];
    pagination: PageResponse | undefined;
}
export interface QueryGetShortBookRequest {
    id: number;
}
export interface QueryGetShortBookResponse {
    ShortBook: ShortBook | undefined;
}
export interface QueryAllShortBookRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllShortBookResponse {
    ShortBook: ShortBook[];
    pagination: PageResponse | undefined;
}
export declare const QueryParamsRequest: {
    encode(_: QueryParamsRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest;
    fromJSON(_: any): QueryParamsRequest;
    toJSON(_: QueryParamsRequest): unknown;
    fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest;
};
export declare const QueryParamsResponse: {
    encode(message: QueryParamsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse;
    fromJSON(object: any): QueryParamsResponse;
    toJSON(message: QueryParamsResponse): unknown;
    fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse;
};
export declare const QueryGetLongBookRequest: {
    encode(message: QueryGetLongBookRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetLongBookRequest;
    fromJSON(object: any): QueryGetLongBookRequest;
    toJSON(message: QueryGetLongBookRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetLongBookRequest>): QueryGetLongBookRequest;
};
export declare const QueryGetLongBookResponse: {
    encode(message: QueryGetLongBookResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetLongBookResponse;
    fromJSON(object: any): QueryGetLongBookResponse;
    toJSON(message: QueryGetLongBookResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetLongBookResponse>): QueryGetLongBookResponse;
};
export declare const QueryAllLongBookRequest: {
    encode(message: QueryAllLongBookRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllLongBookRequest;
    fromJSON(object: any): QueryAllLongBookRequest;
    toJSON(message: QueryAllLongBookRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllLongBookRequest>): QueryAllLongBookRequest;
};
export declare const QueryAllLongBookResponse: {
    encode(message: QueryAllLongBookResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllLongBookResponse;
    fromJSON(object: any): QueryAllLongBookResponse;
    toJSON(message: QueryAllLongBookResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllLongBookResponse>): QueryAllLongBookResponse;
};
export declare const QueryGetShortBookRequest: {
    encode(message: QueryGetShortBookRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetShortBookRequest;
    fromJSON(object: any): QueryGetShortBookRequest;
    toJSON(message: QueryGetShortBookRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetShortBookRequest>): QueryGetShortBookRequest;
};
export declare const QueryGetShortBookResponse: {
    encode(message: QueryGetShortBookResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetShortBookResponse;
    fromJSON(object: any): QueryGetShortBookResponse;
    toJSON(message: QueryGetShortBookResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetShortBookResponse>): QueryGetShortBookResponse;
};
export declare const QueryAllShortBookRequest: {
    encode(message: QueryAllShortBookRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllShortBookRequest;
    fromJSON(object: any): QueryAllShortBookRequest;
    toJSON(message: QueryAllShortBookRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllShortBookRequest>): QueryAllShortBookRequest;
};
export declare const QueryAllShortBookResponse: {
    encode(message: QueryAllShortBookResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllShortBookResponse;
    fromJSON(object: any): QueryAllShortBookResponse;
    toJSON(message: QueryAllShortBookResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllShortBookResponse>): QueryAllShortBookResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Parameters queries the parameters of the module. */
    Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
    /** Queries a LongBook by id. */
    LongBook(request: QueryGetLongBookRequest): Promise<QueryGetLongBookResponse>;
    /** Queries a list of LongBook items. */
    LongBookAll(request: QueryAllLongBookRequest): Promise<QueryAllLongBookResponse>;
    /** Queries a ShortBook by id. */
    ShortBook(request: QueryGetShortBookRequest): Promise<QueryGetShortBookResponse>;
    /** Queries a list of ShortBook items. */
    ShortBookAll(request: QueryAllShortBookRequest): Promise<QueryAllShortBookResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
    LongBook(request: QueryGetLongBookRequest): Promise<QueryGetLongBookResponse>;
    LongBookAll(request: QueryAllLongBookRequest): Promise<QueryAllLongBookResponse>;
    ShortBook(request: QueryGetShortBookRequest): Promise<QueryGetShortBookResponse>;
    ShortBookAll(request: QueryAllShortBookRequest): Promise<QueryAllShortBookResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
