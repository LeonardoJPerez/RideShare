package utils

// SearchString returns true/false if elem exists within array collection.
func SearchString(arr []string, elem string) bool {
	found := false
	if len(arr) == 0 {
		return found
	}

	if elem == "" {
		return found
	}

	for _, s := range arr {
		if s == elem {
			found = true
			break
		}
	}

	return found
}

// SearchUint returns true/false if elem exists within array collection.
func SearchUint(arr []uint, elem uint) bool {
	found := false
	if len(arr) == 0 {
		return found
	}

	for _, s := range arr {
		if s == elem {
			found = true
			break
		}
	}
	return found
}
