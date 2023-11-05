package nftables

type Chain interface {
	// AddRule adds the specified rule at the end of the chain
	AddRule(rule *Rule) error
	// InsertRule adds the specified rule at the beginning of the chain or before the specified rule
	InsertRule(rule *Rule) error
	// ReplaceRule replaces the specified rule
	ReplaceRule(rule *Rule) error
	// DeleteRule removes the specified rule or returns an error if the rule doesn't exist
	DeleteRule(rule *Rule) error
	// DestroyRule removes the specified rule but does NOT fail if the rule doesn't exist
	DestroyRule(rule *Rule) error
}
