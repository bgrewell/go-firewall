package _old

// Counter is a helper type that wraps the packet and byte counters
type Counter struct {
	Packets int `json:"packets,omitempty" yaml:"packets" xml:"packets"`
	Bytes   int `json:"bytes,omitempty" yaml:"bytes" xml:"bytes"`
}
