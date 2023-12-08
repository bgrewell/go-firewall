package elements

// CTTimeout represents a named connection tracking timeout policy in nftables.
// It includes attributes for defining the timeout policy for various connection states.
type CTTimeout struct {
	Family   string `json:"family"`   // The table's family.
	Table    string `json:"table"`    // The table's name.
	Name     string `json:"name"`     // The CT timeout object's name.
	Handle   int    `json:"handle"`   // The CT timeout object's handle, used in delete command.
	Protocol string `json:"protocol"` // The CT timeout object's layer 4 protocol.
	State    string `json:"state"`    // The connection state name for which the timeout value is updated.
	Value    int    `json:"value"`    // The updated timeout value for the specified connection state.
	L3Proto  string `json:"l3proto"`  // The CT timeout object's layer 3 protocol.
}
