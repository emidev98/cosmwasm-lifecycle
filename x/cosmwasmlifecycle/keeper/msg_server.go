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

// CreateContract defines a method for creating a contract
func (msgServer msgServer) EnableBlockExecution(ctx context.Context, msg *types.MsgEnableBlockExecution) (res *types.MsgEnableBlockExecutionResponse, err error) {

	return res, err
}

// CreateContract defines a method for creating a contract
func (msgServer msgServer) DisableBlockExecution(ctx context.Context, msg *types.MsgDisableBlockExecution) (res *types.MsgDisableBlockExecutionResponse, err error) {

	return res, err
}
