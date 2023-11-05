package _old

import "fmt"

const ()

func NewMatchGeneric(name string, option string, value string, negated bool) *MatchGeneric {
	m := &MatchGeneric{}
	m.SetName(name)
	m.SetOption(option)
	m.SetValue(value)
	m.SetNegated(negated)
	return m
}

type MatchGeneric struct {
	name    string
	option  string
	value   string
	negated bool
}

func (m MatchGeneric) Name() string {
	return m.name
}

func (m *MatchGeneric) SetName(name string) error {
	m.name = name
	return nil
}

func (m MatchGeneric) Option() string {
	return m.option
}

func (m *MatchGeneric) SetOption(option string) error {
	m.option = option
	return nil
}

func (m MatchGeneric) Value() string {
	return m.value
}

func (m *MatchGeneric) SetValue(value string) error {
	m.value = value
	return nil
}

func (m MatchGeneric) Negated() bool {
	return m.negated
}

func (m *MatchGeneric) SetNegated(negated bool) {
	m.negated = negated
}

func (m MatchGeneric) String() string {
	return fmt.Sprintf("-m %s --%s \"%s\"", m.Name(), m.Option(), m.Value())
}

func (m *MatchGeneric) Validate(rule Rule) error {
	// Generic doesn't have any validation
	return nil
}
