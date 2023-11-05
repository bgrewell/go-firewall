package nftables

type RuleLocation struct {
	TableHandle int
	ChainHandle int
	RuleHandle  int
}

type Rule struct {
	Family     AddressFamily
	Locations  []RuleLocation
	Expression NotYetFiguredOutStruct
}
