package keeper

import (
	"github.com/emidev98/cosmwasm-lifecycle/x/cosmwasmlifecycle/types"
)

var _ types.QueryServer = Keeper{}
