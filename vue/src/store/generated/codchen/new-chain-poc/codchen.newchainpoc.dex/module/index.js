// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgLimitBuy } from "./types/dex/tx";
import { MsgLimitSell } from "./types/dex/tx";
import { MsgMarketBuy } from "./types/dex/tx";
import { MsgMarketSell } from "./types/dex/tx";
const types = [
    ["/codchen.newchainpoc.dex.MsgLimitBuy", MsgLimitBuy],
    ["/codchen.newchainpoc.dex.MsgLimitSell", MsgLimitSell],
    ["/codchen.newchainpoc.dex.MsgMarketBuy", MsgMarketBuy],
    ["/codchen.newchainpoc.dex.MsgMarketSell", MsgMarketSell],
];
export const MissingWalletError = new Error("wallet is required");
export const registry = new Registry(types);
const defaultFee = {
    amount: [],
    gas: "200000",
};
const txClient = async (wallet, { addr: addr } = { addr: "http://localhost:26657" }) => {
    if (!wallet)
        throw MissingWalletError;
    let client;
    if (addr) {
        client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
    }
    else {
        client = await SigningStargateClient.offline(wallet, { registry });
    }
    const { address } = (await wallet.getAccounts())[0];
    return {
        signAndBroadcast: (msgs, { fee, memo } = { fee: defaultFee, memo: "" }) => client.signAndBroadcast(address, msgs, fee, memo),
        msgLimitBuy: (data) => ({ typeUrl: "/codchen.newchainpoc.dex.MsgLimitBuy", value: MsgLimitBuy.fromPartial(data) }),
        msgLimitSell: (data) => ({ typeUrl: "/codchen.newchainpoc.dex.MsgLimitSell", value: MsgLimitSell.fromPartial(data) }),
        msgMarketBuy: (data) => ({ typeUrl: "/codchen.newchainpoc.dex.MsgMarketBuy", value: MsgMarketBuy.fromPartial(data) }),
        msgMarketSell: (data) => ({ typeUrl: "/codchen.newchainpoc.dex.MsgMarketSell", value: MsgMarketSell.fromPartial(data) }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
