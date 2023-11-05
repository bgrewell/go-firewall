package _old

type Match interface {
	Name() string
	SetName(name string) error
	Option() string
	SetOption(option string) error
	Value() string
	SetValue(value string) error
	Negated() bool
	SetNegated(negated bool)
	String() string
	Validate(rule Rule) error
}
