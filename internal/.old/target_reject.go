package _old

import (
	"strings"
)

const (
	TargetRejectStr string = "--reject-with"
)

type TargetReject struct {
	RejectType string `json:"reject_type" yaml:"reject_type" xml:"reject_type"`
}

func (t TargetReject) String() string {
	parts := make([]string, 0)
	parts = append(parts, "REJECT")
	if t.RejectType != "" {
		parts = append(parts, TargetRejectStr, t.RejectType)
	} else {
		parts = append(parts, TargetRejectStr, "--icmp-port-unreachable")
	}

	return TargetJump{
		Value: strings.Join(parts, " "),
	}.String()
}

func (t TargetReject) Validate(rule Rule) error {
	return nil
}

func (t *TargetReject) Parse(option string, value string) {
	if option == "--reject-with" {
		t.RejectType = value
	}
}
