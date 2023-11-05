package _old

import (
	"fmt"
	"strings"
)

const (
	TargetMasqueradeStr string = "--to-source"
)

type TargetMasquerade struct {
	DestinationPort      string `json:"destination_port" yaml:"destination_port" xml:"destination_port"`
	DestinationPortRange string `json:"destination_port_range" yaml:"destination_port_range" xml:"destination_port_range"`
}

//func (t *TargetMasquerade) MarshalJSON() (b []byte, e error) {
//	type TargetMasqueradeHelper struct {
//		Type string `json:"type"`
//		Value *TargetMasquerade `json:"value"`
//	}
//
//	th := TargetMasqueradeHelper{
//		Type: "masquerade",
//		Value: t,
//	}
//	return json.Marshal(th)
//}

func (t TargetMasquerade) String() string {
	parts := make([]string, 0)
	parts = append(parts, "MASQUERADE")
	if t.DestinationPort != "" {
		parts = append(parts, TargetMasqueradeStr, t.DestinationPort)
	} else if t.DestinationPortRange != "" {
		parts = append(parts, TargetMasqueradeStr, t.DestinationPortRange)
	}

	return TargetJump{
		Value: strings.Join(parts, " "),
	}.String()
}

// Returns if the target is valid when applied with the specified rule
func (t TargetMasquerade) Validate(rule Rule) error {
	// Only valid on the mangle table
	if rule.Table != TableNat {
		return fmt.Errorf("target MASQUERADE is only valid on the 'nat' table")
	}
	if rule.Chain != ChainPostRouting {
		return fmt.Errorf("target MASQUERADE is only valid on the 'POSTROUTING' chain") //TODO: need more smarts as this is valid if it's on a chain that is jumped to from POSTROUTING I believe
	}

	return nil
}

func (t *TargetMasquerade) Parse(option string, value string) {
	if option == "--to-ports" {
		if strings.Contains(value, "-") {
			t.DestinationPortRange = value
		} else {
			t.DestinationPort = value
		}
	}
}
