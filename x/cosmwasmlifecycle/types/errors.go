package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/cosmwasmlifecycle module sentinel errors
var (
	// Message validation errors
	ErrorTitleCannotBeEmptry                       = sdkerrors.Register(ModuleName, 50000, "title cannot be empty")
	ErrorDescriptionCannotBeEmptry                 = sdkerrors.Register(ModuleName, 50001, "description cannot be empty")
	ErrorMinDepositCannotBeZero                    = sdkerrors.Register(ModuleName, 50002, "min deposit cannot be zero")
	ErrorStrikesToDisableExecutionCannotBeNegative = sdkerrors.Register(ModuleName, 50003, "strikes to disable execution cannot be negative")
	ErrorInvalidContractAddr                       = sdkerrors.Register(ModuleName, 50004, "invalid contract address")
	ErrorInvalidDepositRefundAddr                  = sdkerrors.Register(ModuleName, 50005, "invalid deposit refund address")
	ErrorInvalidSigner                             = sdkerrors.Register(ModuleName, 50006, "invalid message signer")

	// Msg Server Errors
	ErrorInvalidAuthority                   = sdkerrors.Register(ModuleName, 60000, "invalid authority on message execution")
	ErrorExecutionTypeAlreadyExists         = sdkerrors.Register(ModuleName, 60001, "execution type already exists")
	ErrorContractNotFoundWithAddress        = sdkerrors.Register(ModuleName, 60002, "contract not found with specified address")
	ErrorContractAlreadyExists              = sdkerrors.Register(ModuleName, 60003, "contract with specified number already exists")
	ErrorInvalidDenom                       = sdkerrors.Register(ModuleName, 60004, "denom invalid denom specified")
	ErrorContractAlreadyEnabled             = sdkerrors.Register(ModuleName, 60005, "contract already enabled with the specified execution type")
	ErrorCannotDisableAllContractExecutions = sdkerrors.Register(ModuleName, 60006, "cannot disable all contract execution, please remove it instead")
)
