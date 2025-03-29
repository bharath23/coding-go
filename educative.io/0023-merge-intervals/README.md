#### Merge Intervals

We are given an array of closed intervals, `intervals`, where each interval
has a start time and an end time. The input array is sorted with respect to
the start times of each interval. For example,
`intervals` = `[[1, 4], [3,6],[7, 9]]` is sorted in terms of start times
`1, 3, 7`.

Your task is to merge the <ins>overlapping intervals</ins> and return a new
output array consisting of only the non-overlapping intervals.

**Constraints**:

- <pre><code>1 <= intervals.Length <= 10<sup>3</sup></code></pre>
- ` intervals[i].Length = 2`
- <pre><code>0 <= start time <= end time <= 10<sup>4</sup></code></pre>
