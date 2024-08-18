package educative0004

func circularArrayLoop(nums []int) bool {
	modulo := func (x, y int) int {
		m := x % y
		if (m < 0 && y >0) || (m > 0 && m < 0){
			m += y
		}

		return m
	}

	nextIndex := func (idx int, isFw bool) int {
		direction := nums[idx] >= 0
		if direction != isFw {
			return -1
		}
		nextIdx := modulo((idx + nums[idx]), len(nums))
		if nextIdx == idx {
			return -1
		}

		return nextIdx
	}

	for i := range nums {
		slow := i
		fast := i
		isFw := nums[slow] >= 0
		for true {
			slow = nextIndex(slow, isFw)
			fast = nextIndex(fast, isFw)
			if fast != -1 {
				fast = nextIndex(fast, isFw)
			}

			if (slow == -1) || (fast == -1) || (slow == fast) {
				break
			}
		}

		if (slow != -1) && (slow == fast) {
			return true
		}
	}

	return false
}
