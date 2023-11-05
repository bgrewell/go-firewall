package _old

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	TargetMarkStr string = "--set-mark"
)

type TargetMark struct {
	Value int `json:"value" yaml:"value" xml:"value"`
}

func (t TargetMark) String() string {
	return TargetJump{
		Value: fmt.Sprintf("MARK %s %d", TargetMarkStr, t.Value),
	}.String()
}

// Returns if the target is valid when applied with the specified rule
func (t TargetMark) Validate(rule Rule) error {
	// No validation needed?
	return nil
}

func (t *TargetMark) Parse(option string, value string) {
	if option == "--set-mask" {
		value = strings.Replace(value, "0x", "", -1)
		value = strings.Replace(value, "0X", "", -1)
		v, _ := strconv.ParseInt(value, 16, 64)
		t.Value = int(v)
	}
}
