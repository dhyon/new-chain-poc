/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "codchen.newchainpoc.dex";

export interface MsgLimitBuy {
  creator: string;
  price: number;
  quantity: number;
}

export interface MsgLimitBuyResponse {}

export interface MsgLimitSell {
  creator: string;
  price: number;
  quantity: number;
}

export interface MsgLimitSellResponse {}

export interface MsgMarketBuy {
  creator: string;
  quantity: number;
}

export interface MsgMarketBuyResponse {}

export interface MsgMarketSell {
  creator: string;
  quantity: number;
}

export interface MsgMarketSellResponse {}

const baseMsgLimitBuy: object = { creator: "", price: 0, quantity: 0 };

export const MsgLimitBuy = {
  encode(message: MsgLimitBuy, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.price !== 0) {
      writer.uint32(16).int32(message.price);
    }
    if (message.quantity !== 0) {
      writer.uint32(24).int32(message.quantity);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgLimitBuy {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgLimitBuy } as MsgLimitBuy;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.price = reader.int32();
          break;
        case 3:
          message.quantity = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgLimitBuy {
    const message = { ...baseMsgLimitBuy } as MsgLimitBuy;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.price !== undefined && object.price !== null) {
      message.price = Number(object.price);
    } else {
      message.price = 0;
    }
    if (object.quantity !== undefined && object.quantity !== null) {
      message.quantity = Number(object.quantity);
    } else {
      message.quantity = 0;
    }
    return message;
  },

  toJSON(message: MsgLimitBuy): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.price !== undefined && (obj.price = message.price);
    message.quantity !== undefined && (obj.quantity = message.quantity);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgLimitBuy>): MsgLimitBuy {
    const message = { ...baseMsgLimitBuy } as MsgLimitBuy;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.price !== undefined && object.price !== null) {
      message.price = object.price;
    } else {
      message.price = 0;
    }
    if (object.quantity !== undefined && object.quantity !== null) {
      message.quantity = object.quantity;
    } else {
      message.quantity = 0;
    }
    return message;
  },
};

const baseMsgLimitBuyResponse: object = {};

export const MsgLimitBuyResponse = {
  encode(_: MsgLimitBuyResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgLimitBuyResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgLimitBuyResponse } as MsgLimitBuyResponse;
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

  fromJSON(_: any): MsgLimitBuyResponse {
    const message = { ...baseMsgLimitBuyResponse } as MsgLimitBuyResponse;
    return message;
  },

  toJSON(_: MsgLimitBuyResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgLimitBuyResponse>): MsgLimitBuyResponse {
    const message = { ...baseMsgLimitBuyResponse } as MsgLimitBuyResponse;
    return message;
  },
};

const baseMsgLimitSell: object = { creator: "", price: 0, quantity: 0 };

export const MsgLimitSell = {
  encode(message: MsgLimitSell, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.price !== 0) {
      writer.uint32(16).int32(message.price);
    }
    if (message.quantity !== 0) {
      writer.uint32(24).int32(message.quantity);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgLimitSell {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgLimitSell } as MsgLimitSell;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.price = reader.int32();
          break;
        case 3:
          message.quantity = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgLimitSell {
    const message = { ...baseMsgLimitSell } as MsgLimitSell;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.price !== undefined && object.price !== null) {
      message.price = Number(object.price);
    } else {
      message.price = 0;
    }
    if (object.quantity !== undefined && object.quantity !== null) {
      message.quantity = Number(object.quantity);
    } else {
      message.quantity = 0;
    }
    return message;
  },

  toJSON(message: MsgLimitSell): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.price !== undefined && (obj.price = message.price);
    message.quantity !== undefined && (obj.quantity = message.quantity);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgLimitSell>): MsgLimitSell {
    const message = { ...baseMsgLimitSell } as MsgLimitSell;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.price !== undefined && object.price !== null) {
      message.price = object.price;
    } else {
      message.price = 0;
    }
    if (object.quantity !== undefined && object.quantity !== null) {
      message.quantity = object.quantity;
    } else {
      message.quantity = 0;
    }
    return message;
  },
};

const baseMsgLimitSellResponse: object = {};

export const MsgLimitSellResponse = {
  encode(_: MsgLimitSellResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgLimitSellResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgLimitSellResponse } as MsgLimitSellResponse;
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

  fromJSON(_: any): MsgLimitSellResponse {
    const message = { ...baseMsgLimitSellResponse } as MsgLimitSellResponse;
    return message;
  },

  toJSON(_: MsgLimitSellResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgLimitSellResponse>): MsgLimitSellResponse {
    const message = { ...baseMsgLimitSellResponse } as MsgLimitSellResponse;
    return message;
  },
};

const baseMsgMarketBuy: object = { creator: "", quantity: 0 };

export const MsgMarketBuy = {
  encode(message: MsgMarketBuy, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.quantity !== 0) {
      writer.uint32(16).int32(message.quantity);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgMarketBuy {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgMarketBuy } as MsgMarketBuy;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.quantity = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgMarketBuy {
    const message = { ...baseMsgMarketBuy } as MsgMarketBuy;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.quantity !== undefined && object.quantity !== null) {
      message.quantity = Number(object.quantity);
    } else {
      message.quantity = 0;
    }
    return message;
  },

  toJSON(message: MsgMarketBuy): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.quantity !== undefined && (obj.quantity = message.quantity);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgMarketBuy>): MsgMarketBuy {
    const message = { ...baseMsgMarketBuy } as MsgMarketBuy;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.quantity !== undefined && object.quantity !== null) {
      message.quantity = object.quantity;
    } else {
      message.quantity = 0;
    }
    return message;
  },
};

const baseMsgMarketBuyResponse: object = {};

export const MsgMarketBuyResponse = {
  encode(_: MsgMarketBuyResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgMarketBuyResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgMarketBuyResponse } as MsgMarketBuyResponse;
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

  fromJSON(_: any): MsgMarketBuyResponse {
    const message = { ...baseMsgMarketBuyResponse } as MsgMarketBuyResponse;
    return message;
  },

  toJSON(_: MsgMarketBuyResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgMarketBuyResponse>): MsgMarketBuyResponse {
    const message = { ...baseMsgMarketBuyResponse } as MsgMarketBuyResponse;
    return message;
  },
};

const baseMsgMarketSell: object = { creator: "", quantity: 0 };

export const MsgMarketSell = {
  encode(message: MsgMarketSell, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.quantity !== 0) {
      writer.uint32(16).int32(message.quantity);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgMarketSell {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgMarketSell } as MsgMarketSell;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.quantity = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgMarketSell {
    const message = { ...baseMsgMarketSell } as MsgMarketSell;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.quantity !== undefined && object.quantity !== null) {
      message.quantity = Number(object.quantity);
    } else {
      message.quantity = 0;
    }
    return message;
  },

  toJSON(message: MsgMarketSell): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.quantity !== undefined && (obj.quantity = message.quantity);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgMarketSell>): MsgMarketSell {
    const message = { ...baseMsgMarketSell } as MsgMarketSell;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.quantity !== undefined && object.quantity !== null) {
      message.quantity = object.quantity;
    } else {
      message.quantity = 0;
    }
    return message;
  },
};

const baseMsgMarketSellResponse: object = {};

export const MsgMarketSellResponse = {
  encode(_: MsgMarketSellResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgMarketSellResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgMarketSellResponse } as MsgMarketSellResponse;
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

  fromJSON(_: any): MsgMarketSellResponse {
    const message = { ...baseMsgMarketSellResponse } as MsgMarketSellResponse;
    return message;
  },

  toJSON(_: MsgMarketSellResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgMarketSellResponse>): MsgMarketSellResponse {
    const message = { ...baseMsgMarketSellResponse } as MsgMarketSellResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  LimitBuy(request: MsgLimitBuy): Promise<MsgLimitBuyResponse>;
  LimitSell(request: MsgLimitSell): Promise<MsgLimitSellResponse>;
  MarketBuy(request: MsgMarketBuy): Promise<MsgMarketBuyResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  MarketSell(request: MsgMarketSell): Promise<MsgMarketSellResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  LimitBuy(request: MsgLimitBuy): Promise<MsgLimitBuyResponse> {
    const data = MsgLimitBuy.encode(request).finish();
    const promise = this.rpc.request(
      "codchen.newchainpoc.dex.Msg",
      "LimitBuy",
      data
    );
    return promise.then((data) => MsgLimitBuyResponse.decode(new Reader(data)));
  }

  LimitSell(request: MsgLimitSell): Promise<MsgLimitSellResponse> {
    const data = MsgLimitSell.encode(request).finish();
    const promise = this.rpc.request(
      "codchen.newchainpoc.dex.Msg",
      "LimitSell",
      data
    );
    return promise.then((data) =>
      MsgLimitSellResponse.decode(new Reader(data))
    );
  }

  MarketBuy(request: MsgMarketBuy): Promise<MsgMarketBuyResponse> {
    const data = MsgMarketBuy.encode(request).finish();
    const promise = this.rpc.request(
      "codchen.newchainpoc.dex.Msg",
      "MarketBuy",
      data
    );
    return promise.then((data) =>
      MsgMarketBuyResponse.decode(new Reader(data))
    );
  }

  MarketSell(request: MsgMarketSell): Promise<MsgMarketSellResponse> {
    const data = MsgMarketSell.encode(request).finish();
    const promise = this.rpc.request(
      "codchen.newchainpoc.dex.Msg",
      "MarketSell",
      data
    );
    return promise.then((data) =>
      MsgMarketSellResponse.decode(new Reader(data))
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
