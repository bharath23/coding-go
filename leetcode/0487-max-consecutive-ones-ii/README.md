#### Max Consecutive Ones II

Given a binary array `nums`, return _the maximum number of consecutive_ `1`_'s in the array if you can flip at most one_  `0`.

**Example 1**:
<pre><code><b>Input</b>: nums = [1,0,1,1,0]
<b>Output</b>: 4
<b>Explanation</b>: Flip the first zero will get the maximum number of consecutive 1s. After flipping, the maximum number of consecutive 1s is 4.
</code></pre>

**Example 2**:
<pre><code><b>Input</b>: nums = [1,0,1,1,0,1]
<b>Output</b>: 4
</code></pre>

**Constraints**:
* <code>1 <= nums.length <= 10<sup>5</sup></code>
* `nums[i]`  is either  `0`  or  `1`.

**Follow up**: What if the input numbers come in one by one as an infinite stream? In other words, you can't store all numbers coming from the stream as it's too large to hold in memory. Could you solve it efficiently?
