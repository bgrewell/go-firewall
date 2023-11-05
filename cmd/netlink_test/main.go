package main

import (
	"fmt"
	"github.com/bgrewell/go-firewall/internal/netlink"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// Create a new netlink monitor - error can be ignored since the current implementation can't throw an error
	nlm, _ := netlink.NewNetlinkMonitor()
	events, errors, err := nlm.Monitor()
	if err != nil {
		fmt.Printf("Error while setting up Netlink monitor: %v\n", err)
		os.Exit(1)
	}

	// Watch the events channel
	go func() {
		for {
			event := <-events
			header, err := netlink.ParseNetlinkMessageHeader(event)
			if err != nil {
				fmt.Printf("Failed to parse message header: %v\n", err)
			}
			fmt.Printf("Got event - len: 0x%x, type: 0x%x, flags: 0x%x, seq: 0x%x, pid: %d [%d bytes]\n",
				header.Len, header.Type, header.Flags, header.Seq, header.Pid, len(event))
			if len(event) > 128 {
				payload := event[128:]
				for _, b := range payload {
					fmt.Printf("%02x ", b)
				}
				fmt.Println("")
			}
		}
	}()

	// Watch the errors channel
	go func() {
		for {
			err := <-errors
			fmt.Printf("Error occurred: %v\n", err)
		}
	}()

	// Handle ctrl-c
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)
	fmt.Println("\nPress Ctrl-C to exit...")
	<-sigChan
	err = nlm.Stop()
	if err != nil {
		fmt.Printf("Error while stopping Netlink monitor: %v\n", err)
	}
}
