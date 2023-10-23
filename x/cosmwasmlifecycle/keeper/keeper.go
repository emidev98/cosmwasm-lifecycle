package keeper

import (
	"fmt"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/emidev98/cosmwasm-lifecycle/x/cosmwasmlifecycle/types"
)

type Keeper struct {
	cdc           codec.BinaryCodec
	storeKey      storetypes.StoreKey
	wasmKeeper    types.WasmKeeper
	bankKeeper    types.BankKeeper
	authorityAddr string
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,
	wasmKeeper types.WasmKeeper,
	bankKeeper types.BankKeeper,
	authorityAddr string,
) *Keeper {

	return &Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		wasmKeeper:    wasmKeeper,
		bankKeeper:    bankKeeper,
		authorityAddr: authorityAddr,
	}
}

// Function used to execute the begin_block as sudo of a contract
func (k Keeper) BeginBlock(ctx sdk.Context, currentBlockHeight int64) (err error) {
	// First gets the module params and check if the module is enabled
	params := k.GetParams(ctx)
	if params.IsEnabled {
		// Iterate over all contracts and check if the contract can be executed in the current context
		err = k.IterateContracts(ctx, func(contractAddr sdk.AccAddress, contract types.Contract) error {
			return k.executeContract(ctx, contractAddr, contract, params, currentBlockHeight, types.ExecutionType_BEGIN_BLOCK)
		})

		if err != nil {
			return err
		}
	}

	return nil
}

func (k Keeper) EndBlock(ctx sdk.Context, currentBlockHeight int64) (err error) {
	// First gets the module params and check if the module is enabled
	params := k.GetParams(ctx)
	if params.IsEnabled {
		// Iterate over all contracts and check if the contract can be executed in the current context
		err = k.IterateContracts(ctx, func(contractAddr sdk.AccAddress, contract types.Contract) error {
			return k.executeContract(ctx, contractAddr, contract, params, currentBlockHeight, types.ExecutionType_END_BLOCK)
		})

		if err != nil {
			return err
		}
	}

	return nil
}

// Function used to execute the contract as sudo
func (k Keeper) executeContract(
	ctx sdk.Context,
	contractAddr sdk.AccAddress,
	contract types.Contract,
	params types.Params,
	currentBlockHeight int64,
	currentLifecycle types.ExecutionType,
) error {
	// If the contract has enough strikes it will be removed
	// and the contract deposit will be burned
	if contract.HaveMaxStrikesToPenalize(params.StrikesToDisableExecution) {
		return k.deleteContractAndBurnDeposit(ctx, contractAddr, contract.Deposit)
	}

	// If the contract does not match min deposit, it will not be executed
	if contract.HaveLessThanMinDeposit(params.MinDeposit) {
		return nil
	}

	if contract.CanExecute(currentBlockHeight, currentLifecycle) {
		// Knowing that the contract can be executed
		// the correct method is defined to be executed
		// based on the current execution type
		msgExecute := []byte("{\"sudo_begin_block\": {}}")
		if currentLifecycle == types.ExecutionType_END_BLOCK {
			msgExecute = []byte("{\"sudo_end_block\": {}}")
		}
		// Execute the contract method as sudo and if this execution fails,
		// increment the strikes of the contract and emit the ContractStrikeEvent
		_, err := k.wasmKeeper.Sudo(ctx, contractAddr, msgExecute)
		if err != nil {
			contract.Strikes++

			// Emit a contract strike event
			if err := EmitContractStrikeEvent(ctx, contractAddr, contract.Strikes, err); err != nil {
				return err
			}

			// If the contract has enough strikes it will be removed
			// and the contract deposit will be burned
			if contract.HaveMaxStrikesToPenalize(params.StrikesToDisableExecution) {

				if err := k.deleteContractAndBurnDeposit(ctx, contractAddr, contract.Deposit); err != nil {
					return err
				}
			}
		}
		contract.LatestBlockExecution = currentBlockHeight
		k.SetContract(ctx, contractAddr, contract)
	}

	return nil
}

// Receives the contractAddr and the respective contract deposit
// to delete the contract from the store and burn the deposit
func (k Keeper) deleteContractAndBurnDeposit(ctx sdk.Context, contractAddr sdk.AccAddress, deposit sdk.Coin) (err error) {
	k.DeleteContract(ctx, contractAddr)
	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(deposit))
	if err != nil {
		return err
	}

	err = EmitContractDeleteEvent(ctx, contractAddr)
	if err != nil {
		return err
	}

	return err
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) GetAuthority() string {
	return k.authorityAddr
}
