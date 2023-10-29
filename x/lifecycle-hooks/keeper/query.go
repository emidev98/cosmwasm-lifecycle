package keeper

import (
	"github.com/emidev98/lifecycle-hooks/x/lifecycle-hooks/types"
)

var _ types.QueryServer = Keeper{}
