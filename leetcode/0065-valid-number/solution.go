package leetcode0065

/*
We iterate through the input string exactly once, performing a constant amount
for each iteration. This results in a time complexity of O(n). Since no
additional data structures are used that grow with input size, the space
complexity remains O(1).

Time complexity: O(n)
Space complexity: O(1)
*/
func isNumberv0(s string) bool {
	var prev rune
	seenDigit := false
	seenDot := false
	seenExp := false
	for i, cur := range s {
		switch {
		case cur == '+' || cur == '-':
			if i != 0 && prev != 'e' && prev != 'E' {
				return false
			}
		case cur == '.':
			if seenDot || seenExp {
				return false
			}
			seenDot = true
		case '0' <= cur && cur <= '9':
			seenDigit = true
		case cur == 'e' || cur == 'E':
			if seenExp || !seenDigit {
				return false
			}
			seenExp = true
			seenDigit = false
		default:
			// invalid input
			return false
		}
		prev = cur
	}
	return seenDigit
}

/*
We iterate through the input string exactly once, performing a constant amount
for each iteration. This results in a time complexity of O(n). The cost of
the transitions is constant and takes the same space regarless of the legnth of
the string. The space complexity  remains O(1).

Time complexity: O(n)
Space complexity: O(1)
*/

func isNumberv1(s string) bool {
	transitions := []map[string]int{
		{"space": 1, "sign": 2, "digit": 3, "dot": 4},
		{"sign": 2, "digit": 3, "dot": 4},
		{"digit": 3, "dot": 4},
		{"digit": 3, "dot": 5, "exp": 6},
		{"digit": 5},
		{"digit": 5, "exp": 6},
		{"sign": 7, "digit": 8},
		{"digit": 8},
		{"digit": 8},
	}

	cur := 0
	trans := ""
	var ok bool
	for _, c := range s {
		switch {
		case c == ' ':
			trans = "space"
		case c == '-' || c == '+':
			trans = "sign"
		case '0' <= c && c <= '9':
			trans = "digit"
		case c == '.':
			trans = "dot"
		case c == 'e' || c == 'E':
			trans = "exp"
		default:
			return false
		}

		cur, ok = transitions[cur][trans]
		if !ok {
			return false
		}
	}

	return cur == 3 || cur == 5 || cur == 8
}
