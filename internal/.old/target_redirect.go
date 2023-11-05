package _old

import (
	"fmt"
	"strings"
)

const (
	TargetRedirectStr string = "--to-ports"
)

type TargetRedirect struct {
	DestinationPort      string `json:"destination_port" yaml:"destination_port" xml:"destination_port"`
	DestinationPortRange string `json:"destination_port_range" yaml:"destination_port_range" xml:"destination_port_range"`
}

func (t TargetRedirect) String() string {
	parts := make([]string, 0)
	parts = append(parts, "REDIRECT")
	parts = append(parts, TargetRedirectStr)
	if t.DestinationPortRange != "" {
		parts = append(parts, t.DestinationPortRange)
	} else if t.DestinationPort != "" {
		parts = append(parts, t.DestinationPort)
	}

	return TargetJump{
		Value: strings.Join(parts, " "),
	}.String()
}

// Returns if the target is valid when applied with the specified rule
func (t TargetRedirect) Validate(rule Rule) error {
	// Only valid on the nat table
	if rule.Table != TableNat {
		return fmt.Errorf("target REDIRECT is only valid on the 'nat' table")
	}
	if rule.Chain == ChainOutput || rule.Chain == ChainPreRouting {
		return fmt.Errorf("target REDIRECT is only valid on the 'OUTPUT', 'PREROUTING' or custom chains")
	}
	if t.DestinationPort != "" && t.DestinationPortRange != "" && rule.Protocol == "" {
		return fmt.Errorf("target REDIRECT destination port(s) are only valid when a protocol is specified on the rule")
	}
	return nil
}

func (t *TargetRedirect) Parse(option string, value string) {
	if option == "--to-ports" {
		portrange := strings.Contains(value, "-")
		if portrange {
			t.DestinationPortRange = value
		} else {
			t.DestinationPort = value
		}
	}
}
