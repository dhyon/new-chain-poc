package keeper

import (
	"github.com/codchen/new-chain-poc/x/dex/types"
)

var _ types.QueryServer = Keeper{}
