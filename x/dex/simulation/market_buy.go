package simulation

import (
	"math/rand"

	"github.com/codchen/new-chain-poc/x/dex/keeper"
	"github.com/codchen/new-chain-poc/x/dex/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgMarketBuy(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgMarketBuy{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the MarketBuy simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "MarketBuy simulation not implemented"), nil, nil
	}
}
