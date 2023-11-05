package main

import (
	"bufio"
	"fmt"
	"github.com/bgrewell/go-firewall/internal"
	"github.com/bgrewell/go-firewall/internal/.old"
	"os"
)

func main() {

	mo := _old.NewMatchGeneric("physdev", "physdev-out", "eth-up", false)
	mi := _old.NewMatchGeneric("physdev", "physdev-in", "eth-dn", false)
	r := &internal.iptables{
		Id:           "rule-2-ul-0",
		Name:         "super-test-ul-0",
		Table:        _old.TableFilter,
		Chain:        _old.ChainForward,
		IpVersion:    _old.IPv4,
		TargetAction: internal.iptables.ActionJump,
		Target: &_old.TargetMark{
			Value: 0x123,
		},
	}
	r.AddMatch(mo)
	r.SetApp("wanemd-super-test")
	r.Debug = true

	err := r.Append()
	if err != nil {
		fmt.Printf("err: %s\n", err)
	} else {
		fmt.Println("added ul-rule")
	}

	r = &internal.iptables{
		Id:           "rule-2-dl-0",
		Name:         "super-test-dl-0",
		Table:        _old.TableFilter,
		Chain:        _old.ChainForward,
		IpVersion:    _old.IPv4,
		TargetAction: internal.iptables.ActionJump,
		Target: &_old.TargetMark{
			Value: 0x123,
		},
	}
	r.AddMatch(mi)
	r.SetApp("wanemd-super-test")
	r.Debug = true

	err = r.Append()
	if err != nil {
		fmt.Printf("err: %s\n", err)
	} else {
		fmt.Println("added dl-rule")
	}

	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	fmt.Println("attempting to delete rules")

	err = _old.DeleteAllMatchingApp("wanemd-super-test")
	if err != nil {
		fmt.Printf("err: %s\n", err)
	} else {
		fmt.Println("deleted rules")
	}

}
