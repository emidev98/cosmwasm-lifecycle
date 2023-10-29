package keeper

import (
	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/emidev98/lifecycle-hooks/x/lifecycle-hooks/types"
)

func (k Keeper) InitGenesis(ctx sdk.Context, g types.GenesisState) []abci.ValidatorUpdate {
	k.SetParams(ctx, g.Params)
	// TODO: add contract to stores from genesis
	return []abci.ValidatorUpdate{}
}

func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	// TODO: export contracts from stores to genesis
	return genesis
}
