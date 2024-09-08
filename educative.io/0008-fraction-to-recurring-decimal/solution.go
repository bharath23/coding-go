package educative0008

import (
	"strconv"
)

func fractionToDecimal(numerator, denominator int) string {
	result := ""
	negResult := 1
	if numerator < 0 {
		negResult = -1 * negResult
		numerator = -1 * numerator
	}

	if denominator < 0 {
		negResult = -1 * negResult
		denominator = -1 * denominator
	}

	if negResult < 0 {
		result += "-"
	}

	if numerator == 0 {
		return "0"
	}

	quotient := numerator / denominator
	remainder := numerator % denominator
	result += strconv.Itoa(quotient)
	if remainder == 0 {
		return result
	}

	result += "."
	remainderMap := make(map[int]int)
	remainderMap[remainder] = len(result)
	for remainder != 0 {
		quotient = remainder * 10 / denominator
		remainder = remainder * 10 % denominator
		result += strconv.Itoa(quotient)
		if _, ok := remainderMap[remainder]; ok {
			break
		}

		remainderMap[remainder] = len(result)
	}

	if remainder != 0 {
		start := remainderMap[remainder]
		result = result[:start] + "(" + result[start:] + ")"
	}

	return result
}
