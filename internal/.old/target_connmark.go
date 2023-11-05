package _old

import (
	"fmt"
)

//TODO: Need to fully implement by supporting masks for the mark

type ConnMarkType int

const (
	ConnMarkTypeSet = iota
	ConnMarkTypeSave
	ConnMarkTypeRestore
)

const (
	TargetConnMarkSet     string = "--set-mark"
	TargetConnMarkSave    string = "--save-mark"
	TargetConnMarkRestore string = "--restore-mark"
)

type TargetConnMark struct {
	MarkType ConnMarkType `json:"mark_type" yaml:"mark_type" xml:"mark_type"`
	Value    int          `json:"value" yaml:"value" xml:"value"`
}

//func (t *TargetConnMark) MarshalJSON() (b []byte, e error) {
//	type TargetConnMarkHelper struct {
//		Type string `json:"type"`
//		Value *TargetConnMark `json:"value"`
//	}
//
//	th := TargetConnMarkHelper{
//		Type: "connmark",
//		Value: t,
//	}
//	return json.Marshal(th)
//}

func (t TargetConnMark) String() string {
	target := ""
	switch t.MarkType {
	case ConnMarkTypeSet:
		target = TargetConnMarkSet
	case ConnMarkTypeSave:
		target = TargetConnMarkSave
	case ConnMarkTypeRestore:
		target = TargetConnMarkRestore
	}
	return TargetJump{
		Value: fmt.Sprintf("CONNMARK %s %d", target, t.Value),
	}.String()
}

// Returns if the target is valid when applied with the specified rule
func (t TargetConnMark) Validate(rule Rule) error {
	// Only valid on the mangle table
	if rule.Table != TableMangle {
		return fmt.Errorf("target CONNMARK is only valid on the 'mangle' table")
	}
	return nil
}

func (t *TargetConnMark) Parse(option string, value string) {

}
