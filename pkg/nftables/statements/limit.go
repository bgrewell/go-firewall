package statements

// Limit represents a rate limiting configuration.
type Limit struct {
	Rate      int    `json:"rate,omitempty"`       // Rate value.
	RateUnit  string `json:"rate_unit,omitempty"`  // Unit of rate.
	Per       string `json:"per,omitempty"`        // Time unit for rate.
	Burst     int    `json:"burst,omitempty"`      // Burst value.
	BurstUnit string `json:"burst_unit,omitempty"` // Unit of burst.
	Inv       bool   `json:"inv,omitempty"`        // Match if limit exceeded.
	Name      string `json:"limit,omitempty"`      // Reference to a named limit.
}
