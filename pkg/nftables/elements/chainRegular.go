package elements

import (
	"github.com/bgrewell/go-firewall/pkg/nftables/consts"
)

type RegularChain struct {
	Family consts.AddressFamily `json:"family"`
	Table  Table                `json:"table"`
	Name   string               `json:"name"`
	Handle int                  `json:"handle"`
}
