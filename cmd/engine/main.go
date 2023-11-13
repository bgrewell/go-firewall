package main

import (
	"fmt"
	"github.com/bgrewell/go-firewall/pkg/engines"
	"time"
)

func main() {
	engine, err := engines.NewEngine()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

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
