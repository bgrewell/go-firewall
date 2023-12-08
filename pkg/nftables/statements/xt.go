package statements

// XT represents an xt statement from xtables compat interface.
type XT struct {
	Type string `json:"type"` // Type of xt statement.
	Name string `json:"name"` // XT name.
}
