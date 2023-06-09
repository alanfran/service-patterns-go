package status

// Enum is a pseudo-enum type, that should not be constructed from bare integers.
type Enum int

const (
	// private zero-value to denote an uninitialized variable
	unspecified Enum = iota
	// Reachable - device is connected and responsive
	Reachable
	// Unreachable - device is not responding
	Unreachable
	// Maintenance - down for planned maintenance
	Maintenance
)

var index = map[Enum]string{
	Reachable:   "REACHABLE",
	Unreachable: "UNREACHABLE",
	Maintenance: "MAINTENANCE",
}

// Parse parses a string and returns either the enum it represents,
// or an error.
func Parse(input string) (Enum, error) {
	for enum, str := range index {
		if str == input {
			return enum, nil
		}
	}

	return unspecified, ErrorCouldNotParse
}

// String implements the standard Stringer interface
func (e Enum) String() string {
	text, exists := index[e]
	if !exists {
		return ""
	}

	return text
}

// IsAnyOf returns true if this enum is one of the target values.
func (e Enum) IsAnyOf(values ...Enum) bool {
	for _, val := range values {
		if e == val {
			return true
		}
	}

	return false
}
