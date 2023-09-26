package types

import "strings"

const (
	ModuleName = "datamarket"
)

// KeySeparator is used to combine parts of the keys in the store
const KeySeparator = "|"

var (
	DataSetPrefix = "dataset"
	BuyerPrefix   = "buyer"
)

func GetDataSetPrefix(class string) []byte {
	return []byte(strings.Join([]string{DataSetPrefix, class, ""}, KeySeparator))
}

func GetBuyerPrefix(buyer string) []byte {
	return []byte(strings.Join([]string{BuyerPrefix, buyer, ""}, KeySeparator))
}
