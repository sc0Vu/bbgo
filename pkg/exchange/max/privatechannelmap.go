// Code generated by go:generate; DO NOT EDIT.
package max

var AllPrivateChannels = map[PrivateChannel]struct{}{
	PrivateChannelOrder:                  {},
	PrivateChannelOrderUpdate:            {},
	PrivateChannelTrade:                  {},
	PrivateChannelTradeUpdate:            {},
	PrivateChannelFastTradeUpdate:        {},
	PrivateChannelAccount:                {},
	PrivateChannelAccountUpdate:          {},
	PrivateChannelAveragePrice:           {},
	PrivateChannelFavoriteMarket:         {},
	PrivateChannelMWalletOrder:           {},
	PrivateChannelMWalletTrade:           {},
	PrivateChannelMWalletTradeFastUpdate: {},
	PrivateChannelMWalletAccount:         {},
	PrivateChannelMWalletAveragePrice:    {},
	PrivateChannelBorrowing:              {},
	PrivateChannelAdRatio:                {},
	PrivateChannelPoolQuota:              {},
}
var AllMarginPrivateChannels = map[PrivateChannel]struct{}{
	PrivateChannelMWalletOrder:           {},
	PrivateChannelMWalletTrade:           {},
	PrivateChannelMWalletTradeFastUpdate: {},
	PrivateChannelMWalletAccount:         {},
	PrivateChannelMWalletAveragePrice:    {},
	PrivateChannelBorrowing:              {},
	PrivateChannelAdRatio:                {},
	PrivateChannelPoolQuota:              {},
}

// AllMarginPrivateChannelsKeys converts the Margin group map of PrivateChannel to a slice of PrivateChannel
func AllMarginPrivateChannelsKeys() []PrivateChannel {
	keys := make([]PrivateChannel, 0, len(AllMarginPrivateChannels))
	for k := range AllMarginPrivateChannels {
		keys = append(keys, k)
	}
	return keys
}

// ValidateMarginPrivateChannels validates if a value belongs to the Margin group of PrivateChannel
func ValidateMarginPrivateChannels(ch PrivateChannel) bool {
	_, ok := AllMarginPrivateChannels[ch]
	return ok
}

// IsMarginPrivateChannel checks if the value is in the Margin group of PrivateChannel
func IsMarginPrivateChannel(ch PrivateChannel) bool {
	_, exist := AllMarginPrivateChannels[ch]
	return exist
}

var AllMiscPrivateChannels = map[PrivateChannel]struct{}{
	PrivateChannelAveragePrice:   {},
	PrivateChannelFavoriteMarket: {},
}

// AllMiscPrivateChannelsKeys converts the Misc group map of PrivateChannel to a slice of PrivateChannel
func AllMiscPrivateChannelsKeys() []PrivateChannel {
	keys := make([]PrivateChannel, 0, len(AllMiscPrivateChannels))
	for k := range AllMiscPrivateChannels {
		keys = append(keys, k)
	}
	return keys
}

// ValidateMiscPrivateChannels validates if a value belongs to the Misc group of PrivateChannel
func ValidateMiscPrivateChannels(ch PrivateChannel) bool {
	_, ok := AllMiscPrivateChannels[ch]
	return ok
}

// IsMiscPrivateChannel checks if the value is in the Misc group of PrivateChannel
func IsMiscPrivateChannel(ch PrivateChannel) bool {
	_, exist := AllMiscPrivateChannels[ch]
	return exist
}

var AllPrivateChannelsSlice = []PrivateChannel{
	PrivateChannelOrder,
	PrivateChannelOrderUpdate,
	PrivateChannelTrade,
	PrivateChannelTradeUpdate,
	PrivateChannelFastTradeUpdate,
	PrivateChannelAccount,
	PrivateChannelAccountUpdate,
	PrivateChannelAveragePrice,
	PrivateChannelFavoriteMarket,
	PrivateChannelMWalletOrder,
	PrivateChannelMWalletTrade,
	PrivateChannelMWalletTradeFastUpdate,
	PrivateChannelMWalletAccount,
	PrivateChannelMWalletAveragePrice,
	PrivateChannelBorrowing,
	PrivateChannelAdRatio,
	PrivateChannelPoolQuota,
}

// PrivateChannelStrings converts a slice of PrivateChannel to a slice of string
func PrivateChannelStrings(slice []PrivateChannel) (out []string) {
	for _, el := range slice {
		out = append(out, string(el))
	}
	return out
}

// PrivateChannelKeys converts a map of PrivateChannel to a slice of PrivateChannel
func PrivateChannelKeys(values map[PrivateChannel]struct{}) (slice []PrivateChannel) {
	for k := range values {
		slice = append(slice, k)
	}
	return slice
}

// ValidatePrivateChannel validates a value of type PrivateChannel
func ValidatePrivateChannel(ch PrivateChannel) bool {
	_, ok := AllPrivateChannels[ch]
	return ok
}
