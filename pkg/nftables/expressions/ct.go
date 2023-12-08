package expressions

type CTExpression struct {
	Key    string `json:"key"`
	Family string `json:"family,omitempty"`
	Dir    string `json:"dir,omitempty"`
}
