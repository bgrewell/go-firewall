package _old

// Chain represents the default chains
type Chain string

const (
	ChainInput       Chain = "INPUT"
	ChainOutput      Chain = "OUTPUT"
	ChainForward     Chain = "FORWARD"
	ChainPreRouting  Chain = "PREROUTING"
	ChainPostRouting Chain = "POSTROUTING"
)
