package educative0006

type Set struct {
	value rune
	index int
}

type Stack []Set

// Push pushes the new element into the stack
func (s Stack) Push(v Set) Stack {
	return append(s, v)
}

// Pop pops the last pushed element from the stack
func (s Stack) Pop() (Stack, Set) {

	l := len(s)
	if l == 0 {
		return s, Set{0, 0}
	}
	return s[:l-1], s[l-1]
}

// Top returns the top value of the stack
func (s Stack) Top() Set {
	l := len(s)
	return s[l-1]
}

// Empty checks if the stack is empty or not
func (s Stack) Empty() bool {
	return len(s) == 0
}

func minRemoveParentheses(s string) string {
	var stack Stack
	for i, r := range s {
		switch r {
		case '(':
			set := Set{
				value: r,
				index: i,
			}
			stack = stack.Push(set)
		case ')':
			if !stack.Empty() {
				set := stack.Top()
				if set.value == '(' {
					stack, _ = stack.Pop()
				} else {
					set := Set{
						value: r,
						index: i,
					}
					stack = stack.Push(set)
				}
			} else {
				set := Set{
					value: r,
					index: i,
				}
				stack = stack.Push(set)
			}
		}
	}

	result := []rune{}
	idx := 0
	for i, r := range s {
		if !stack.Empty() && idx < len(stack) {
			if stack[idx].index == i {
				idx++
				continue
			}
		}
		result = append(result, r)
	}
	return string(result)
}
