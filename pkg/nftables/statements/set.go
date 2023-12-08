package statements

// Set represents a dynamic set operation.
type Set struct {
	Op   string `json:"op"`   // Operator on set.
	Elem string `json:"elem"` // Set element to add or update.
	Set  string `json:"set"`  // Set reference.
}
