package elements

// Limit represents a named limit in nftables.
// It defines the limits on the rate and burst of packets or bytes.
type Limit struct {
	Family string `json:"family"` // The table's family.
	Table  string `json:"table"`  // The table's name.
	Name   string `json:"name"`   // The limit's name.
	Handle int    `json:"handle"` // The limit's handle, used in delete command.
	Rate   int    `json:"rate"`   // The limit's rate value.
	Per    string `json:"per"`    // Time unit for the limit (e.g., "second", "minute").
	Burst  int    `json:"burst"`  // The limit's burst value.
	Unit   string `json:"unit"`   // Unit of rate and burst (e.g., "packets", "bytes").
	Inv    bool   `json:"inv"`    // If true, match if the limit was exceeded.
}
