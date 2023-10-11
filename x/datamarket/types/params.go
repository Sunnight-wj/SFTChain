package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (p Params) Validate() error {
	err := validateDataPrice(p.DataPrice)
	return err
}

func DefaultParams() Params {
	return Params{
		DataPrice:     sdk.NewInt64Coin("ujuno", 100),
		FeePercentage: 0.1,
	}
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
