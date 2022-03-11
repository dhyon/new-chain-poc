package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgMarketBuy = "market_buy"

var _ sdk.Msg = &MsgMarketBuy{}

func NewMsgMarketBuy(creator string, quantity int32) *MsgMarketBuy {
	return &MsgMarketBuy{
		Creator:  creator,
		Quantity: quantity,
	}
}

func (msg *MsgMarketBuy) Route() string {
	return RouterKey
}

func (msg *MsgMarketBuy) Type() string {
	return TypeMsgMarketBuy
}

func (msg *MsgMarketBuy) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgMarketBuy) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMarketBuy) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
