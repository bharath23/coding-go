#### 3Sum

Given an integer array nums, return all the triplets  `[nums[i], nums[j], nums[k]]`  such that  `i != j`,  `i != k`, and  `j != k`, and  `nums[i] + nums[j] + nums[k] == 0`.

Notice that the solution set must not contain duplicate triplets.

**Example 1**:
<pre><b>Input</b>: nums = [-1,0,1,2,-1,-4]
<b>Output</b>: [[-1,-1,2],[-1,0,1]]
</pre>

**Example 2**:
<pre><b>Input</b>: nums = []
<b>Output</b>: []
</pre>

**Example 3**:
<pre><b>Input</b>: nums = [0]
<b>Output</b>: []
</pre>

**Constraints**:
* `0 <= nums.length <= 3000`
* <code>-10<sup>5</sup>  <= nums[i] <= 10<sup>5</sup></code>
