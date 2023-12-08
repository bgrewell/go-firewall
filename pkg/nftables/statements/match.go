package statements

// Match represents a comparison between two expressions.
type Match struct {
	Left  string `json:"left"`  // Left-hand side expression.
	Right string `json:"right"` // Right-hand side expression.
	Op    string `json:"op"`    // Comparison operator.
}
