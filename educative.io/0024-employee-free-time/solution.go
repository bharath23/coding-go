package educative0024

import (
	"fmt"
	"sort"
)

type Interval struct {
	start int
	end   int
}

func (i *Interval) String() string {
	return fmt.Sprintf("(%d, %d)", i.start, i.end)
}

func employeeFreeTime(schedule [][]*Interval) []*Interval {
	intervals := []*Interval{}
	for _, s := range schedule {
		intervals = append(intervals, s...)
	}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i].start == intervals[j].start {
			return intervals[i].end < intervals[j].end
		}

		return intervals[i].start < intervals[j].start
	})

	res := []*Interval{}
	end := intervals[0].end
	for _, i := range intervals {
		if i.start > end {
			res = append(res, &Interval{end, i.start})
		}

		end = max(end, i.end)
	}

	return res
}
