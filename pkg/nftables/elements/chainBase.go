package elements

import "github.com/bgrewell/go-firewall/pkg/nftables/consts"

type BaseChain struct {
	RegularChain
	Type   consts.ChainType `json:"type"`
	Hook   string           `json:"hook"`
	Prio   int              `json:"prio"`
	Dev    string           `json:"dev"` // only used for netdev family
	Policy string           `json:"policy"`
}
