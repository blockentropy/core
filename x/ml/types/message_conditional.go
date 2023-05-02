package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgConditional = "conditional"

var _ sdk.Msg = &MsgConditional{}

func NewMsgConditional(creator string, modality string, model string, ctrlmodel string, ctrlinput string, prompt string, negprompt string, seed string, machine string, endpoint string) *MsgConditional {
	return &MsgConditional{
		Creator:   creator,
		Modality:  modality,
		Model:     model,
		Ctrlmodel: ctrlmodel,
		Ctrlinput: ctrlinput,
		Prompt:    prompt,
		Negprompt: negprompt,
		Seed:      seed,
		Machine:   machine,
		Endpoint:  endpoint,
	}
}

func (msg *MsgConditional) Route() string {
	return RouterKey
}

func (msg *MsgConditional) Type() string {
	return TypeMsgConditional
}

func (msg *MsgConditional) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgConditional) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgConditional) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
