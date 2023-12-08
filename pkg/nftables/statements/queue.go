package statements

// Queue represents a queuing configuration.
type Queue struct {
	Num   string   `json:"num"`   // Queue number.
	Flags []string `json:"flags"` // Queue flags.
}
