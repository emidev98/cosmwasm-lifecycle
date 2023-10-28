package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/emidev98/cosmwasm-lifecycle/x/cosmwasmlifecycle/types"
)

// Emits an event to notify that a contract has been
// registerd for execution on begin and end block
func EmitRegisterContractEvent(
	ctx sdk.Context,
	contractAddress sdk.AccAddress,
	contractDeposit sdk.Coin,
	executionType types.ExecutionType,
	blockFrequency int64,
) error {
	return ctx.EventManager().EmitTypedEvent(&types.RegisterContractEvent{
		ModuleName:      types.ModuleName,
		ContractAddress: contractAddress.String(),
		ContractDeposit: contractDeposit,
		ExecutionType:   executionType,
		BlockFrequency:  blockFrequency,
	})
}

// Emits an event to notify that a contract has been
// mondified its execution on begin and end block
func EmitModifyContractEvent(
	ctx sdk.Context,
	contractAddress sdk.AccAddress,
	executionType types.ExecutionType,
	blockFrequency int64,
) error {
	return ctx.EventManager().EmitTypedEvent(&types.ModifyContractEvent{
		ModuleName:        types.ModuleName,
		ContractAddress:   contractAddress.String(),
		NewExecutionType:  executionType,
		NewBlockFrequency: blockFrequency,
	})
}

// Emits an event to notify that a contract has been
// remove from its execution on begin and end block
func EmitRemoveContractEvent(
	ctx sdk.Context,
	contractAddress sdk.AccAddress,
	refundAccount sdk.AccAddress,
	refundAmount sdk.Coin,
) error {
	return ctx.EventManager().EmitTypedEvent(&types.RemoveContractEvent{
		ModuleName:      types.ModuleName,
		ContractAddress: contractAddress.String(),
		RefundAccount:   refundAccount.String(),
		RefundAmount:    refundAmount,
	})
}

// Emits an event to notify that the contract has been striked
// including the address, current strike and the reason
func EmitFundExistentContractEvent(
	ctx sdk.Context,
	contractAddress sdk.AccAddress,
	senderAddress sdk.AccAddress,
	depositAmount sdk.Coin,
) error {
	return ctx.EventManager().EmitTypedEvent(&types.FundExistentContractEvent{
		ModuleName:      types.ModuleName,
		ContractAddress: contractAddress.String(),
		SenderAddress:   senderAddress.String(),
		DepositAmount:   depositAmount,
	})
}

// Emits an event to notify that the contract has been striked
// including the address, current strike and the reason
func EmitContractStrikeEvent(
	ctx sdk.Context,
	contractAddress sdk.AccAddress,
	currentStrike int64,
	strikeReason error,
) error {
	return ctx.EventManager().EmitTypedEvent(&types.ContractStrikeEvent{
		ModuleName:      types.ModuleName,
		ContractAddress: contractAddress.String(),
		CurrentStrike:   currentStrike,
		StrikeReason:    strikeReason.Error(),
	})
}

// Emits an event to notify that the contract has received the
// last strike that make it be removed from the state execution
func EmitForceRemoveContractEvent(
	ctx sdk.Context,
	contractAddress sdk.AccAddress,
) error {
	return ctx.EventManager().EmitTypedEvent(&types.ForceRemoveContractEvent{
		ModuleName:      types.ModuleName,
		ContractAddress: contractAddress.String(),
	})
}
