package elements

// Element represents an element in a named set in nftables.
// It includes properties to identify the set and the elements to manipulate.
type Element struct {
	Family string   `json:"family"` // The table's family.
	Table  string   `json:"table"`  // The table's name.
	Name   string   `json:"name"`   // The set's name.
	Elem   []string `json:"elem"`   // The elements in the set, expressed as expressions or a list of expressions.
}
