package statements

// Notrack disables connection tracking for the packet.
type Notrack struct {
	Notrack *struct{} `json:"notrack,omitempty"`
}
