package _old

import "fmt"

const (
	matchCommentName   = "comment"
	matchCommentOption = "comment"
)

type MatchComment struct {
	value string
}

func (m MatchComment) Name() string {
	return matchCommentName
}

func (m *MatchComment) SetName(name string) error {
	return fmt.Errorf("comment match doesn't support setting the name")
}

func (m MatchComment) Option() string {
	return matchCommentOption
}

func (m *MatchComment) SetOption(option string) error {
	return fmt.Errorf("comment match doesn't support setting the option")
}

func (m MatchComment) Value() string {
	return m.value
}

func (m *MatchComment) SetValue(value string) error {
	m.value = value
	return nil
}

func (m MatchComment) Negated() bool {
	return false
}

func (m *MatchComment) SetNegated(negated bool) {

}

func (m MatchComment) String() string {
	return fmt.Sprintf("-m %s --%s \"%s\"", m.Name(), m.Option(), m.Value())
}

func (m *MatchComment) Validate(rule Rule) error {
	if m.Value() == "" {
		return fmt.Errorf("required field comment not set")
	}

	return nil
}
