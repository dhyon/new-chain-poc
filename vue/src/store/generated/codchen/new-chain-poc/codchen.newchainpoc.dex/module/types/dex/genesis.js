/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Params } from "../dex/params";
import { LongBook } from "../dex/long_book";
import { ShortBook } from "../dex/short_book";
export const protobufPackage = "codchen.newchainpoc.dex";
const baseGenesisState = { longBookCount: 0, shortBookCount: 0 };
export const GenesisState = {
    encode(message, writer = Writer.create()) {
        if (message.params !== undefined) {
            Params.encode(message.params, writer.uint32(10).fork()).ldelim();
        }
        for (const v of message.longBookList) {
            LongBook.encode(v, writer.uint32(18).fork()).ldelim();
        }
        if (message.longBookCount !== 0) {
            writer.uint32(24).uint64(message.longBookCount);
        }
        for (const v of message.shortBookList) {
            ShortBook.encode(v, writer.uint32(34).fork()).ldelim();
        }
        if (message.shortBookCount !== 0) {
            writer.uint32(40).uint64(message.shortBookCount);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGenesisState };
        message.longBookList = [];
        message.shortBookList = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.params = Params.decode(reader, reader.uint32());
                    break;
                case 2:
                    message.longBookList.push(LongBook.decode(reader, reader.uint32()));
                    break;
                case 3:
                    message.longBookCount = longToNumber(reader.uint64());
                    break;
                case 4:
                    message.shortBookList.push(ShortBook.decode(reader, reader.uint32()));
                    break;
                case 5:
                    message.shortBookCount = longToNumber(reader.uint64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseGenesisState };
        message.longBookList = [];
        message.shortBookList = [];
        if (object.params !== undefined && object.params !== null) {
            message.params = Params.fromJSON(object.params);
        }
        else {
            message.params = undefined;
        }
        if (object.longBookList !== undefined && object.longBookList !== null) {
            for (const e of object.longBookList) {
                message.longBookList.push(LongBook.fromJSON(e));
            }
        }
        if (object.longBookCount !== undefined && object.longBookCount !== null) {
            message.longBookCount = Number(object.longBookCount);
        }
        else {
            message.longBookCount = 0;
        }
        if (object.shortBookList !== undefined && object.shortBookList !== null) {
            for (const e of object.shortBookList) {
                message.shortBookList.push(ShortBook.fromJSON(e));
            }
        }
        if (object.shortBookCount !== undefined && object.shortBookCount !== null) {
            message.shortBookCount = Number(object.shortBookCount);
        }
        else {
            message.shortBookCount = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.params !== undefined &&
            (obj.params = message.params ? Params.toJSON(message.params) : undefined);
        if (message.longBookList) {
            obj.longBookList = message.longBookList.map((e) => e ? LongBook.toJSON(e) : undefined);
        }
        else {
            obj.longBookList = [];
        }
        message.longBookCount !== undefined &&
            (obj.longBookCount = message.longBookCount);
        if (message.shortBookList) {
            obj.shortBookList = message.shortBookList.map((e) => e ? ShortBook.toJSON(e) : undefined);
        }
        else {
            obj.shortBookList = [];
        }
        message.shortBookCount !== undefined &&
            (obj.shortBookCount = message.shortBookCount);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGenesisState };
        message.longBookList = [];
        message.shortBookList = [];
        if (object.params !== undefined && object.params !== null) {
            message.params = Params.fromPartial(object.params);
        }
        else {
            message.params = undefined;
        }
        if (object.longBookList !== undefined && object.longBookList !== null) {
            for (const e of object.longBookList) {
                message.longBookList.push(LongBook.fromPartial(e));
            }
        }
        if (object.longBookCount !== undefined && object.longBookCount !== null) {
            message.longBookCount = object.longBookCount;
        }
        else {
            message.longBookCount = 0;
        }
        if (object.shortBookList !== undefined && object.shortBookList !== null) {
            for (const e of object.shortBookList) {
                message.shortBookList.push(ShortBook.fromPartial(e));
            }
        }
        if (object.shortBookCount !== undefined && object.shortBookCount !== null) {
            message.shortBookCount = object.shortBookCount;
        }
        else {
            message.shortBookCount = 0;
        }
        return message;
    },
};
var globalThis = (() => {
    if (typeof globalThis !== "undefined")
        return globalThis;
    if (typeof self !== "undefined")
        return self;
    if (typeof window !== "undefined")
        return window;
    if (typeof global !== "undefined")
        return global;
    throw "Unable to locate global object";
})();
function longToNumber(long) {
    if (long.gt(Number.MAX_SAFE_INTEGER)) {
        throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
    }
    return long.toNumber();
}
if (util.Long !== Long) {
    util.Long = Long;
    configure();
}
