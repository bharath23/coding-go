package meta0001

import (
	"sort"
)

func ceil(a, b int) int {
	return (a + b - 1) / b
}

// The maximum number of seats that can be occupied when there must be at least
// k empty seats between any two occupied seats.
//
// Examples:
//
//	n=5, k=0  → x x x x x   (all seats can be occupied)
//	n=5, k=1  → x - x - x
//	n=5, k=2  → x - - x -
//	n=5, k=3  → x - - - x
//
// General formula:
//
//	1 + floor((n - 1) / (k + 1))
//	or
//	ceil(n/(k + 1))
//
// This accounts for one initial occupied seat, followed by one additional seat
// occupied for for every (k + 1) seats remaining.
func getMaxAdditionalDinersCount(N, K int, M int32, S []int) int {
	// sort the occupied seats
	sort.Ints(S)
	additionalDiners := 0
	firstOpenSeat := 1
	for _, seat := range S {
		numOpenSeats := seat - K - firstOpenSeat
		if numOpenSeats > 0 {
			additionalDiners += ceil(numOpenSeats, K+1)
		}
		firstOpenSeat = seat + K + 1
	}

	// Calculate for any remaining seats
	numOpenSeats := N + 1 - firstOpenSeat
	if numOpenSeats > 0 {
		additionalDiners += ceil(numOpenSeats, K+1)
	}

	return additionalDiners
}
