#### Employee Free Time

You’re given a list containing the schedules of multiple employees. Each
person’s schedule is a list of non-overlapping intervals in sorted order. An
interval is specified with the start and end time, both being positive
integers. Your task is to find the list of finite intervals representing the
free time for all the employees.

> Note: The common free intervals are calculated between the earliest start
> time and the latest end time of all meetings across all employees.

**Constraints**:

- `1 <= schedule.length, schedule[i].length <= 50`
- <pre><code>0 <= interval.start < interval.end <= 10<sup>8</sup></code></pre>
