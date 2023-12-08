package expressions

type NumgenExpression struct {
	Mode   string `json:"mode"`
	Mod    int    `json:"mod"`
	Offset int    `json:"offset,omitempty"`
}
