package elements

import "github.com/bgrewell/go-firewall/pkg/nftables"

type Ruleset struct {
	Info   nftables.MetaInfo `json:"info"`
	Tables []Table           `json:"tables,omitempty"`
	Chains []Chain           `json:"chains,omitempty"`
	Rules  []Rule            `json:"rules,omitempty"`
}
