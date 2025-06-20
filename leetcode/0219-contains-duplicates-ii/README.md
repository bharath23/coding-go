#### Contains Duplicate II

Given an integer array `nums` and an integer `k`, return `true` _if there are
**two distinct indices** `i` and `j` in the array such that `nums[i] == nums[j]` and `abs(i - j) <= k`_.

**Example 1**:

<pre><b>Input</b>: nums = [1,2,3,1], k = 3
<b>Output</b>: true
</pre>

**Example 2**:

<pre><b>Input</b>: nums = [1,0,1,1], k = 1
<b>Output</b>: true
</pre>

**Example 3**:

<pre><b>Input</b>: nums = [1,2,3,1,2,3], k = 2
<b>Output</b>: false
</pre>

**Constraints**:

- <code>1 <= nums.length <= 10<sup>5</sup></code>
- <code>-10<sup>9</sup> <= nums[i] <= 10<sup>9</sup></code>
- <code>0 <= k <= 10<sup>5</sup></code>
