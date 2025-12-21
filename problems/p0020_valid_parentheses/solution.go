package p0020_valid_parentheses

func matchCase(c rune) rune {

	switch c {
	case ')':
		return '('
	case ']':
		return '['
	case '}':
		return '{'
	}

	return 0
}

func isValid(s string) bool {
	// Create a stack
	st := make([]rune, 0)

	for _, char := range s {

		if len(st) > 0 && st[len(st)-1] == matchCase(char) {
			st = st[:len(st)-1]
			continue
		}

		st = append(st, char)
	}

	return len(st) == 0
}
