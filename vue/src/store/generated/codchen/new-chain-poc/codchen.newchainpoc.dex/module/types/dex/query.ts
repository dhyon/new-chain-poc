/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Params } from "../dex/params";
import { LongBook } from "../dex/long_book";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";
import { ShortBook } from "../dex/short_book";

export const protobufPackage = "codchen.newchainpoc.dex";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

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

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryGetLongBookRequest: object = { id: 0 };

export const QueryGetLongBookRequest = {
  encode(
    message: QueryGetLongBookRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetLongBookRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetLongBookRequest,
    } as QueryGetLongBookRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetLongBookRequest {
    const message = {
      ...baseQueryGetLongBookRequest,
    } as QueryGetLongBookRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: QueryGetLongBookRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetLongBookRequest>
  ): QueryGetLongBookRequest {
    const message = {
      ...baseQueryGetLongBookRequest,
    } as QueryGetLongBookRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseQueryGetLongBookResponse: object = {};

export const QueryGetLongBookResponse = {
  encode(
    message: QueryGetLongBookResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.LongBook !== undefined) {
      LongBook.encode(message.LongBook, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetLongBookResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetLongBookResponse,
    } as QueryGetLongBookResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.LongBook = LongBook.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetLongBookResponse {
    const message = {
      ...baseQueryGetLongBookResponse,
    } as QueryGetLongBookResponse;
    if (object.LongBook !== undefined && object.LongBook !== null) {
      message.LongBook = LongBook.fromJSON(object.LongBook);
    } else {
      message.LongBook = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetLongBookResponse): unknown {
    const obj: any = {};
    message.LongBook !== undefined &&
      (obj.LongBook = message.LongBook
        ? LongBook.toJSON(message.LongBook)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetLongBookResponse>
  ): QueryGetLongBookResponse {
    const message = {
      ...baseQueryGetLongBookResponse,
    } as QueryGetLongBookResponse;
    if (object.LongBook !== undefined && object.LongBook !== null) {
      message.LongBook = LongBook.fromPartial(object.LongBook);
    } else {
      message.LongBook = undefined;
    }
    return message;
  },
};

const baseQueryAllLongBookRequest: object = {};

export const QueryAllLongBookRequest = {
  encode(
    message: QueryAllLongBookRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllLongBookRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllLongBookRequest,
    } as QueryAllLongBookRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllLongBookRequest {
    const message = {
      ...baseQueryAllLongBookRequest,
    } as QueryAllLongBookRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllLongBookRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllLongBookRequest>
  ): QueryAllLongBookRequest {
    const message = {
      ...baseQueryAllLongBookRequest,
    } as QueryAllLongBookRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllLongBookResponse: object = {};

export const QueryAllLongBookResponse = {
  encode(
    message: QueryAllLongBookResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.LongBook) {
      LongBook.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllLongBookResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllLongBookResponse,
    } as QueryAllLongBookResponse;
    message.LongBook = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.LongBook.push(LongBook.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllLongBookResponse {
    const message = {
      ...baseQueryAllLongBookResponse,
    } as QueryAllLongBookResponse;
    message.LongBook = [];
    if (object.LongBook !== undefined && object.LongBook !== null) {
      for (const e of object.LongBook) {
        message.LongBook.push(LongBook.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllLongBookResponse): unknown {
    const obj: any = {};
    if (message.LongBook) {
      obj.LongBook = message.LongBook.map((e) =>
        e ? LongBook.toJSON(e) : undefined
      );
    } else {
      obj.LongBook = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllLongBookResponse>
  ): QueryAllLongBookResponse {
    const message = {
      ...baseQueryAllLongBookResponse,
    } as QueryAllLongBookResponse;
    message.LongBook = [];
    if (object.LongBook !== undefined && object.LongBook !== null) {
      for (const e of object.LongBook) {
        message.LongBook.push(LongBook.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetShortBookRequest: object = { id: 0 };

export const QueryGetShortBookRequest = {
  encode(
    message: QueryGetShortBookRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetShortBookRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetShortBookRequest,
    } as QueryGetShortBookRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetShortBookRequest {
    const message = {
      ...baseQueryGetShortBookRequest,
    } as QueryGetShortBookRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: QueryGetShortBookRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetShortBookRequest>
  ): QueryGetShortBookRequest {
    const message = {
      ...baseQueryGetShortBookRequest,
    } as QueryGetShortBookRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseQueryGetShortBookResponse: object = {};

export const QueryGetShortBookResponse = {
  encode(
    message: QueryGetShortBookResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.ShortBook !== undefined) {
      ShortBook.encode(message.ShortBook, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetShortBookResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetShortBookResponse,
    } as QueryGetShortBookResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.ShortBook = ShortBook.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetShortBookResponse {
    const message = {
      ...baseQueryGetShortBookResponse,
    } as QueryGetShortBookResponse;
    if (object.ShortBook !== undefined && object.ShortBook !== null) {
      message.ShortBook = ShortBook.fromJSON(object.ShortBook);
    } else {
      message.ShortBook = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetShortBookResponse): unknown {
    const obj: any = {};
    message.ShortBook !== undefined &&
      (obj.ShortBook = message.ShortBook
        ? ShortBook.toJSON(message.ShortBook)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetShortBookResponse>
  ): QueryGetShortBookResponse {
    const message = {
      ...baseQueryGetShortBookResponse,
    } as QueryGetShortBookResponse;
    if (object.ShortBook !== undefined && object.ShortBook !== null) {
      message.ShortBook = ShortBook.fromPartial(object.ShortBook);
    } else {
      message.ShortBook = undefined;
    }
    return message;
  },
};

const baseQueryAllShortBookRequest: object = {};

export const QueryAllShortBookRequest = {
  encode(
    message: QueryAllShortBookRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllShortBookRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllShortBookRequest,
    } as QueryAllShortBookRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllShortBookRequest {
    const message = {
      ...baseQueryAllShortBookRequest,
    } as QueryAllShortBookRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllShortBookRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllShortBookRequest>
  ): QueryAllShortBookRequest {
    const message = {
      ...baseQueryAllShortBookRequest,
    } as QueryAllShortBookRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllShortBookResponse: object = {};

export const QueryAllShortBookResponse = {
  encode(
    message: QueryAllShortBookResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.ShortBook) {
      ShortBook.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllShortBookResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllShortBookResponse,
    } as QueryAllShortBookResponse;
    message.ShortBook = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.ShortBook.push(ShortBook.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllShortBookResponse {
    const message = {
      ...baseQueryAllShortBookResponse,
    } as QueryAllShortBookResponse;
    message.ShortBook = [];
    if (object.ShortBook !== undefined && object.ShortBook !== null) {
      for (const e of object.ShortBook) {
        message.ShortBook.push(ShortBook.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllShortBookResponse): unknown {
    const obj: any = {};
    if (message.ShortBook) {
      obj.ShortBook = message.ShortBook.map((e) =>
        e ? ShortBook.toJSON(e) : undefined
      );
    } else {
      obj.ShortBook = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllShortBookResponse>
  ): QueryAllShortBookResponse {
    const message = {
      ...baseQueryAllShortBookResponse,
    } as QueryAllShortBookResponse;
    message.ShortBook = [];
    if (object.ShortBook !== undefined && object.ShortBook !== null) {
      for (const e of object.ShortBook) {
        message.ShortBook.push(ShortBook.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a LongBook by id. */
  LongBook(request: QueryGetLongBookRequest): Promise<QueryGetLongBookResponse>;
  /** Queries a list of LongBook items. */
  LongBookAll(
    request: QueryAllLongBookRequest
  ): Promise<QueryAllLongBookResponse>;
  /** Queries a ShortBook by id. */
  ShortBook(
    request: QueryGetShortBookRequest
  ): Promise<QueryGetShortBookResponse>;
  /** Queries a list of ShortBook items. */
  ShortBookAll(
    request: QueryAllShortBookRequest
  ): Promise<QueryAllShortBookResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "codchen.newchainpoc.dex.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  LongBook(
    request: QueryGetLongBookRequest
  ): Promise<QueryGetLongBookResponse> {
    const data = QueryGetLongBookRequest.encode(request).finish();
    const promise = this.rpc.request(
      "codchen.newchainpoc.dex.Query",
      "LongBook",
      data
    );
    return promise.then((data) =>
      QueryGetLongBookResponse.decode(new Reader(data))
    );
  }

  LongBookAll(
    request: QueryAllLongBookRequest
  ): Promise<QueryAllLongBookResponse> {
    const data = QueryAllLongBookRequest.encode(request).finish();
    const promise = this.rpc.request(
      "codchen.newchainpoc.dex.Query",
      "LongBookAll",
      data
    );
    return promise.then((data) =>
      QueryAllLongBookResponse.decode(new Reader(data))
    );
  }

  ShortBook(
    request: QueryGetShortBookRequest
  ): Promise<QueryGetShortBookResponse> {
    const data = QueryGetShortBookRequest.encode(request).finish();
    const promise = this.rpc.request(
      "codchen.newchainpoc.dex.Query",
      "ShortBook",
      data
    );
    return promise.then((data) =>
      QueryGetShortBookResponse.decode(new Reader(data))
    );
  }

  ShortBookAll(
    request: QueryAllShortBookRequest
  ): Promise<QueryAllShortBookResponse> {
    const data = QueryAllShortBookRequest.encode(request).finish();
    const promise = this.rpc.request(
      "codchen.newchainpoc.dex.Query",
      "ShortBookAll",
      data
    );
    return promise.then((data) =>
      QueryAllShortBookResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
