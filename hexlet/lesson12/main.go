
package solution

// BEGIN (write your solution here)
const (
	OK = iota
	CANCELLED
	UNKNOWN
)

func ErrorMessageToCode(msg string) int {
	switch msg {
		case "OK":
		 	return OK
		case "CANCELLED":
			return CANCELLED
		default:
			return UNKNOWN
	}
}
// END
