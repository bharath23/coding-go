package educative0003

func isHappy(num int) bool {
	sumSquareDigits := func(n int) int {
		sum := 0
		for n > 0 {
			d := n % 10
			n = n / 10
			sum += d * d
		}
		return sum
	}

	slow := num
	fast := num
	for {
		slow = sumSquareDigits(slow)
		fast = sumSquareDigits(sumSquareDigits(fast))
		if fast == 1 {
			return true
		}
		if slow == fast {
			return false
		}
	}
}
