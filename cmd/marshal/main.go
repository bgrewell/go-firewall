package main

import (
	"encoding/json"
	"fmt"
	"github.com/bgrewell/go-firewall/internal"
	"github.com/bgrewell/go-firewall/internal/.old"
)

func main() {

	r := &internal.iptables{
		Id:                 "abc",
		Name:               "postman-test",
		Table:              "filter",
		Chain:              "forward",
		IpVersion:          "ipv4",
		Protocol:           "tcp",
		Source:             "12.34.56.78",
		Destination:        "1.2.3.4",
		DestinationNegated: true,
		TargetAction:       "jump",
		Target:             &_old.TargetJump{Value: "ACCEPT"},
	}

	b, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(b))

	var rd internal.iptables
	err = json.Unmarshal(b, &rd)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rd.String())

}
