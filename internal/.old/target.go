package _old

const (
	TargetAccept string = "ACCEPT"
	TargetDrop   string = "DROP"
	TargetQueue  string = "QUEUE"
	TargetReturn string = "RETURN"
)

// Target is an interface for the target extensions
type Target interface {
	String() string
	Validate(rule Rule) error
	Parse(option string, value string)
}
