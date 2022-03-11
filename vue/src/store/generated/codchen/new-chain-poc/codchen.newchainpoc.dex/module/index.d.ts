import { StdFee } from "@cosmjs/launchpad";
import { Registry, OfflineSigner, EncodeObject } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgLimitBuy } from "./types/dex/tx";
import { MsgLimitSell } from "./types/dex/tx";
import { MsgMarketBuy } from "./types/dex/tx";
import { MsgMarketSell } from "./types/dex/tx";
export declare const MissingWalletError: Error;
export declare const registry: Registry;
interface TxClientOptions {
    addr: string;
}
interface SignAndBroadcastOptions {
    fee: StdFee;
    memo?: string;
}
declare const txClient: (wallet: OfflineSigner, { addr: addr }?: TxClientOptions) => Promise<{
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }?: SignAndBroadcastOptions) => any;
    msgLimitBuy: (data: MsgLimitBuy) => EncodeObject;
    msgLimitSell: (data: MsgLimitSell) => EncodeObject;
    msgMarketBuy: (data: MsgMarketBuy) => EncodeObject;
    msgMarketSell: (data: MsgMarketSell) => EncodeObject;
}>;
interface QueryClientOptions {
    addr: string;
}
declare const queryClient: ({ addr: addr }?: QueryClientOptions) => Promise<Api<unknown>>;
export { txClient, queryClient, };
