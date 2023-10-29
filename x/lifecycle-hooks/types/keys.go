package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
)

const (
	// ModuleName defines the module name
	ModuleName = "lifecyclehooks"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_lifecyclehooks"
)

var (
	ParamKey    = []byte{0x01}
	ContractKey = []byte{0x02}
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

// GetContractKeyByAddress returns the key for the contract
func GetContractKeyByAddress(contractAddr sdk.AccAddress) []byte {
	return append(ContractKey, address.MustLengthPrefix(contractAddr)...)
}
