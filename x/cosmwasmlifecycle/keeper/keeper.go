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
	cdc        codec.BinaryCodec
	storeKey   storetypes.StoreKey
	wasmKeeper types.WasmKeeper
	bankKeeper types.BankKeeper
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,
	wasmKeeper types.WasmKeeper,
	bankKeeper types.BankKeeper,
) *Keeper {

	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		wasmKeeper: wasmKeeper,
		bankKeeper: bankKeeper,
	}
}

// Function used to execute the begin_block as sudo of a contract
func (k Keeper) BeginBlock(ctx sdk.Context) (err error) {
	// First gets the module params and check if the module is enabled
	params := k.GetParams(ctx)
	if params.IsEnabled {
		// Iterate over all contracts and check if the contract can be executed in the current context
		err = k.IterateContracts(ctx, func(contractAddr sdk.AccAddress, contract types.Contract) error {
			return k.executeContract(ctx, contractAddr, contract, params, types.ExecutionType_BEGIN_BLOCK)
		})

		if err != nil {
			return err
		}
	}

	return nil
}

func (k Keeper) EndBlock(ctx sdk.Context) (err error) {
	// First gets the module params and check if the module is enabled
	params := k.GetParams(ctx)
	if params.IsEnabled {
		// Iterate over all contracts and check if the contract can be executed in the current context
		err = k.IterateContracts(ctx, func(contractAddr sdk.AccAddress, contract types.Contract) error {
			return k.executeContract(ctx, contractAddr, contract, params, types.ExecutionType_END_BLOCK)
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
	executionType types.ExecutionType,
) error {
	if contract.CanExecute(params, executionType) {
		// Execute the begin_block as sudo and if this execution fails,
		// increment the strikes of the contract and emit the ContractStrikeEvent
		_, err := k.wasmKeeper.Sudo(ctx, contractAddr, []byte("{\"begin_block\": {}}"))
		if err != nil {
			contract.Strikes++
			err := EmitContractStrikeEvent(ctx, contractAddr, contract.Strikes, err)
			if err != nil {
				panic(err)
			}

			// If the contract has enough strikes to be disabled,
			// it will be deleted form the store and the deposit
			// will be burned
			if contract.HasEnoughStrikesToDisableExecution(params.StrikesToDisableExecution) {
				err := k.deleteContractAndBurnDeposit(ctx, contractAddr, contract.Deposit)
				if err != nil {
					return err
				}
			} else {
				k.SetContract(ctx, contractAddr, contract)
			}
		}
	} else {
		err := k.deleteContractAndBurnDeposit(ctx, contractAddr, contract.Deposit)
		if err != nil {
			return err
		}
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
