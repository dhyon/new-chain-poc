/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";
export const protobufPackage = "codchen.newchainpoc.dex";
const baseOrderEntry = {
    price: 0,
    quantity: 0,
    allocationCreator: "",
    allocation: 0,
};
export const OrderEntry = {
    encode(message, writer = Writer.create()) {
        if (message.price !== 0) {
            writer.uint32(8).int32(message.price);
        }
        if (message.quantity !== 0) {
            writer.uint32(16).int32(message.quantity);
        }
        for (const v of message.allocationCreator) {
            writer.uint32(26).string(v);
        }
        writer.uint32(34).fork();
        for (const v of message.allocation) {
            writer.int32(v);
        }
        writer.ldelim();
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseOrderEntry };
        message.allocationCreator = [];
        message.allocation = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.price = reader.int32();
                    break;
                case 2:
                    message.quantity = reader.int32();
                    break;
                case 3:
                    message.allocationCreator.push(reader.string());
                    break;
                case 4:
                    if ((tag & 7) === 2) {
                        const end2 = reader.uint32() + reader.pos;
                        while (reader.pos < end2) {
                            message.allocation.push(reader.int32());
                        }
                    }
                    else {
                        message.allocation.push(reader.int32());
                    }
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseOrderEntry };
        message.allocationCreator = [];
        message.allocation = [];
        if (object.price !== undefined && object.price !== null) {
            message.price = Number(object.price);
        }
        else {
            message.price = 0;
        }
        if (object.quantity !== undefined && object.quantity !== null) {
            message.quantity = Number(object.quantity);
        }
        else {
            message.quantity = 0;
        }
        if (object.allocationCreator !== undefined &&
            object.allocationCreator !== null) {
            for (const e of object.allocationCreator) {
                message.allocationCreator.push(String(e));
            }
        }
        if (object.allocation !== undefined && object.allocation !== null) {
            for (const e of object.allocation) {
                message.allocation.push(Number(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.price !== undefined && (obj.price = message.price);
        message.quantity !== undefined && (obj.quantity = message.quantity);
        if (message.allocationCreator) {
            obj.allocationCreator = message.allocationCreator.map((e) => e);
        }
        else {
            obj.allocationCreator = [];
        }
        if (message.allocation) {
            obj.allocation = message.allocation.map((e) => e);
        }
        else {
            obj.allocation = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseOrderEntry };
        message.allocationCreator = [];
        message.allocation = [];
        if (object.price !== undefined && object.price !== null) {
            message.price = object.price;
        }
        else {
            message.price = 0;
        }
        if (object.quantity !== undefined && object.quantity !== null) {
            message.quantity = object.quantity;
        }
        else {
            message.quantity = 0;
        }
        if (object.allocationCreator !== undefined &&
            object.allocationCreator !== null) {
            for (const e of object.allocationCreator) {
                message.allocationCreator.push(e);
            }
        }
        if (object.allocation !== undefined && object.allocation !== null) {
            for (const e of object.allocation) {
                message.allocation.push(e);
            }
        }
        return message;
    },
};
