package leetcode0020

type stack struct {
	values []rune
	size   int
}

func newStack() *stack {
	return &stack{
		values: make([]rune, 0),
		size:   0,
	}
}

func (s *stack) isEmpty() bool {
	return s.size == 0
}

func (s *stack) pop() rune {
	s.size--
	v := s.values[s.size]
	s.values = s.values[:s.size]
	return v
}

func (s *stack) push(r rune) {
	s.values = append(s.values, r)
	s.size++
}

/*
A stack based solution. Traverse the string character by character:
	if it is an open paratheses, push it to stack
	if it is close parantheses, pop from stack and check if it is matching

Complexity:
Time complexity: O(n)
Space complexity: O(n)
*/

func isValid(str string) bool {
	s := newStack()
	for _, r := range str {
		if r == '(' || r == '[' || r == '{' {
			s.push(r)
			continue
		}

		if s.isEmpty() {
			return false
		}

		p := s.pop()
		switch r {
		case ')':
			if p != '(' {
				return false
			}
		case ']':
			if p != '[' {
				return false
			}
		case '}':
			if p != '{' {
				return false
			}
		}
	}

	if !s.isEmpty() {
		return false
	}

	return true
}
