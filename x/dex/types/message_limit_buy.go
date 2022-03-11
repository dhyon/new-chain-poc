package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgLimitBuy = "limit_buy"

var _ sdk.Msg = &MsgLimitBuy{}

func NewMsgLimitBuy(creator string, price int32, quantity int32) *MsgLimitBuy {
	return &MsgLimitBuy{
		Creator:  creator,
		Price:    price,
		Quantity: quantity,
	}
}

func (msg *MsgLimitBuy) Route() string {
	return RouterKey
}

func (msg *MsgLimitBuy) Type() string {
	return TypeMsgLimitBuy
}

func (msg *MsgLimitBuy) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgLimitBuy) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgLimitBuy) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
