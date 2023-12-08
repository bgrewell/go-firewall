package statements

// Reject represents a packet rejection configuration.
type Reject struct {
	Type string `json:"type,omitempty"` // Type of reject.
	Expr string `json:"expr,omitempty"` // ICMP code to reject with.
}
