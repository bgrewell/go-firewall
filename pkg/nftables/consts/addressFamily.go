package consts

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
