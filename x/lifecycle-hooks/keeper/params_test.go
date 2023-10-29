package keeper_test

import (
	"testing"

	testkeeper "github.com/emidev98/lifecycle-hooks/testutil/keeper"
	"github.com/emidev98/lifecycle-hooks/x/lifecycle-hooks/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.LifecycleHooksKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
