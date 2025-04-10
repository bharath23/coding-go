package educative0024

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		schedule [][]*Interval
		want     []*Interval
	}{
		{
			name: "test 1",
			schedule: [][]*Interval{
				{
					&Interval{start: 1, end: 2},
					&Interval{start: 5, end: 6},
				},
				{
					&Interval{start: 1, end: 3},
				},
				{
					&Interval{start: 4, end: 10},
				},
			},
			want: []*Interval{
				&Interval{start: 3, end: 4},
			},
		},
		{
			name: "test 2",
			schedule: [][]*Interval{
				{
					&Interval{start: 1, end: 3},
					&Interval{start: 6, end: 7},
				},
				{
					&Interval{start: 2, end: 4},
				},
				{
					&Interval{start: 2, end: 5},
					&Interval{start: 9, end: 12},
				},
			},
			want: []*Interval{
				&Interval{start: 5, end: 6},
				&Interval{start: 7, end: 9},
			},
		},
		{
			name: "test 3",
			schedule: [][]*Interval{
				{
					&Interval{start: 2, end: 3},
					&Interval{start: 7, end: 9},
				},
				{
					&Interval{start: 1, end: 4},
					&Interval{start: 6, end: 7},
				},
			},
			want: []*Interval{
				&Interval{start: 4, end: 6},
			},
		},
		{
			name: "test 4",
			schedule: [][]*Interval{
				{
					&Interval{start: 3, end: 5},
					&Interval{start: 8, end: 10},
				},
				{
					&Interval{start: 4, end: 6},
					&Interval{start: 9, end: 12},
				},
				{
					&Interval{start: 5, end: 6},
					&Interval{start: 8, end: 10},
				},
			},
			want: []*Interval{
				&Interval{start: 6, end: 8},
			},
		},
		{
			name: "test 5",
			schedule: [][]*Interval{
				{
					&Interval{start: 1, end: 2},
					&Interval{start: 3, end: 4},
					&Interval{start: 5, end: 6},
					&Interval{start: 7, end: 8},
					&Interval{start: 9, end: 10},
					&Interval{start: 11, end: 12},
				},
				{
					&Interval{start: 1, end: 2},
					&Interval{start: 3, end: 4},
					&Interval{start: 5, end: 6},
					&Interval{start: 7, end: 8},
					&Interval{start: 9, end: 10},
					&Interval{start: 11, end: 12},
				},
				{
					&Interval{start: 1, end: 2},
					&Interval{start: 3, end: 4},
					&Interval{start: 5, end: 6},
					&Interval{start: 7, end: 8},
					&Interval{start: 9, end: 10},
					&Interval{start: 11, end: 12},
				},
				{
					&Interval{start: 1, end: 2},
					&Interval{start: 3, end: 4},
					&Interval{start: 5, end: 6},
					&Interval{start: 7, end: 8},
					&Interval{start: 9, end: 10},
					&Interval{start: 11, end: 12},
				},
			},
			want: []*Interval{
				&Interval{start: 2, end: 3},
				&Interval{start: 4, end: 5},
				&Interval{start: 6, end: 7},
				&Interval{start: 8, end: 9},
				&Interval{start: 10, end: 11},
			},
		},
	}

	for _, test := range tests {
		have := employeeFreeTime(test.schedule)
		assert.Equalf(t, test.want, have, "%s: employeeFreeTime returned "+
			"interval is incorrect", test.name)
	}
}
