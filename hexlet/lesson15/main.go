
package solution

// BEGIN (write your solution here)
func Map(strs []string, mapFunc func(s strin) string) []string {
	mapped := make([]string, len(strs))
	for i, s := range strs {
		mapped[i] = mapFunc(s)
	}

	return mapped
}

// END
