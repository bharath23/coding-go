#### Maximum Subarray

Given an integer array `nums`, find the subarray with the largest sum, and
return _its sum_.

**Example 1**:

<pre><b>Input</b>: nums = [-2,1,-3,4,-1,2,1,-5,4]
<b>Output</b>: 6
<b>Explanation</b>: The subarray [4,-1,2,1] has the largest sum 6.
</pre>

**Example 2**:

<pre><b>Input</b>: nums = [1]
<b>Output: 1</b>
<b>Explanation</b>: The subarray [1] has the largest sum 1.
</pre>

**Example 3**:

<pre><b>Input<b>: nums = [5,4,-1,7,8]
<b>Output</b>: 23
<b>Explanation<b>: The subarray [5,4,-1,7,8] has the largest sum 23.
</pre>

**Constraints**:

- <code>1 <= nums.length <= 10<sup>5</sup></code>
- <code>-10<sup>4</sup> <= nums[i] <= 10<sup>4</sup></code>

Follow up: If you have figured out the `O(n)` solution, try coding another
solution using the divide and conquer approach, which is more subtle.
