package statements

// Mangle represents a change to packet data or meta info.
type Mangle struct {
	Key   string `json:"key"`   // Packet data key to change.
	Value string `json:"value"` // Value to change data to.
}
