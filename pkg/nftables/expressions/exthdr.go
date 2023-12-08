package expressions

type ExthdrExpression struct {
	Name   string `json:"name"`
	Field  string `json:"field,omitempty"`
	Offset int    `json:"offset,omitempty"`
}
