package nftables

import (
	"encoding/json"
	"fmt"
)

type NftablesData struct {
	Nftables []json.RawMessage `json:"nftables"`
}

func UnmarshalNftables(nftablesData string) (*Ruleset, error) {
	var data NftablesData
	if err := json.Unmarshal([]byte(nftablesData), &data); err != nil {
		return nil, err
	}

	ruleset := &Ruleset{}

	for _, item := range data.Nftables {
		var obj map[string]json.RawMessage
		if err := json.Unmarshal(item, &obj); err != nil {
			return nil, err
		}

		if val, ok := obj["metainfo"]; ok {
			var info MetaInfo
			if err := json.Unmarshal(val, &info); err != nil {
				return nil, err
			}
			ruleset.Info = info
		} else if val, ok := obj["table"]; ok {
			table, err := UnmarshalTable(val)
			if err != nil {
				return nil, err
			}
			ruleset.Tables = append(ruleset.Tables, table)
		} else if val, ok := obj["chain"]; ok {
			chain, err := UnmarshalChain(val)
			if err != nil {
				return nil, err
			}
			ruleset.Chains = append(ruleset.Chains, chain)
		} else if val, ok := obj["rule"]; ok {
			rule, err := UnmarshalRule(val)
			if err != nil {
				return nil, err
			}
			ruleset.Rules = append(ruleset.Rules, rule)
		} else {
			return nil, fmt.Errorf("unknown element type in nftables data: %v\n", string(item))
		}
	}

	return ruleset, nil
}
