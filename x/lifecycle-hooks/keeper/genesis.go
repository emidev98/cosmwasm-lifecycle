package keeper

import (
	"cosmossdk.io/errors"
	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/emidev98/lifecycle-hooks/x/lifecycle-hooks/types"
)

func (k Keeper) InitGenesis(ctx sdk.Context, g types.GenesisState) []abci.ValidatorUpdate {
	k.SetParams(ctx, g.Params)

	for _, val := range g.Contracts {
		contractAddress, err := sdk.AccAddressFromBech32(val.ContractAddress)
		if err != nil {
			panic(errors.Wrap(types.ErrorInvalidContractAddr, val.ContractAddress))
		}

		k.SetContract(ctx, contractAddress, val.Contract)
	}

	return []abci.ValidatorUpdate{}
}

func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	k.IterateContracts(ctx, func(contractAddr sdk.AccAddress, contract types.Contract) error {
		genesisContract := types.NewGenesisContract(contractAddr, contract)
		genesis.Contracts = append(genesis.Contracts, &genesisContract)
		return nil
	})
	return genesis
}
