package leetcode0227

/*
 * The solution scans each character in the input string once. Stack operations
 * (push and update) take constant time O(1). Thus, the total time complexity
 * is O(n), where n is the length of the string.
 *
 * The stack may store up to O(n) elements in the worst case (e.g., all '+' or
 * '-'). Therefore, the space complexity is O(n).
 *
 * Time Complexity: O(n)
 * Space Complexity: O(n)
 */
func calculatev0(s string) int {
	stack := []int{}
	num := 0
	op := '+'
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= '0' && c <= '9' {
			num = num*10 + int(c-'0')
		}

		if (c < '0' || c > '9') && c != ' ' || i == len(s)-1 {
			switch op {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] = stack[len(stack)-1] * num
			case '/':
				stack[len(stack)-1] = stack[len(stack)-1] / num
			}
			op = rune(c)
			num = 0
		}
	}

	result := 0
	for _, v := range stack {
		result += v
	}

	return result
}

/*
 * The solution scans each character in the input string once. Thus, the total
 * time complexity is O(n), where n is the length of the string.
 *
 * We need no additional data structures to store any result. So the space
 * complexity is O(1).
 *
 * Time Complexity: O(n)
 * Space Complexity: O(1)
 */
func calculatev1(s string) int {
	lastNum, curNum := 0, 0
	op := '+'
	result := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= '0' && c <= '9' {
			curNum = curNum*10 + int(c-'0')
		}

		if (c < '0' || c > '9') && c != ' ' || i == len(s)-1 {
			switch op {
			case '+':
				result += lastNum
				lastNum = curNum
			case '-':
				result += lastNum
				lastNum = -curNum
			case '*':
				lastNum = lastNum * curNum
			case '/':
				lastNum = lastNum / curNum
			}
			op = rune(c)
			curNum = 0
		}
	}

	result += lastNum

	return result
}
