package statements

// CTTimeout assigns a connection tracking timeout policy.
type CTTimeout struct {
	CTTimeout string `json:"ct timeout"` // CT timeout reference.
}
