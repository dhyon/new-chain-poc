package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgLimitSell = "limit_sell"

var _ sdk.Msg = &MsgLimitSell{}

func NewMsgLimitSell(creator string, price int32, quantity int32) *MsgLimitSell {
	return &MsgLimitSell{
		Creator:  creator,
		Price:    price,
		Quantity: quantity,
	}
}

func (msg *MsgLimitSell) Route() string {
	return RouterKey
}

func (msg *MsgLimitSell) Type() string {
	return TypeMsgLimitSell
}

func (msg *MsgLimitSell) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgLimitSell) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgLimitSell) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
