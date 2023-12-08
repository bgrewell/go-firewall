package statements

// Meter represents a metering configuration.
type Meter struct {
	Name string `json:"name"` // Meter name.
	Key  string `json:"key"`  // Meter key.
	Stmt string `json:"stmt"` // Meter statement.
}
