package elements

import "encoding/json"

type Rule interface{}

func UnmarshalRule(ruleJSON []byte) (rule Rule, err error) {
	var defaultRule DefaultRule
	if err := json.Unmarshal(ruleJSON, &defaultRule); err != nil {
		return nil, err
	}

	return &defaultRule, nil
}
