package nftables

import "github.com/bgrewell/go-firewall/pkg/nftables/consts"

type DefaultRule struct {
	Family     consts.AddressFamily `json:"family,omitempty"`
	Table      Table                `json:"-"`
	Chain      Chain                `json:"-"`
	Handle     int                  `json:"handle,omitempty"`
	Expression interface{}          `json:"expression,omitempty"`
}
