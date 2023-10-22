package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

const (
	UpdateParamsProposalType             = "wasm_lifecycle_update_params_proposal_type"
	EnableContractExecutionProposalType  = "wasm_lifecycle_enable_contract_execution_proposal_type"
	DisableContractExecutionProposalType = "wasm_lifecycle_disable_contract_execution_proposal_type"
)

var (
	_ govtypes.Content = &MsgUpdateParamsProposal{}
	_ govtypes.Content = &MsgEnableContractExecutionProposal{}
	_ govtypes.Content = &MsgDisableContractExecutionProposal{}
)

func init() {
	govtypes.RegisterProposalType(UpdateParamsProposalType)
	govtypes.RegisterProposalType(EnableContractExecutionProposalType)
	govtypes.RegisterProposalType(DisableContractExecutionProposalType)
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

func NewMsgEnableContractExecutionProposal(
	title, description, authority string,
	contractDeposit sdk.Coin,
	contractAddr string,
	execution ExecutionType,
) govtypes.Content {
	return &MsgEnableContractExecutionProposal{
		Title:           title,
		Description:     description,
		Authority:       authority,
		ContractDeposit: contractDeposit,
		ContractAddr:    contractAddr,
		Execution:       execution,
	}
}
func (m *MsgEnableContractExecutionProposal) ProposalRoute() string { return RouterKey }
func (m *MsgEnableContractExecutionProposal) ProposalType() string  { return UpdateParamsProposalType }
func (m *MsgEnableContractExecutionProposal) ValidateBasic() error {
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

func NewMsgDisableContractExecutionProposal(
	title, description, authority string,
	contractDeposit sdk.Coin,
	contractAddr string,
	execution ExecutionType,
	depositRefundAccount string,
) govtypes.Content {
	return &MsgDisableContractExecutionProposal{
		Title:                title,
		Description:          description,
		Authority:            authority,
		ContractAddr:         contractAddr,
		Execution:            execution,
		DepositRefundAccount: depositRefundAccount,
	}
}
func (m *MsgDisableContractExecutionProposal) ProposalRoute() string { return RouterKey }
func (m *MsgDisableContractExecutionProposal) ProposalType() string  { return UpdateParamsProposalType }
func (m *MsgDisableContractExecutionProposal) ValidateBasic() error {
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
