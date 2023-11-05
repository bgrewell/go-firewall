package main

import (
	"fmt"
	"github.com/bgrewell/go-firewall/internal"
	"github.com/bgrewell/go-firewall/internal/.old"
)

func main() {
	// Test to make sure that the format that wanemd needs is fully supported
	// c.IpTablesFmtString = "/sbin/iptables -A rule-%d -m physdev --physdev-out %s -o %s %s -j MARK --set-mark %d" // [rule.Id, device, bridge, filter, handle]
	mo := _old.NewMatchGeneric("physdev", "physdev-out", "eth-up", false)

	r := &internal.iptables{
		Id:           "rule-123",
		Name:         "serverXYZ",
		Table:        _old.TableFilter,
		Chain:        _old.ChainForward,
		IpVersion:    _old.IPv4,
		TargetAction: internal.iptables.ActionJump,
		Target: &_old.TargetMark{
			Value: 0x123,
		},
	}
	r.AddMatch(mo)
	r.SetApp("wanemd")
	r.Debug = true

	err := r.Append()
	if err != nil {
		fmt.Printf("err: %s\n", err)
	} else {
		fmt.Println("Done!")
	}
}
