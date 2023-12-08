package elements

import "github.com/bgrewell/go-firewall/pkg/nftables/consts"

// Counter represents a named counter in nftables.
// It includes attributes for identifying the counter and its metrics.
type Counter struct {
	Family  consts.AddressFamily `json:"family"`  // The table's family.
	Table   string               `json:"table"`   // The table's name.
	Name    string               `json:"name"`    // The counter's name.
	Handle  int                  `json:"handle"`  // The counter's handle, used in delete command.
	Packets int                  `json:"packets"` // Packet counter value.
	Bytes   int                  `json:"bytes"`   // Byte counter value.
}
