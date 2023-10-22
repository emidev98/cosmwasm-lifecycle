package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/emidev98/cosmwasm-lifecycle/x/cosmwasmlifecycle/types"
)

func (k Keeper) GetContract(ctx sdk.Context, contractAddr sdk.AccAddress) (d types.Contract, found bool) {
	key := types.GetContractKeyByAddress(contractAddr)
	b := ctx.KVStore(k.storeKey).Get(key)
	if b == nil {
		return d, false
	}
	k.cdc.MustUnmarshal(b, &d)
	return d, true
}

func (k Keeper) SetContract(ctx sdk.Context, contractAddr sdk.AccAddress, contract types.Contract) {
	key := types.GetContractKeyByAddress(contractAddr)
	b := k.cdc.MustMarshal(&contract)
	ctx.KVStore(k.storeKey).Set(key, b)
}

func (k Keeper) DeleteContract(ctx sdk.Context, contractAddr sdk.AccAddress) {
	key := types.GetContractKeyByAddress(contractAddr)
	store := ctx.KVStore(k.storeKey)
	store.Delete(key)
}

func (k Keeper) IterateContracts(ctx sdk.Context, cb func(contractAddr sdk.AccAddress, contract types.Contract) error) (err error) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.ContractKey)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var contract types.Contract
		b := iter.Value()
		k.cdc.MustUnmarshal(b, &contract)
		err = cb(iter.Key(), contract)
		if err != nil {
			return err
		}
	}

	return nil
}
