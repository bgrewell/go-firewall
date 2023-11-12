package consts

// IPFamilyHooks represent the packet processing stages in the network stack
type IPFamilyHooks string

const (
	// HookIPPrerouting handles all packets entering the system are processed by the prerouting hook. It is invoked before the routing process and is used for early filtering or changing packet attributes that affect routing
	HookIPPrerouting IPFamilyHooks = "prerouting"
	// HookIPInput handles all packets delivered to the local system are processed by the input hook
	HookIPInput IPFamilyHooks = "input"
	// HookIPForward handles all packets forwarded to a different host are processed by the forward hook
	HookIPForward IPFamilyHooks = "forward"
	// HookIPOutput handles all packets sent by local processes are processed by the output hook
	HookIPOutput IPFamilyHooks = "output"
	// HookIPPostRouting handles all packets leaving the system are processed by the postrouting hook
	HookIPPostRouting IPFamilyHooks = "postrouting"
	// HookIPIngress handles all packets entering the system are processed by this hook. It is invoked before layer 3 protocol handlers, hence before the prerouting hook, and it can be used for filtering and policing. Ingress is only available for Inet family
	HookIPIngress IPFamilyHooks = "ingress"
)
