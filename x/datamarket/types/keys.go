package types

import "strings"

var ParamsKey = []byte{0x00}

const (
	ModuleName = "datamarket"

	StoreKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// RouterKey to be used for message routing
	RouterKey = ModuleName
)

// KeySeparator is used to combine parts of the keys in the store
const KeySeparator = "|"

var (
	DataSetKey     = "dataset"
	UploaderSetKey = "uploaderset"
	BuyerKey       = "buyer"
	VipInfoKey     = "vipinfo"
	DataSetPrefix  = "dataset"
	BuyerPrefix    = "buyer"
	VipInfoPrefix  = "vipinfo"
)

func GetDataSetPrefix(class string) []byte {
	return []byte(strings.Join([]string{DataSetPrefix, class, ""}, KeySeparator))
}

func GetBuyerPrefix(buyer string) []byte {
	return []byte(strings.Join([]string{BuyerPrefix, buyer, ""}, KeySeparator))
}

func GetVipInfoPrefix() []byte {
	return []byte(strings.Join([]string{VipInfoPrefix}, KeySeparator))
}
