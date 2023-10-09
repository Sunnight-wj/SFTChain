package types

import (
	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.Msg = &MsgUploadData{}
	_ sdk.Msg = &MsgBuyData{}
	_ sdk.Msg = &MsgUpdateParams{}
)

const (
	TypeMsgUploadData   = "upload_data"
	TypeMsgBuyData      = "buy_data"
	TypeMsgUpdateParams = "update_params"
)

func NewMsgUploadData(
	uploader sdk.AccAddress,
	class,
	url string,
) *MsgUploadData {
	return &MsgUploadData{
		Uploader: uploader.String(),
		Class:    class,
		Url:      url,
	}
}

// Route returns the name of the module
func (msg MsgUploadData) Route() string { return RouterKey }

// Type returns the the action
func (msg MsgUploadData) Type() string { return TypeMsgUploadData }

// ValidateBasic runs stateless checks on the message
func (msg MsgUploadData) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Uploader); err != nil {
		return errorsmod.Wrapf(err, "invalid upolader address %s", msg.Uploader)
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgUploadData) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgUploadData) GetSigners() []sdk.AccAddress {
	from, _ := sdk.AccAddressFromBech32(msg.Uploader)
	return []sdk.AccAddress{from}
}

// NewMsgBuyData creates new instance of MsgBuyData
func NewMsgBuyData(
	buyer sdk.AccAddress,
	class string,
) *MsgBuyData {

	return &MsgBuyData{
		Buyer: buyer.String(),
		Class: class,
	}
}

// Route returns the name of the module
func (msg MsgBuyData) Route() string { return RouterKey }

// Type returns the the action
func (msg MsgBuyData) Type() string { return TypeMsgBuyData }

// ValidateBasic runs stateless checks on the message
func (msg MsgBuyData) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Buyer); err != nil {
		return errorsmod.Wrapf(err, "invalid buyer address %s", msg.Buyer)
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgBuyData) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgBuyData) GetSigners() []sdk.AccAddress {
	from, _ := sdk.AccAddressFromBech32(msg.Buyer)
	return []sdk.AccAddress{from}
}

// NewMsgUpdateParams creates new instance of MsgUpdateParams
func NewMsgUpdateParams(
	authority sdk.Address,
	denom string,
	amount int64,
	feePercentage float64,
) *MsgUpdateParams {

	dataPrice := &sdk.Coin{
		Denom:  denom,
		Amount: math.NewInt(amount),
	}

	return &MsgUpdateParams{
		Authority: authority.String(),
		Params: &Params{
			FeePercentage: feePercentage,
			DataPrice:     dataPrice,
		},
	}
}

// Route returns the name of the module
func (msg MsgUpdateParams) Route() string { return RouterKey }

// Type returns the the action
func (msg MsgUpdateParams) Type() string { return TypeMsgUpdateParams }

// ValidateBasic runs stateless checks on the message
func (msg MsgUpdateParams) ValidateBasic() error {

	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgUpdateParams) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgUpdateParams) GetSigners() []sdk.AccAddress {
	from, _ := sdk.AccAddressFromBech32(msg.Authority)
	return []sdk.AccAddress{from}
}
