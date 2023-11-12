package consts

// ChainType represents the types of chains supported by nftables
type ChainType string

// Constants relating to the chain types
const (
	// ChainTypeBase is an entry point for packets from the network stack, where a hook is specified
	ChainTypeBase ChainType = "base"
	// ChainTypeRegular is a jump target used for rule organization
	ChainTypeRegular ChainType = "regular"
)
