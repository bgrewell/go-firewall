package nftables

import "github.com/bgrewell/go-firewall/pkg/nftables/consts"

type BaseChain struct {
	RegularChain
	Type   consts.ChainType `json:"type"`
	Hook   string           `json:"hook"`
	Prio   int              `json:"prio"`
	Policy string           `json:"policy"`
}
