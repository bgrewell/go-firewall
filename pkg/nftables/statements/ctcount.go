package statements

// CTCount limits the number of connections using conntrack.
type CTCount struct {
	Val int  `json:"val"` // Connection count threshold.
	Inv bool `json:"inv"` // Match if val exceeded.
}
