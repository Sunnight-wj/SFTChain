package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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
	TypeMsgMintTo       = "mint_to"
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
	amount sdk.Coin,
	feePercentage float64,
) *MsgUpdateParams {

	return &MsgUpdateParams{
		Authority: authority.String(),
		Params: Params{
			FeePercentage: feePercentage,
			DataPrice:     amount,
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

// NewMsgUpdateParams creates new instance of MsgUpdateParams
func NewMsgMintTo(
	sender sdk.AccAddress,
	amount sdk.Coin,
	to sdk.AccAddress,
) *MsgMintTo {

	return &MsgMintTo{
		Sender:        sender.String(),
		Amount:        amount,
		MintToAddress: to.String(),
	}
}

// Route returns the name of the module
func (msg MsgMintTo) Route() string { return RouterKey }

// Type returns the the action
func (msg MsgMintTo) Type() string { return TypeMsgMintTo }

// ValidateBasic runs stateless checks on the message
func (msg MsgMintTo) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return err
	}
	_, err = sdk.AccAddressFromBech32(msg.MintToAddress)
	if err != nil {
		return err
	}
	if !msg.Amount.IsValid() || msg.Amount.Amount.Equal(sdk.ZeroInt()) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgMintTo) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgMintTo) GetSigners() []sdk.AccAddress {
	from, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{from}
}
