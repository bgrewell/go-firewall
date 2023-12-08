package expressions

type ElemExpression struct {
	Val     interface{} `json:"val"`
	Timeout int         `json:"timeout,omitempty"`
	Expires int         `json:"expires,omitempty"`
	Comment string      `json:"comment,omitempty"`
}
