package lib

func StringStartsWith(targetString string, startsWith ...string) bool {
	for _, start := range startsWith {
		i := len(start)
		if targetString[:i] == start {
			return true
		}
	}
	return false
}
