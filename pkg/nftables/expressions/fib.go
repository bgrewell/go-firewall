package expressions

type FibExpression struct {
	Result string   `json:"result"`
	Flags  []string `json:"flags,omitempty"`
}
