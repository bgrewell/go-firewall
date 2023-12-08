package elements

import (
	"github.com/bgrewell/go-firewall/pkg/nftables/consts"
)

type DefaultRule struct {
	// The table’s family.
	Family consts.AddressFamily `json:"family"`
	// The table's name.
	Table Table `json:"table"` // TODO: Needs to be string name of the table
	// The chain's name.
	Chain Chain `json:"chain"` // TODO: Needs to be string name of the chain
	// The rule’s handle. In delete/replace commands, it serves as an identifier of the rule to delete/replace. In add/insert commands, it serves as an identifier of an existing rule to append/prepend the rule to.
	Handle int `json:"handle"`
	// The rule’s position for add/insert commands. It is used as an alternative to handle then.
	Index int `json:"index"`
	// Optional rule comment
	Comment string `json:"comment"`
	// An array of statements this rule consists of. In input, it is used in add/insert/replace commands only.
	Expression interface{} `json:"expr"`
}
