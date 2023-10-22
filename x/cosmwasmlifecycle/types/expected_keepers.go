package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type WasmKeeper interface {
	// Sudo allows privileged access to a contract. This can never be called by an external tx, but only by
	// another native Go module directly, or on-chain governance (if sudo proposals are enabled). Thus, the keeper doesn't
	// place any access controls on it, that is the responsibility or the app developer (who passes the wasm.Keeper in app.go)
	//
	// Sub-messages returned from the sudo call to the contract are executed with the default authorization policy. This can be
	// customized though by passing a new policy with the context. See types.WithSubMsgAuthzPolicy.
	// The policy will be read in msgServer.selectAuthorizationPolicy and used for sub-message executions.
	// This is an extension point for some very advanced scenarios only. Use with care!
	Sudo(ctx sdk.Context, contractAddress sdk.AccAddress, msg []byte) ([]byte, error)
}

type BankKeeper interface {
	// BurnCoins burns coins deletes coins from the balance of the module account.
	// It will panic if the module account does not exist or is unauthorized.
	BurnCoins(ctx sdk.Context, name string, amt sdk.Coins) error
}
