package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgLlm = "llm"

var _ sdk.Msg = &MsgLlm{}

func NewMsgLlm(creator string, modality string, model string, prompt string, context string, machine string, endpoint string) *MsgLlm {
	return &MsgLlm{
		Creator:  creator,
		Modality: modality,
		Model:    model,
		Prompt:   prompt,
		Context:  context,
		Machine:  machine,
		Endpoint: endpoint,
	}
}

func (msg *MsgLlm) Route() string {
	return RouterKey
}

func (msg *MsgLlm) Type() string {
	return TypeMsgLlm
}

func (msg *MsgLlm) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgLlm) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgLlm) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
