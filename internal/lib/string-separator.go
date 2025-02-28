package lib

func StringStartsWith(targetString string, startsWith string) bool {
	if targetString == "" || len(targetString) <= len(startsWith) {
		return false
	}

	i := len(startsWith)
	if targetString[:i] == startsWith {
		return true
	}
	return false
}
