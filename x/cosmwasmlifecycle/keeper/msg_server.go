package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
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

// Update module params, only the module authority can do this.
func (s msgServer) UpdateParams(ctx context.Context, msg *types.MsgUpdateParamsProposal) (res *types.MsgUpdateParamsProposalResponse, err error) {
	if s.GetAuthority() != msg.Authority {
		return nil, types.ErrorInvalidAuthority
	}
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	s.Keeper.SetParams(sdkCtx, msg.Params)
	return res, err
}

// Enable a contract execution, only the module authority can do this.
func (s msgServer) EnableContractExecution(ctx context.Context, msg *types.MsgEnableContractExecutionProposal) (res *types.MsgEnableContractExecutionProposalResponse, err error) {
	if s.GetAuthority() != msg.Authority {
		return nil, types.ErrorInvalidAuthority
	}
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	contract, found := s.Keeper.GetContract(sdkCtx, sdk.MustAccAddressFromBech32(msg.GetContractAddr()))

	if !found {
		s.Keeper.SetContract(sdkCtx,
			sdk.AccAddress(msg.GetContractAddr()),
			types.NewCleanContract(msg.GetExecution(), msg.ContractDeposit),
		)
	} else {
		if contract.ExecutionType == msg.GetExecution() {
			return nil, types.ErrorExecutionTypeAlreadyExists
		}
		contract.ExecutionType = types.ExecutionType_BEGIN_AND_END_BLOCK
		s.Keeper.SetContract(sdkCtx,
			sdk.MustAccAddressFromBech32(msg.GetContractAddr()),
			contract,
		)
	}

	return res, err
}

// Disable a contract execution and return the funds to the external account, only the module authority can do this.
func (s msgServer) DisableContractExecution(ctx context.Context, msg *types.MsgDisableContractExecutionProposal) (res *types.MsgDisableContractExecutionProposalResponse, err error) {
	if s.GetAuthority() != msg.Authority {
		return nil, types.ErrorInvalidAuthority
	}
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	contract, found := s.Keeper.GetContract(sdkCtx, sdk.MustAccAddressFromBech32(msg.GetContractAddr()))

	if !found {
		return nil, types.ErrorContractNotFoundWithAddress
	}

	s.Keeper.DeleteContract(sdkCtx, sdk.AccAddress(msg.GetContractAddr()))
	err = s.bankKeeper.SendCoinsFromModuleToAccount(
		sdkCtx,
		types.ModuleName,
		sdk.AccAddress(msg.DepositRefundAccount),
		sdk.NewCoins(contract.Deposit),
	)
	if err != nil {
		return nil, err
	}

	return res, err
}

// Fund existent contract execution.
func (s msgServer) FundExistentContract(ctx context.Context, msg *types.MsgFundExistentContract) (res *types.MsgFundExistentContractResponse, err error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	contract, found := s.Keeper.GetContract(sdkCtx, sdk.MustAccAddressFromBech32(msg.GetContractAddr()))
	if !found {
		return nil, types.ErrorContractNotFoundWithAddress
	}
	params := s.Keeper.GetParams(sdkCtx)
	if msg.Deposit.Denom != params.MinDeposit.Denom {
		return nil, types.ErrorInvalidDenom
	}
	contract.Deposit = contract.Deposit.Add(msg.Deposit)
	s.Keeper.SetContract(sdkCtx, sdk.MustAccAddressFromBech32(msg.GetContractAddr()), contract)

	return res, err
}
