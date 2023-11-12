package consts

type NetdevFamilyHooks string

const (
	// HookNetdevIngress handles packets entering the system are processed by this hook. It is invoked after the network taps (ie. tcpdump), right after tc ingress and before layer 3 protocol handlers, it can be used for early filtering and policing
	HookNetdevIngress NetdevFamilyHooks = "ingress"
	// HookNetdevEgress handles packets leaving the system are processed by this hook. It is invoked after layer 3 protocol handlers and before tc egress. It can be used for late filtering and policing
	HookNetdevEgress NetdevFamilyHooks = "egress"
)
