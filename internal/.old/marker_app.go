package _old

import (
	"fmt"
	"strings"
)

type MarkerGeneric struct {
	name  string
	value string
}

func (m MarkerGeneric) Name() string {
	return m.name
}

func (m *MarkerGeneric) SetName(name string) error {
	m.name = name
	return nil
}

func (m MarkerGeneric) Value() string {
	return m.value
}

func (m *MarkerGeneric) SetValue(value string) error {
	m.value = value
	return nil
}

func (m *MarkerGeneric) Parse(comment string) error {
	parts := strings.Split(comment, ":")
	if len(parts) != 2 {
		return fmt.Errorf("invalid marker format. expected <name>:<value> got %s", comment)
	}
	err := m.SetName(parts[0])
	if err != nil {
		return err
	}
	err = m.SetValue(parts[1])
	if err != nil {
		return err
	}
	return nil
}

func (m MarkerGeneric) String() string {
	return fmt.Sprintf("%s:%s", m.Name(), m.Value())
}
