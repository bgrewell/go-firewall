package elements

// CTExpectation represents a named connection tracking expectation in nftables.
// It includes attributes for defining the expected behavior of connections.
type CTExpectation struct {
	Family   string `json:"family"`   // The table's family.
	Table    string `json:"table"`    // The table's name.
	Name     string `json:"name"`     // The CT expectation object's name.
	Handle   int    `json:"handle"`   // The CT expectation object's handle, used in delete command.
	L3Proto  string `json:"l3proto"`  // The CT expectation object's layer 3 protocol.
	Protocol string `json:"protocol"` // The CT expectation object's layer 4 protocol.
	Dport    int    `json:"dport"`    // The destination port of the expected connection.
	Timeout  int    `json:"timeout"`  // The time in milliseconds that this expectation will live.
	Size     int    `json:"size"`     // The maximum count of expectations to be living at the same time.
}
