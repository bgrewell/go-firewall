package expressions

type BinaryOperationExpression struct {
	Operator string        `json:"operator"`
	Operands []interface{} `json:"operands"`
}
