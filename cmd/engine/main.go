package main

import (
	"fmt"
	"github.com/bgrewell/go-firewall/internal/netlink"
	"github.com/bgrewell/go-firewall/pkg/engines"
	"time"
)

func main() {
	notifyChan := make(chan struct{})
	monitor, err := netlink.NewNetlinkMonitor()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	engine, err := engines.NewEngine(engines.WithMonitor(monitor), engines.WithChangeNotify(notifyChan))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	go func() {
		for {
			select {
			case <-notifyChan:
				fmt.Println("[!] nftables rules change detected")
			}
		}
	}()

	err = engine.Start()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	time.Sleep(100 * time.Second)

	err = engine.Stop()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
