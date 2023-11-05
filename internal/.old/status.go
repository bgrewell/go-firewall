package _old

type Status struct {
	IPv4Status *IPStatus `json:"ipv4_status,omitempty" yaml:"ipv4_status" xml:"ipv4_status"`
	IPv6Status *IPStatus `json:"ipv6_status,omitempty" yaml:"ipv6_status" xml:"ipv6_status"`
}

type IPStatus struct {
	ModulesLoaded  bool     `json:"modules_loaded,omitempty" yaml:"modules_loaded" xml:"modules_loaded"`
	InputPolicy    string   `json:"input_policy,omitempty" yaml:"input_policy" xml:"input_policy"`
	ForwardPolicy  string   `json:"forward_policy,omitempty" yaml:"forward_policy" xml:"forward_policy"`
	OutputPolicy   string   `json:"output_policy,omitempty" yaml:"output_policy" xml:"output_policy"`
	FilterChains   []string `json:"filter_chains,omitempty" yaml:"filter_chains" xml:"filter_chains"`
	NatChains      []string `json:"nat_chains,omitempty" yaml:"nat_chains" xml:"nat_chains"`
	MangleChains   []string `json:"mangle_chains,omitempty" yaml:"mangle_chains" xml:"mangle_chains"`
	RawChains      []string `json:"raw_chains,omitempty" yaml:"raw_chains" xml:"raw_chains"`
	SecurityChains []string `json:"security_chains,omitempty" yaml:"security_chains" xml:"security_chains"`
}

func GetStatus() (status *Status, err error) {
	ip4, err := GetPolicy(string(TableFilter), string(ChainInput))
	if err != nil {
		return nil, err
	}
	fp4, err := GetPolicy(string(TableFilter), string(ChainForward))
	if err != nil {
		return nil, err
	}
	op4, err := GetPolicy(string(TableFilter), string(ChainOutput))
	if err != nil {
		return nil, err
	}
	fc4, err := EnumerateChains("filter")
	if err != nil {
		return nil, err
	}
	nc4, err := EnumerateChains("nat")
	if err != nil {
		return nil, err
	}
	mc4, err := EnumerateChains("mangle")
	if err != nil {
		return nil, err
	}
	rc4, err := EnumerateChains("raw")
	if err != nil {
		return nil, err
	}
	sc4, err := EnumerateChains("security")
	if err != nil {
		return nil, err
	}

	s4 := &IPStatus{
		ModulesLoaded:  true,
		InputPolicy:    ip4,
		ForwardPolicy:  fp4,
		OutputPolicy:   op4,
		FilterChains:   fc4,
		NatChains:      nc4,
		MangleChains:   mc4,
		RawChains:      rc4,
		SecurityChains: sc4,
	}

	ip6, err := Get6Policy(string(TableFilter), string(ChainInput))
	if err != nil {
		return nil, err
	}
	fp6, err := Get6Policy(string(TableFilter), string(ChainForward))
	if err != nil {
		return nil, err
	}
	op6, err := Get6Policy(string(TableFilter), string(ChainOutput))
	if err != nil {
		return nil, err
	}
	fc6, err := Enumerate6Chains("filter")
	if err != nil {
		return nil, err
	}
	nc6, err := Enumerate6Chains("nat")
	if err != nil {
		return nil, err
	}
	mc6, err := Enumerate6Chains("mangle")
	if err != nil {
		return nil, err
	}
	rc6, err := Enumerate6Chains("raw")
	if err != nil {
		return nil, err
	}
	sc6, err := Enumerate6Chains("security")
	if err != nil {
		return nil, err
	}

	s6 := &IPStatus{
		ModulesLoaded:  true,
		InputPolicy:    ip6,
		ForwardPolicy:  fp6,
		OutputPolicy:   op6,
		FilterChains:   fc6,
		NatChains:      nc6,
		MangleChains:   mc6,
		RawChains:      rc6,
		SecurityChains: sc6,
	}

	s := &Status{
		IPv4Status: s4,
		IPv6Status: s6,
	}

	return s, nil
}
