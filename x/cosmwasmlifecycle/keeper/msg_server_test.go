package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/emidev98/cosmwasm-lifecycle/testutil/keeper"
	"github.com/emidev98/cosmwasm-lifecycle/x/cosmwasmlifecycle/keeper"
	"github.com/emidev98/cosmwasm-lifecycle/x/cosmwasmlifecycle/types"
	"github.com/stretchr/testify/require"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CosmwasmlifecycleKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

func TestMsgServer(t *testing.T) {
	ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
}
