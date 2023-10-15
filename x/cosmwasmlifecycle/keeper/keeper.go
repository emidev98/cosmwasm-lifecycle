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
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,
	wasmKeeper types.WasmKeeper,
) *Keeper {

	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		wasmKeeper: wasmKeeper,
	}
}

func (k Keeper) BeginBlock(ctx sdk.Context) {
	params := k.GetParams(ctx)
	if params.IsEnabled {
		for _, address := range params.BeginBlockExecution {
			addr, err := sdk.AccAddressFromBech32(address)
			if err != nil {
				panic(err)
			}

			_, err = k.wasmKeeper.Sudo(ctx, addr, []byte("{\"begin_block\": {}}"))
			if err != nil {

			}
		}
	}
}

func (k Keeper) EndBlock(ctx sdk.Context) {

}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
