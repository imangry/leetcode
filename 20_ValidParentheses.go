package main



func isValid(s string) bool {

	left := make(map[rune]rune)
	left['{'] = '}'
	left['('] = ')'
	left['['] = ']'

	right := make(map[rune]rune)
	right['}'] = '{'
	right[')'] = '('
	right[']'] = '['

	stack := make([]rune, len(s))
	ptr := 0

	for _, value := range s {
		if _, ok := left[value]; ok {
			stack[ptr] = value
			ptr ++
		}

		if r, ok := right[value]; ok {
			if ptr == 0 {
				return false
			}
			if stack[ptr-1] != r {
				return false
			}
			ptr--
		}
	}
	if ptr != 0 {
		return false
	}
	return true
}
