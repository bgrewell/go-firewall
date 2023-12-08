package statements

// Log represents a logging configuration.
type Log struct {
	Prefix         string   `json:"prefix,omitempty"`          // Prefix for log entries.
	Group          int      `json:"group,omitempty"`           // Log group.
	Snaplen        int      `json:"snaplen,omitempty"`         // Snaplen for logging.
	QueueThreshold int      `json:"queue-threshold,omitempty"` // Queue threshold.
	Level          string   `json:"level,omitempty"`           // Log level.
	Flags          []string `json:"flags,omitempty"`           // Log flags.
}
