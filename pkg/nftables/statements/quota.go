package statements

// Quota represents a quota limit.
type Quota struct {
	Val      int    `json:"val,omitempty"`       // Quota value.
	ValUnit  string `json:"val_unit,omitempty"`  // Unit of quota value.
	Used     *int   `json:"used,omitempty"`      // Quota used so far.
	UsedUnit string `json:"used_unit,omitempty"` // Unit of used quota.
	Inv      bool   `json:"inv,omitempty"`       // Match if quota was exceeded.
	Name     string `json:"quota,omitempty"`     // Reference to a named quota.
}
