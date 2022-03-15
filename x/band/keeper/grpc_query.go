package keeper

import (
	"github.com/codchen/new-chain-poc/x/band/types"
)

var _ types.QueryServer = Keeper{}
