package statements

// Dup duplicates a packet to a different destination.
type Dup struct {
	Addr string `json:"addr,omitempty"` // Address to duplicate packet to.
	Dev  string `json:"dev,omitempty"`  // Interface to duplicate packet on.
}
