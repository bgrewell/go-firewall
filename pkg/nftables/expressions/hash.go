package expressions

type JhashExpression struct {
	Mod    int         `json:"mod"`
	Offset int         `json:"offset,omitempty"`
	Expr   interface{} `json:"expr"`
	Seed   int         `json:"seed,omitempty"`
}

type SymhashExpression struct {
	Mod    int `json:"mod"`
	Offset int `json:"offset,omitempty"`
}
