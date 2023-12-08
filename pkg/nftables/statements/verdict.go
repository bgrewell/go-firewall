package statements

// Verdict represents the action to be taken for a packet.
type Verdict struct {
	Accept   *struct{} `json:"accept,omitempty"`
	Drop     *struct{} `json:"drop,omitempty"`
	Continue *struct{} `json:"continue,omitempty"`
	Return   *struct{} `json:"return,omitempty"`
	Jump     *Target   `json:"jump,omitempty"`
	Goto     *Target   `json:"goto,omitempty"`
}

// Target represents the target chain for jump and goto statements.
type Target struct {
	Target string `json:"target"`
}
