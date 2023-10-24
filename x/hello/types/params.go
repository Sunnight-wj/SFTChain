package types

func (p Params) Validate() error {
	return nil
}

func DefaultParams() Params {
	return Params{
		Name: "xiaoming",
		Age:  24,
	}
}
