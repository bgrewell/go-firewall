package expressions

type PayloadExpression struct {
	Base     string `json:"base,omitempty"`
	Offset   int    `json:"offset,omitempty"`
	Len      int    `json:"len,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	Field    string `json:"field,omitempty"`
}
