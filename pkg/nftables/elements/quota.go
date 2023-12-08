package elements

// Quota represents a named quota in nftables.
// It includes attributes for identifying the quota and its current usage.
type Quota struct {
	Family string `json:"family"` // The table's family.
	Table  string `json:"table"`  // The table's name.
	Name   string `json:"name"`   // The quota's name.
	Handle int    `json:"handle"` // The quota's handle, used in delete command.
	Bytes  int    `json:"bytes"`  // Quota threshold in bytes.
	Used   int    `json:"used"`   // Quota used so far in bytes.
	Inv    bool   `json:"inv"`    // If true, match if the quota has been exceeded.
}
