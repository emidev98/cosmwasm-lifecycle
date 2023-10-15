package types

const (
	// ModuleName defines the module name
	ModuleName = "cosmwasmlifecycle"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_cosmwasmlifecycle"
)

var (
	ParamsKey  = []byte{0x01}
	StrikesKey = []byte{0x02}
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
