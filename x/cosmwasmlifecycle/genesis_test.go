package cosmwasmlifecycle_test

import (
	"testing"

	keepertest "github.com/emidev98/cosmwasm-lifecycle/testutil/keeper"
	"github.com/emidev98/cosmwasm-lifecycle/testutil/nullify"
	"github.com/emidev98/cosmwasm-lifecycle/x/cosmwasmlifecycle"
	"github.com/emidev98/cosmwasm-lifecycle/x/cosmwasmlifecycle/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CosmwasmlifecycleKeeper(t)
	cosmwasmlifecycle.InitGenesis(ctx, *k, genesisState)
	got := cosmwasmlifecycle.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
