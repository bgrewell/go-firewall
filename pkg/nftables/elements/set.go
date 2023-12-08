package elements

// Set represents a named set in nftables.
// It includes various attributes like family, table, name, and more.
type Set struct {
	Family     string   `json:"family"`      // The table's family.
	Table      string   `json:"table"`       // The table's name.
	Name       string   `json:"name"`        // The set's name.
	Handle     int      `json:"handle"`      // The set's handle, used in delete command.
	Type       []string `json:"type"`        // The set's datatype. The set type might be a string, such as "ipv4_addr" or an array consisting of strings (for concatenated types).
	Policy     string   `json:"policy"`      // The set's policy.
	Flags      []string `json:"flags"`       // The set's flags.
	Elem       []string `json:"elem"`        // Initial set element(s).
	Timeout    int      `json:"timeout"`     // Element timeout in seconds.
	GcInterval int      `json:"gc-interval"` // Garbage collector interval in seconds.
	Size       int      `json:"size"`        // Maximum number of elements supported.
}
