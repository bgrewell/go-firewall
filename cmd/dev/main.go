package main

import (
	"fmt"
	"github.com/bgrewell/go-firewall/pkg/nftables/helpers"
	"github.com/bgrewell/go-firewall/pkg/nftables/statements"
	"log"
)

func main() {

	p := 12345
	b := 67890
	counter := statements.Counter{
		Packets: &p,
		Bytes:   &b,
		Name:    "counter_xyz",
	}

	// convert to JSON that nftables will understand
	nftJSON, err := helpers.ToPrettyJSON(counter)
	if err != nil {
		log.Fatalf("Error marshalling: %v\n", err)
	}
	fmt.Println(nftJSON)

	// convert back
	var ctype statements.Counter
	err = helpers.FromJSON([]byte(nftJSON), &ctype)
	if err != nil {
		log.Fatalf("Error unmarshalling: %v\n", err)
	}
	fmt.Printf("%#v\n", ctype)
}
