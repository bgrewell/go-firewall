package statements

// NAT represents a Network Address Translation configuration.
type NAT struct {
	SNAT       *NATConfig `json:"snat,omitempty"`
	DNAT       *NATConfig `json:"dnat,omitempty"`
	Masquerade *NATConfig `json:"masquerade,omitempty"`
	Redirect   *NATConfig `json:"redirect,omitempty"`
}

// NATConfig holds the configuration for a NAT rule.
type NATConfig struct {
	Addr   string   `json:"addr,omitempty"`   // Address to translate to.
	Family string   `json:"family,omitempty"` // Family of addr.
	Port   string   `json:"port,omitempty"`   // Port to translate to.
	Flags  []string `json:"flags,omitempty"`  // Flags for NAT.
}
