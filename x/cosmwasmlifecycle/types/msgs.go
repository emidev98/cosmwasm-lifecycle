package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.Msg = &MsgFundExistentContract{}

	MsgFundExistentContractType = "wasm_lifecycle_msg_fund_existent_contract_type"
)

func NewMsgFundExistentContract(sender, contractAddr string, deposit sdk.Coin) *MsgFundExistentContract {
	return &MsgFundExistentContract{
		Sender:       sender,
		ContractAddr: contractAddr,
		Deposit:      deposit,
	}
}

func (msg MsgFundExistentContract) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgFundExistentContract) Route() string {
	return sdk.MsgTypeURL(&msg)
}

func (msg MsgFundExistentContract) ValidateBasic() error {

	return nil
}

func (msg MsgFundExistentContract) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(ErrorInvalidSigner)
	}
	return []sdk.AccAddress{signer}
}

func (msg MsgFundExistentContract) Type() string { return MsgFundExistentContractType }
