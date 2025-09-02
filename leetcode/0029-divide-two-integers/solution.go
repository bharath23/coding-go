package leetcode0029

import "math"

/*
 * The outer loop runs while the dividend is larger than the divisor. For each
 * iteration, the inner loop doubles the divisor until it cannot fit in the
 * current dividend. Each doubling corresponds to a power of two in the
 * quotient, and the dividend shrinks exponentially with each outer loop
 * iteration. Overall time complexity is O(log n).
 *
 * The space complexity is O(1) since no additional space is required.
 *
 * Time complexity: O(log n)
 * Space complexity: O(1)
 */

func dividev0(dividend, divisor int32) int32 {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	negatives := 0
	if dividend > 0 {
		dividend = -dividend
		negatives++
	}
	if divisor > 0 {
		divisor = -divisor
		negatives++
	}

	quotient := int32(0)
	for dividend <= divisor {
		powersOfTwo := int32(1)
		value := divisor
		for (value + value) >= dividend {
			powersOfTwo += powersOfTwo
			value += value
		}

		quotient += powersOfTwo
		dividend -= value
	}

	if negatives == 1 {
		return -quotient
	}
	return quotient
}

/*
 * The first loop builds slices of doubled divisors and corresponding powers
 * of two. It runs O(log n) times since the divisor doubles each iteration
 * until it exceeds the dividend. The second loop iterates over these slices
 * to subtract from the dividend and build the quotient, which also runs
 * O(log n) times. Overall time complexity is O(log n).
 *
 * The space complexity is O(log n) due to slices storing doubles and powers
 * of two.
 *
 * Time complexity: O(log n)
 * Space complexity: O(log n)
 */

func dividev1(dividend, divisor int32) int32 {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	negatives := 0
	if dividend > 0 {
		dividend = -dividend
		negatives++
	}
	if divisor > 0 {
		divisor = -divisor
		negatives++
	}

	doubles := make([]int32, 0, 32)
	powersOfTwo := make([]int32, 0, 32)
	powerOfTwo := int32(1)
	for divisor >= dividend {
		doubles = append(doubles, divisor)
		powersOfTwo = append(powersOfTwo, powerOfTwo)
		divisor += divisor
		powerOfTwo += powerOfTwo
	}

	quotient := int32(0)
	for i := len(doubles) - 1; i >= 0; i-- {
		if doubles[i] >= dividend {
			quotient += powersOfTwo[i]
			dividend -= doubles[i]
		}
	}

	if negatives == 1 {
		return -quotient
	}

	return quotient
}

/*
 * The first loop finds the highest double of the divisor that fits into the
 * dividend. It doubles the divisor and corresponding power of two each
 * iteration, running O(log n) times. The second loop subtracts these highest
 * doubles from the dividend while shifting them down, also running O(log n)
 * times since the highest double is halved each iteration. Overall time
 * complexity is O(log n).
 *
 * The space complexity is O(1) since no additional space is required.
 *
 * Time complexity: O(log n)
 * Space complexity: O(1)
 */

func dividev2(dividend, divisor int32) int32 {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	negatives := 0
	if dividend > 0 {
		dividend = -dividend
		negatives++
	}
	if divisor > 0 {
		divisor = -divisor
		negatives++
	}

	highestDouble := divisor
	highestPowerOfTwo := int32(1)
	for dividend <= highestDouble<<1 {
		highestPowerOfTwo <<= 1
		highestDouble <<= 1
	}

	quotient := int32(0)
	for dividend <= divisor {
		if dividend <= highestDouble {
			quotient += highestPowerOfTwo
			dividend -= highestDouble
		}
		highestPowerOfTwo >>= 1
		highestDouble >>= 1
	}

	if negatives == 1 {
		return -quotient
	}

	return quotient
}
