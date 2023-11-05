package _old

import (
	"fmt"
	"strings"
)

const (
	TargetChecksumStr string = "--checksum-fill"
)

type TargetChecksum struct {
	ChecksumFill string `json:"checksum_fill" yaml:"checksum_fill" xml:"checksum_fill"`
}

func (t TargetChecksum) String() string {
	parts := make([]string, 0)
	parts = append(parts, "CHECKSUM")
	parts = append(parts, TargetChecksumStr)

	return TargetJump{
		Value: strings.Join(parts, " "),
	}.String()
}

func (t TargetChecksum) Validate(rule Rule) error {
	if rule.Table != TableMangle {
		return fmt.Errorf("target CHECKSUM is only valid on the 'mangle' table")
	}
	return nil
}

func (t *TargetChecksum) Parse(option string, value string) {
}
