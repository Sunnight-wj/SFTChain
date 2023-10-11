package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (p Params) Validate() error {
	err := validateDataPrice(p.DataPrice)
	return err
}

func validateDataPrice(i interface{}) error {
	v, ok := i.(sdk.Coin)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.Validate() != nil {
		return fmt.Errorf("invalid denom data price: %+v", i)
	}

	return nil
}
