package keeper

import (
	"context"

	"github.com/emidev98/cosmwasm-lifecycle/x/cosmwasmlifecycle/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{
		Keeper: keeper,
	}
}

var _ types.MsgServer = msgServer{}

// Enable a contract execution.
func (msgServer msgServer) EnableContractBlockExecution(ctx context.Context, msg *types.MsgEnableContractBlockExecution) (res *types.MsgEnableContractBlockExecutionResponse, err error) {

	return res, err
}

// Disable a contract execution and return the funds to the external account.
func (msgServer msgServer) DisableContractBlockExecution(ctx context.Context, msg *types.MsgDisableContractBlockExecution) (res *types.MsgDisableContractBlockExecutionResponse, err error) {

	return res, err
}
