package _old

import (
	"fmt"
)

const (
	TargetBalanceStr string = "--to-destination"
)

type TargetBalance struct {
	StartingIpAddress string `json:"starting_ip_address" yaml:"starting_ip_address" xml:"starting_ip_address"`
	EndingIpAddress   string `json:"ending_ip_address" yaml:"ending_ip_address" xml:"ending_ip_address"`
}

//func (t *TargetBalance) MarshalJSON() (b []byte, e error) {
//	type TargetBalanceHelper struct {
//		Type string `json:"type"`
//		Value *TargetBalance `json:"value"`
//	}
//
//	th := TargetBalanceHelper{
//		Type: "balance",
//		Value: t,
//	}
//	return json.Marshal(th)
//}

func (t TargetBalance) String() string {
	return TargetJump{
		Value: fmt.Sprintf("BALANCE %s %s-%s", TargetBalanceStr, t.StartingIpAddress, t.EndingIpAddress),
	}.String()
}

func (t TargetBalance) Validate(rule Rule) error {
	return nil
}

func (t *TargetBalance) Parse(option string, value string) {

}
