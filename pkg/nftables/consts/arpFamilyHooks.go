package consts

// ARPFamilyHooks represents the packet processing stages in the network stack
type ARPFamilyHooks string

const (
	// HookARPInput handles packets delivered to the local system are processed by the input hook
	HookARPInput ARPFamilyHooks = "input"
	// HookARPOutput handles packets sent by the local system are processed by the output hook
	HookARPOutput ARPFamilyHooks = "output"
)
