package nftables

type Ruleset struct {
	Info   MetaInfo `json:"info"`
	Tables []Table  `json:"tables,omitempty"`
	Chains []Chain  `json:"chains,omitempty"`
	Rules  []Rule   `json:"rules,omitempty"`
}
