package types

var ParamsKey = []byte{0x00}

const (
	ModuleName = "hello"

	StoreKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// RouterKey to be used for message routing
	RouterKey = ModuleName
)
