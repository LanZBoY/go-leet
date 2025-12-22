package p0150_evaluate_reverse_polish_notation

import "strconv"

func evalRPN(tokens []string) int {
	s := make([]int, 0)

	oMap := map[string]func(a, b int) int{
		"+": func(a, b int) int { return a + b },
		"-": func(a, b int) int { return a - b },
		"*": func(a, b int) int { return a * b },
		"/": func(a, b int) int { return a / b },
	}

	for _, t := range tokens {

		if !isOperater(t) {
			i, _ := strconv.Atoi(t)
			s = append(s, i)
		} else {
			var a = 0
			var b = 0
			b, s = pull(s)
			a, s = pull(s)
			r := oMap[t](a, b)

			s = append(s, r)
		}

	}

	return s[0]
}

func pull(s []int) (int, []int) {
	data := s[len(s)-1]
	s = s[:len(s)-1]
	return data, s
}

func isOperater(token string) bool {
	oMap := map[string]struct{}{
		"+": {},
		"-": {},
		"*": {},
		"/": {},
	}
	_, ok := oMap[token]

	return ok
}
