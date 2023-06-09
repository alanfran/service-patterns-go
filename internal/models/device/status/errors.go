package status

type Error int

// Declaring errors as const integers makes them immutable,
// so your sentinel values don't get replaced and break error handling.
const (
	errorUnspecified Error = iota
	ErrorCouldNotParse
)

func (e Error) Error() string {
	switch e {
	case errorUnspecified:
		return "unspecified device status error - this is a bug"
	case ErrorCouldNotParse:
		return "could not parse string into device status enum"
	default:
		return "unknown error in device status enum - this is a bug"
	}
}
