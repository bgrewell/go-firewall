package statements

// FWD represents a packet forwarding configuration.
type FWD struct {
	Dev    string `json:"dev,omitempty"`    // Interface for forwarding.
	Family string `json:"family,omitempty"` // Family of addr.
	Addr   string `json:"addr,omitempty"`   // IP address for forwarding.
}
