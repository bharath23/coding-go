Given an array of integers `nums` sorted in non-decreasing order, find the
starting and ending position of a given `target` value.

If `target` is not found in the array, return `[-1, -1]`.

You must write an algorithm with `O(log n)` runtime complexity.

**Example 1**:

<pre><b>Input</b>: nums = [5,7,7,8,8,10], target = 8
<b>Output</b>: [3,4]
</pre>

**Example 2**:

<pre><b>Input</b>: nums = [5,7,7,8,8,10], target = 6
<b>Output</b>: [-1,-1]
</pre>

**Example 3**:

<pre><b>Input</b>: nums = [], target = 0
<b>Output</b>: [-1,-1]
</pre>

**Constraints**:

- `0 <= nums.length <= 105`
- <code> -10<sup>9</sup> <= nums[i] <= 10<sup>9</sup></code>
- `nums` is a non-decreasing array.
- <code> -10<sup>9</sup> <= target <= 10<sup>9</sup></code>
