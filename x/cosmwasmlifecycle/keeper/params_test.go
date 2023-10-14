package keeper_test

import (
	"testing"

	testkeeper "github.com/emidev98/cosmwasm-lifecycle/testutil/keeper"
	"github.com/emidev98/cosmwasm-lifecycle/x/cosmwasmlifecycle/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CosmwasmlifecycleKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
