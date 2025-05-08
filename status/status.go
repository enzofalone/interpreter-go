package status

type ReturnCode int

const (
	SUCCESS       ReturnCode = 0
	LEXICAL_ERROR            = 65
	UNKNOWN_ERROR            = 1
)
