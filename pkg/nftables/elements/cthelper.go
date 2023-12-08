package elements

// CTHelper represents a named connection tracking helper in nftables.
// It includes attributes to describe the properties of the CT helper.
type CTHelper struct {
	Family   string `json:"family"`   // The table's family.
	Table    string `json:"table"`    // The table's name.
	Name     string `json:"name"`     // The CT helper's name.
	Handle   int    `json:"handle"`   // The CT helper's handle, used in delete command.
	Type     string `json:"type"`     // The CT helper type name, e.g., "ftp" or "tftp".
	Protocol string `json:"protocol"` // The CT helper's layer 4 protocol.
	L3Proto  string `json:"l3proto"`  // The CT helper's layer 3 protocol, e.g., "ip" or "ip6".
}
