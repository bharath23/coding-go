package leetcode0042

/*
To calculate where rainwater can be trapped, we perform three linear passes
over the input array. One pass to compute the maximum height to the left of
each index. One pass to compute the maximum height to the right of each index.
A final pass to calculate the trapped water at each index using the minimum of
the left and right maximums. Each of these steps takes O(n) time, so the total
time complexity is O(n). We use two additional arrays of size n to store the
left and right maximum heights. Therefore, the space complexity is also O(n).

Time complexity: O(n)
Space complexity: O(n)
*/
func trapv0(height []int) int {
	size := len(height)
	leftMax, rightMax := make([]int, size), make([]int, size)
	leftMax[0] = 0
	rightMax[size-1] = 0
	for i := 1; i < size; i++ {
		leftMax[i] = max(leftMax[i-1], height[i-1])
	}

	for i := size - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i+1])
	}

	result := 0
	for i := 0; i < size; i++ {
		if h := min(leftMax[i], rightMax[i]) - height[i]; h > 0 {
			result += h
		}
	}
	return result
}

/*
This implementation uses a two-pointer approach:
  - Water can be trapped wherever there is a valley between taller bars.
  - If the height at the left pointer is less than the height at the right
    pointer, we rely on the leftMax to compute trapped water.
  - If the right side is lower, we use rightMax instead.
  - At each step, we move the pointer (left or right) that has the lower
    height, updating max values and accumulating trapped water accordingly.

We make a single pass through the input so the time compexity is O(n). We use
only a constant amount of extra space, so the space complexity is O(1).

Time complexity: O(n)
Space complexit: O(1)
*/
func trapv1(height []int) int {
	leftMax, rightMax := 0, 0
	left, right := 0, len(height)-1
	result := 0
	for left < right {
		if height[left] < height[right] {
			if leftMax <= height[left] {
				leftMax = height[left]
			} else {
				result += leftMax - height[left]
			}
			left++
		} else {
			if rightMax <= height[right] {
				rightMax = height[right]
			} else {
				result += rightMax - height[right]
			}
			right--
		}
	}

	return result
}
