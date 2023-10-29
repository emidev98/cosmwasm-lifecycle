package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

const (
	UpdateParamsProposalType     = "wasm_lifecycle_update_params"
	RegisterContractProposalType = "wasm_lifecycle_register_contract"
	ModifyContractProposalType   = "wasm_lifecycle_modify_contract"
	RemoveContractProposalType   = "wasm_lifecycle_remove_contract"
)

var (
	_ govtypes.Content = &MsgUpdateParamsProposal{}
	_ govtypes.Content = &MsgRegisterContractProposal{}
	_ govtypes.Content = &MsgModifyContractProposal{}
	_ govtypes.Content = &MsgRemoveContractProposal{}

	_ sdk.Msg = &MsgUpdateParamsProposal{}
	_ sdk.Msg = &MsgRegisterContractProposal{}
	_ sdk.Msg = &MsgModifyContractProposal{}
	_ sdk.Msg = &MsgRemoveContractProposal{}
)

func init() {
	govtypes.RegisterProposalType(UpdateParamsProposalType)
	govtypes.RegisterProposalType(RegisterContractProposalType)
	govtypes.RegisterProposalType(ModifyContractProposalType)
	govtypes.RegisterProposalType(RemoveContractProposalType)
}

func NewMsgCreateAllianceProposal(
	title, description, authority string,
	isEnabled bool,
	minDeposit sdk.Coin,
	strikesToDisableExecution int64,
) govtypes.Content {
	return &MsgUpdateParamsProposal{
		Title:       title,
		Description: description,
		Authority:   authority,
		Params: Params{
			IsEnabled:                 isEnabled,
			MinDeposit:                minDeposit,
			StrikesToDisableExecution: strikesToDisableExecution,
		},
	}
}
func (m MsgUpdateParamsProposal) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}
func (m MsgUpdateParamsProposal) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(m.Authority)
	if err != nil {
		panic(ErrorInvalidSigner)
	}
	return []sdk.AccAddress{signer}
}
func (m *MsgUpdateParamsProposal) ProposalRoute() string { return RouterKey }
func (m *MsgUpdateParamsProposal) ProposalType() string  { return UpdateParamsProposalType }
func (m *MsgUpdateParamsProposal) ValidateBasic() error {
	if m.Title == "" {
		return ErrorTitleCannotBeEmptry
	}
	if m.Description == "" {
		return ErrorDescriptionCannotBeEmptry
	}
	if m.Params.MinDeposit.IsZero() {
		return ErrorMinDepositCannotBeZero
	}
	if m.Params.StrikesToDisableExecution <= 0 {
		return ErrorStrikesToDisableExecutionCannotBeNegative
	}
	return nil
}

func NewMsgRegisterContractProposal(
	title, description, authority string,
	contractDeposit sdk.Coin,
	contractAddr string,
	executiontType ExecutionType,
	executionBlocksFrequency int64,
) govtypes.Content {
	return &MsgRegisterContractProposal{
		Title:                    title,
		Description:              description,
		Authority:                authority,
		ContractDeposit:          contractDeposit,
		ContractAddr:             contractAddr,
		ExecutionType:            executiontType,
		ExecutionBlocksFrequency: executionBlocksFrequency,
	}
}
func (m MsgRegisterContractProposal) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}
func (m MsgRegisterContractProposal) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(m.Authority)
	if err != nil {
		panic(ErrorInvalidSigner)
	}
	return []sdk.AccAddress{signer}
}
func (m *MsgRegisterContractProposal) ProposalRoute() string { return RouterKey }
func (m *MsgRegisterContractProposal) ProposalType() string  { return UpdateParamsProposalType }
func (m *MsgRegisterContractProposal) ValidateBasic() error {
	if m.Title == "" {
		return ErrorTitleCannotBeEmptry
	}
	if m.Description == "" {
		return ErrorDescriptionCannotBeEmptry
	}

	if _, err := sdk.AccAddressFromBech32(m.ContractAddr); err != nil {
		return ErrorInvalidContractAddr
	}

	return nil
}

func NewMsgModifyContractProposal(
	title, description, authority string,
	contractAddr string,
	executiontType ExecutionType,
	operation ExecutionTypeOperation,
	executionBlocksFrequency int64,
) govtypes.Content {
	return &MsgModifyContractProposal{
		Title:                    title,
		Description:              description,
		Authority:                authority,
		ContractAddr:             contractAddr,
		ExecutionType:            executiontType,
		Operation:                operation,
		ExecutionBlocksFrequency: executionBlocksFrequency,
	}
}
func (m MsgModifyContractProposal) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}
func (m MsgModifyContractProposal) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(m.Authority)
	if err != nil {
		panic(ErrorInvalidSigner)
	}
	return []sdk.AccAddress{signer}
}
func (m *MsgModifyContractProposal) ProposalRoute() string { return RouterKey }
func (m *MsgModifyContractProposal) ProposalType() string  { return UpdateParamsProposalType }
func (m *MsgModifyContractProposal) ValidateBasic() error {
	if m.Title == "" {
		return ErrorTitleCannotBeEmptry
	}
	if m.Description == "" {
		return ErrorDescriptionCannotBeEmptry
	}

	if _, err := sdk.AccAddressFromBech32(m.ContractAddr); err != nil {
		return ErrorInvalidContractAddr
	}

	return nil
}

func NewMsgRemoveContractProposal(
	title, description, authority string,
	contractDeposit sdk.Coin,
	contractAddr string,
	depositRefundAccount string,
) govtypes.Content {
	return &MsgRemoveContractProposal{
		Title:                title,
		Description:          description,
		Authority:            authority,
		ContractAddr:         contractAddr,
		DepositRefundAccount: depositRefundAccount,
	}
}
func (m MsgRemoveContractProposal) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}
func (m MsgRemoveContractProposal) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(m.Authority)
	if err != nil {
		panic(ErrorInvalidSigner)
	}
	return []sdk.AccAddress{signer}
}
func (m *MsgRemoveContractProposal) ProposalRoute() string { return RouterKey }
func (m *MsgRemoveContractProposal) ProposalType() string  { return UpdateParamsProposalType }
func (m *MsgRemoveContractProposal) ValidateBasic() error {
	if m.Title == "" {
		return ErrorTitleCannotBeEmptry
	}
	if m.Description == "" {
		return ErrorDescriptionCannotBeEmptry
	}

	if _, err := sdk.AccAddressFromBech32(m.ContractAddr); err != nil {
		return ErrorInvalidContractAddr
	}

	if _, err := sdk.AccAddressFromBech32(m.DepositRefundAccount); err != nil {
		return ErrorInvalidDepositRefundAddr
	}

	return nil
}
