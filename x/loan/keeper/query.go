package keeper

import (
	"simpleloan/x/loan/types"
)

var _ types.QueryServer = Keeper{}
