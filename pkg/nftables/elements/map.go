package elements

// Map represents a named map in nftables, a special form of sets.
// It is similar to Set but maps a unique key to a value.
type Map struct {
	Family     string   `json:"family"`      // The table's family.
	Table      string   `json:"table"`       // The table's name.
	Name       string   `json:"name"`        // The map's name.
	Handle     int      `json:"handle"`      // The map's handle, used in delete command.
	Type       []string `json:"type"`        // The map's datatype. The set type might be a string, such as "ipv4_addr" or an array consisting of strings (for concatenated types).
	Map        string   `json:"map"`         // Type of values this set maps to.
	Policy     string   `json:"policy"`      // The map's policy.
	Flags      []string `json:"flags"`       // The map's flags.
	Elem       []string `json:"elem"`        // Initial map element(s).
	Timeout    int      `json:"timeout"`     // Element timeout in seconds.
	GcInterval int      `json:"gc-interval"` // Garbage collector interval in seconds.
	Size       int      `json:"size"`        // Maximum number of elements supported.
}
