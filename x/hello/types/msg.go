package types

import sdk "github.com/cosmos/cosmos-sdk/types"

var _ sdk.Msg = &MsgUpdateParams{}

const TypeMsgUpdateParams = "update_params"

// NewMsgUpdateParams creates new instance of MsgUpdateParams
func NewMsgUpdateParams(
	authority sdk.Address,
	name string,
	age uint64,
) *MsgUpdateParams {

	return &MsgUpdateParams{
		Authority: authority.String(),
		Params: Params{
			Name: name,
			Age:  age,
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
