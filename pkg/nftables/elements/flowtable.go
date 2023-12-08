package elements

// Flowtable represents a named flowtable in nftables.
// It includes various attributes to describe the flowtable's properties.
type Flowtable struct {
	Family string   `json:"family"` // The table's family.
	Table  string   `json:"table"`  // The table's name.
	Name   string   `json:"name"`   // The flow table's name.
	Handle int      `json:"handle"` // The flow table's handle, used in delete command.
	Hook   string   `json:"hook"`   // The flow table's hook.
	Prio   int      `json:"prio"`   // The flow table's priority.
	Dev    []string `json:"dev"`    // The flow table's interface(s).
}
