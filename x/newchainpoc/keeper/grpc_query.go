package keeper

import (
	"github.com/codchen/new-chain-poc/x/newchainpoc/types"
)

var _ types.QueryServer = Keeper{}
