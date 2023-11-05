package types

import sdk "github.com/cosmos/cosmos-sdk/types"

// this line is used by starport scaffolding # genesis/types/import

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// this line is used by starport scaffolding # genesis/types/default
		Params:    DefaultParams(),
		Contracts: []*GenesisContract{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}

// NewGenesisState creates a new genesis state.
func NewGenesisContract(contractAddress sdk.AccAddress, contract Contract) GenesisContract {
	return GenesisContract{
		ContractAddress: contractAddress.String(),
		Contract:        contract,
	}
}
