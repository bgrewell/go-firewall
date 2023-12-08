package elements

import "encoding/json"

type Chain interface {
	Rename(newname string)
	//// AddRule adds the specified rule at the end of the chain
	//AddRule(rule *Rule) error
	//// InsertRule adds the specified rule at the beginning of the chain or before the specified rule
	//InsertRule(rule *Rule) error
	//// ReplaceRule replaces the specified rule
	//ReplaceRule(rule *Rule) error
	//// DeleteRule removes the specified rule or returns an error if the rule doesn't exist
	//DeleteRule(rule *Rule) error
	//// DestroyRule removes the specified rule but does NOT fail if the rule doesn't exist
	//DestroyRule(rule *Rule) error
}

func UnmarshalChain(chainJSON []byte) (chain Chain, err error) {
	// Check to see if the chain has a hook, if it does then it is a base chain, otherwise it's a regular chain
	var baseChain BaseChain
	if err := json.Unmarshal(chainJSON, &baseChain); err != nil {
		return nil, err
	}

	// Check if the 'hook' key is present and not empty
	if baseChain.Hook != "" {
		// Return BaseChain if 'hook' is present
		return &baseChain, nil
	} else {
		// Return RegularChain part of BaseChain if 'hook' is not present
		return &baseChain.RegularChain, nil
	}
}
