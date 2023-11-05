package _old

import (
	"fmt"
	"strings"
)

const (
	TargetSNatStr string = "--to-source"
)

type TargetSNat struct {
	SourceIp        string `json:"source_ip,omitempty" yaml:"source_ip" xml:"source_ip"`
	SourceIpRange   string `json:"source_ip_range,omitempty" yaml:"source_ip_range" xml:"source_ip_range"`
	SourcePort      string `json:"source_port,omitempty" yaml:"source_port" xml:"source_port"`
	SourcePortRange string `json:"source_port_range,omitempty" yaml:"source_port_range" xml:"source_port_range"`
}

//func (t *TargetSNat) MarshalJSON() (b []byte, e error) {
//	type TargetSNatHelper struct {
//		Type string `json:"type"`
//		Value *TargetSNat `json:"value"`
//	}
//
//	th := TargetSNatHelper{
//		Type: "snat",
//		Value: t,
//	}
//	return json.Marshal(th)
//}

func (t TargetSNat) String() string {
	parts := make([]string, 0)
	parts = append(parts, "SNAT")
	parts = append(parts, TargetSNatStr)
	if t.SourceIpRange != "" {
		parts = append(parts, t.SourceIpRange)
	} else {
		parts = append(parts, t.SourceIp)
	}
	if t.SourcePortRange != "" {
		parts = append(parts, fmt.Sprintf(":%s", t.SourcePortRange))
	} else if t.SourcePort != "" {
		parts = append(parts, fmt.Sprintf(":%s", t.SourcePort))
	}

	return TargetJump{
		Value: strings.Join(parts, " "),
	}.String()
}

// Returns if the target is valid when applied with the specified rule
func (t TargetSNat) Validate(rule Rule) error {
	// Only valid on the mangle table
	if rule.Table != TableNat {
		return fmt.Errorf("target SNAT is only valid on the 'nat' table")
	}
	if rule.Chain == ChainPreRouting || rule.Chain == ChainInput || rule.Chain == ChainOutput {
		return fmt.Errorf("target SNAT is only valid on the 'POSTROUTING' or custom chains") //TODO: need more smarts as this is valid if it's on a chain that is jumped to from POSTROUTING I believe
	}
	if t.SourcePort != "" && t.SourcePortRange != "" && rule.Protocol == "" {
		return fmt.Errorf("target DNAT destination port(s) are only valid when a protocol is specified on the rule")
	}
	if t.SourceIp == "" && t.SourceIpRange == "" {
		return fmt.Errorf("target DNAT requires a destination ip address or range")
	}
	return nil
}

func (t *TargetSNat) Parse(option string, value string) {

}
