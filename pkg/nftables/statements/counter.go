package statements

// Counter represents a byte/packet counter.
type Counter struct {
	Packets *int   `json:"packets,omitempty"` // Packets counted.
	Bytes   *int   `json:"bytes,omitempty"`   // Bytes counted.
	Name    string `json:"counter,omitempty"` // Reference to a named counter.
}
