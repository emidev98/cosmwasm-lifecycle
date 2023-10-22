package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/emidev98/cosmwasm-lifecycle/x/cosmwasmlifecycle/types"
)

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

// Emits an event to notify that the contract has been striked
// including the address, current strike and the reason
func EmitContractDeleteEvent(
	ctx sdk.Context,
	contractAddress sdk.AccAddress,
) error {
	return ctx.EventManager().EmitTypedEvent(&types.ContractDeleteEvent{
		ModuleName:      types.ModuleName,
		ContractAddress: contractAddress.String(),
	})
}
