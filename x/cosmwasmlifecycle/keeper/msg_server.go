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
func (s msgServer) RegisterContract(ctx context.Context, msg *types.MsgRegisterContractProposal) (res *types.MsgRegisterContractProposalResponse, err error) {
	if s.GetAuthority() != msg.Authority {
		return nil, types.ErrorInvalidAuthority
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	_, found := s.Keeper.GetContract(sdkCtx, sdk.MustAccAddressFromBech32(msg.GetContractAddr()))
	if found {
		return nil, types.ErrorContractAlreadyExists
	}

	contract := types.NewCleanContract(msg.GetExecutionType(), msg.GetExecutionBlocksFrequency(), msg.ContractDeposit)
	// store contract
	s.Keeper.SetContract(sdkCtx,
		sdk.AccAddress(msg.GetContractAddr()),
		contract,
	)
	// TODO: emit contract enabled event
	return res, err
}

// Modify a registered contract execution, only the module authority can do this.
func (s msgServer) ModifyContract(ctx context.Context, msg *types.MsgModifyContractProposal) (res *types.MsgModifyContractProposalResponse, err error) {
	if s.GetAuthority() != msg.Authority {
		return nil, types.ErrorInvalidAuthority
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	contract, found := s.Keeper.GetContract(sdkCtx, sdk.MustAccAddressFromBech32(msg.GetContractAddr()))
	if !found {
		return nil, types.ErrorContractNotFoundWithAddress
	}

	// When enabling a smart contract execution type must not be
	// already enabled with the same execution type otherwise
	// return an error.
	if msg.GetOperation() == types.ExecutionTypeOperation_ENABLE {
		if contract.GetExecutionType() == types.ExecutionType_BEGIN_AND_END_BLOCK {
			return nil, types.ErrorContractAlreadyEnabled
		} else if contract.GetExecutionType() == msg.GetExecutionType() {
			return nil, types.ErrorContractAlreadyEnabled
		}
		contract.ExecutionType = types.ExecutionType_BEGIN_AND_END_BLOCK
	} else {

		// When disabling a smart contract execution cannot disable both
		// execution types, otherwise return an error.
		if contract.GetExecutionType() == msg.GetExecutionType() ||
			contract.GetExecutionType() == types.ExecutionType_BEGIN_AND_END_BLOCK {
			return nil, types.ErrorCannotDisableAllContractExecutions
		}

		// Based on the current execution type, set the new execution type.
		if msg.GetExecutionType() == types.ExecutionType_BEGIN_BLOCK {
			contract.ExecutionType = types.ExecutionType_END_BLOCK
		} else {
			contract.ExecutionType = types.ExecutionType_BEGIN_BLOCK
		}
	}
	// Override the execution frequency
	contract.ExecutionFrequency = msg.GetExecutionBlocksFrequency()
	// store contract
	s.Keeper.SetContract(sdkCtx, sdk.MustAccAddressFromBech32(msg.GetContractAddr()), contract)
	// TODO: emit contract modify event

	return res, err
}

// Remove a registered contract execution and return the deposit to the refund account.
func (s msgServer) RemoveContract(ctx context.Context, msg *types.MsgRemoveContractProposal) (res *types.MsgRemoveContractProposalResponse, err error) {
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
	// TODO: emit contract removed event
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
	// store contract
	s.Keeper.SetContract(sdkCtx, sdk.MustAccAddressFromBech32(msg.GetContractAddr()), contract)

	return res, err
}
