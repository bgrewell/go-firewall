package elements

import "github.com/bgrewell/go-firewall/pkg/nftables/consts"

type DefaultTable struct {
	Family consts.AddressFamily `json:"family"`
	Name   string               `json:"name"`
	Handle int                  `json:"handle"`
}
