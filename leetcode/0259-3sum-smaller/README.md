 #### 3Sum Smaller

Given an array of `n` integers `nums` and an integer `target`, find the number of index triplets `i`, `j`, `k` with `0 <= i < j < k < n` that satisfy the condition `nums[i] + nums[j] + nums[k] < target`.

**Example 1**:
<pre><b>Input</b>: nums = [-2,0,1,3], target = 2
<b>Output</b>: 2
<b>Explanation</b>: Because there are two triplets which are less that 2:
[-2,0,1]
[-2,0,3]
</pre>

**Example 2**:
<pre><b>Input</b>: nums = [], target = 0
<b>Output</b>: 0
</pre>

**Example 3**:
<pre><b>Input</b>: nums = [0], target = 0
<b>Output</b>: 0
</pre>

**Constraints**:
* `0 <= nums.length <= 3500`
* `-100 <= nums[i] <= 100`
* `-100 <= target <= 100`
