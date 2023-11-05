package _old

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	TargetDSCPStr      string = "--set-dscp"
	TargetDSCPClassStr string = "--set-dscp-class"
)

type TargetDSCP struct {
	Value int `json:"value" yaml:"value" xml:"value"`
}

//func (t *TargetDSCP) MarshalJSON() (b []byte, e error) {
//	type TargetDSCPHelper struct {
//		Type string `json:"type"`
//		Value *TargetDSCP `json:"value"`
//	}
//
//	th := TargetDSCPHelper{
//		Type: "goto",
//		Value: t,
//	}
//	return json.Marshal(th)
//}

func (t TargetDSCP) String() string {
	return TargetJump{
		Value: fmt.Sprintf("DSCP %s %d", TargetDSCPStr, t.Value),
	}.String()
}

// Returns if the target is valid when applied with the specified rule
func (t TargetDSCP) Validate(rule Rule) error {
	// Only valid on the mangle table
	if rule.Table != TableMangle {
		return fmt.Errorf("target DSCP is only valid on the 'mangle' table")
	}
	return nil
}

func (t *TargetDSCP) Parse(option string, value string) {
	if option == "--set-dscp" {
		value = strings.Replace(value, "0x", "", -1)
		value = strings.Replace(value, "0X", "", -1)
		v, _ := strconv.ParseInt(value, 16, 64)
		t.Value = int(v)
	}
}

type TargetDSCPClass struct {
	Class string `json:"class" yaml:"class" xml:"class"`
}

func (t TargetDSCPClass) String() string {
	return TargetJump{
		Value: fmt.Sprintf("DSCP %s %d", TargetDSCPClassStr, t.Class),
	}.String()
}

// Returns if the target is valid when applied with the specified rule
func (t TargetDSCPClass) Validate(rule Rule) error {
	// Only valid on the mangle table
	if rule.Table != TableMangle {
		return fmt.Errorf("target DSCP is only valid on the 'mangle' table")
	}
	return nil
}

func (t *TargetDSCPClass) Parse(option string, value string) {

}
