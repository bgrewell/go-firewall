package _old

// Marker interface is used to allow custom markers to be used. Markers are simply a specially formatted
// comment on the rule that follows the <name>:<value> format
type Marker interface {
	Name() string
	SetName(name string) error
	Value() string
	SetValue(value string) error
	Parse(comment string) error
	String() string
}
