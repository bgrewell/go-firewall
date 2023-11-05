package nftables

// AddressFamily represents the address family types supported by nftables
type AddressFamily string

// Constants relating to the table families
const (
	// AddressFamilyIP The IPv4 address families handle IPv4 types of packets, equivalent to iptables
	AddressFamilyIP AddressFamily = "ip"
	// AddressFamilyIP6 The IPv6 address families handle IPv6 types of packets, equivalent to ip6tables
	AddressFamilyIP6 AddressFamily = "ip6"
	// AddressFamilyINET The Inet address families handle both IPv4 and IPv6 types of packets
	AddressFamilyINET AddressFamily = "inet"
	// AddressFamilyARP The ARP address family handles ARP packets received and sent by the system, equivalent to arptables
	AddressFamilyARP AddressFamily = "arp"
	// AddressFamilyBridge The bridge address family handles Ethernet packets traversing bridge devices, equivalent to ebtables
	AddressFamilyBridge AddressFamily = "bridge"
	// AddressFamilyNetdev The Netdev address family handles packets from the device ingress and egress path. This family allows you to filter packets of any ethertype such as ARP, VLAN 802.1q, VLAN 802.1ad (Q-in-Q) as well as IPv4 and IPv6 packets
	AddressFamilyNetdev AddressFamily = "netdev"
)

// ChainType represents the types of chains supported by nftables
type ChainType string

// Constants relating to the chain types
const (
	// ChainTypeBase is an entry point for packets from the network stack, where a hook is specified
	ChainTypeBase ChainType = "base"
	// ChainTypeRegular is a jump target used for rule organization
	ChainTypeRegular ChainType = "regular"
)

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

// ARPFamilyHooks represents the packet processing stages in the network stack
type ARPFamilyHooks string

const (
	// HookARPInput handles packets delivered to the local system are processed by the input hook
	HookARPInput ARPFamilyHooks = "input"
	// HookARPOutput handles packets sent by the local system are processed by the output hook
	HookARPOutput ARPFamilyHooks = "output"
)

type NetdevFamilyHooks string

const (
	// HookNetdevIngress handles packets entering the system are processed by this hook. It is invoked after the network taps (ie. tcpdump), right after tc ingress and before layer 3 protocol handlers, it can be used for early filtering and policing
	HookNetdevIngress NetdevFamilyHooks = "ingress"
	// HookNetdevEgress handles packets leaving the system are processed by this hook. It is invoked after layer 3 protocol handlers and before tc egress. It can be used for late filtering and policing
	HookNetdevEgress NetdevFamilyHooks = "egress"
)
