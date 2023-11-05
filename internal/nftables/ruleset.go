package nftables

type RulesetOptions func(options *listOptions)

type Ruleset interface {
	List(options ...RulesetOptions) []*Rule
	Flush()
}
