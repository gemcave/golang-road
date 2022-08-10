package solution

import (
	"strings"
)

// BEGIN (write your solution here)
func ModifySpaces(s, mode string) string {
	var replacer string
	switch mode {
	case "dash":
		replacer = "-"
	case "underscore":
		replacer = "_"
	default:
		replacer = "*"
	}
	return strings.ReplaceAll(s, " ", replacer)
}

// END
