package types

import "strings"

const (
	ModuleName = "datamarket"
)

// KeySeparator is used to combine parts of the keys in the store
const KeySeparator = "|"

var (
	ClassPrefix = "class"
	BuyerPrefix = "buyer"
)

func GetClassPrefix(class string) []byte {
	return []byte(strings.Join([]string{ClassPrefix, class, ""}, KeySeparator))
}

func GetClassesPrefix() []byte {
	return []byte(strings.Join([]string{ClassPrefix, ""}, KeySeparator))
}

func GetBuyerPrefix(buyer string) []byte {
	return []byte(strings.Join([]string{BuyerPrefix, buyer, ""}, KeySeparator))
}
