package types

import "strings"

var ParamsKey = []byte{0x00}

const (
	ModuleName = "datamarket"
)

// KeySeparator is used to combine parts of the keys in the store
const KeySeparator = "|"

var (
	DataSetKey    = "dataset"
	DataSetPrefix = "dataset"
	BuyerPrefix   = "buyer"
)

func GetDataSetPrefix(class string) []byte {
	return []byte(strings.Join([]string{DataSetPrefix, class, ""}, KeySeparator))
}

func GetBuyerPrefix(buyer string) []byte {
	return []byte(strings.Join([]string{BuyerPrefix, buyer, ""}, KeySeparator))
}
