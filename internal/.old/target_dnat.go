package _old

import (
	"fmt"
	"strings"
)

const (
	TargetDNatStr string = "--to-destination"
)

type TargetDNat struct {
	DestinationIp        string `json:"destination_ip" yaml:"destination_ip" xml:"destination_ip"`
	DestinationIpRange   string `json:"destination_ip_range" yaml:"destination_ip_range" xml:"destination_ip_range"`
	DestinationPort      string `json:"destination_port" yaml:"destination_port" xml:"destination_port"`
	DestinationPortRange string `json:"destination_port_range" yaml:"destination_port_range" xml:"destination_port_range"`
}

func (t TargetDNat) String() string {
	parts := make([]string, 0)
	parts = append(parts, "DNAT")
	parts = append(parts, TargetDNatStr)
	dstPart := ""
	if t.DestinationIpRange != "" {
		dstPart = t.DestinationIpRange
	} else {
		dstPart = t.DestinationIp
	}
	if t.DestinationPortRange != "" {
		dstPart = fmt.Sprintf("%s:%s", dstPart, t.DestinationPortRange)
	} else if t.DestinationPort != "" {
		dstPart = fmt.Sprintf("%s:%s", dstPart, t.DestinationPort)
	}
	parts = append(parts, dstPart)

	return TargetJump{
		Value: strings.Join(parts, " "),
	}.String()
}

// Returns if the target is valid when applied with the specified rule
func (t TargetDNat) Validate(rule Rule) error {
	// Only valid on the nat table
	if rule.Table != TableNat {
		return fmt.Errorf("target DNAT is only valid on the 'nat' table")
	}
	if rule.Chain == ChainOutput || rule.Chain == ChainPreRouting {
		return fmt.Errorf("target DNAT is only valid on the 'OUTPUT', 'PREROUTING' or custom chains")
	}
	if t.DestinationPort != "" && t.DestinationPortRange != "" && rule.Protocol == "" {
		return fmt.Errorf("target DNAT destination port(s) are only valid when a protocol is specified on the rule")
	}
	if t.DestinationIp == "" && t.DestinationIpRange == "" {
		return fmt.Errorf("target DNAT requires a destination ip address or range")
	}
	return nil
}

func (t *TargetDNat) Parse(option string, value string) {
	if option == "--to-destination" {
		parts := strings.Split(value, ":")
		iprange := strings.Contains(parts[0], "-")
		portrange := strings.Contains(parts[1], "-")
		if iprange {
			t.DestinationIpRange = parts[0]
		} else {
			t.DestinationIp = parts[0]
		}

		if portrange {
			t.DestinationPortRange = parts[1]
		} else {
			t.DestinationPort = parts[1]
		}
	}
}
