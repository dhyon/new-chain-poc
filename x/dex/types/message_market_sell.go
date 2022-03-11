package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgMarketSell = "market_sell"

var _ sdk.Msg = &MsgMarketSell{}

func NewMsgMarketSell(creator string, quantity int32) *MsgMarketSell {
	return &MsgMarketSell{
		Creator:  creator,
		Quantity: quantity,
	}
}

func (msg *MsgMarketSell) Route() string {
	return RouterKey
}

func (msg *MsgMarketSell) Type() string {
	return TypeMsgMarketSell
}

func (msg *MsgMarketSell) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgMarketSell) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMarketSell) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
