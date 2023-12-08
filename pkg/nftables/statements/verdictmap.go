package statements

// VerdictMap applies a verdict conditionally.
type VerdictMap struct {
	Key  string `json:"key"`  // Map key.
	Data string `json:"data"` // Mapping expression.
}
