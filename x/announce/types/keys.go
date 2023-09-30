package types

const (
	// ModuleName for the hyperlane-announce
	ModuleName = "hyperlane-announce"

	// StoreKey is the store key string for hyperlane-announce
	StoreKey = ModuleName

	// RouterKey is the message route for hyperlane-announce
	RouterKey = ModuleName
)

// KVStore keys
var (
	AnnouncedValidators       = []byte{0x00}
	AnnouncedStorageLocations = []byte{0x1}
)

func AnnouncedValidatorsKey() []byte {
	return AnnouncedValidators
}

func StorageLocationsKey(val string, messageId []byte) []byte {
	return append(AnnouncedStorageLocations, []byte(val)...)
}
