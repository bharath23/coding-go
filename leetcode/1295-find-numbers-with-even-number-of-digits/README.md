#### Find Numbers with Even Number of digits

Given an array `nums` of integers, return how many of them contain an **even number** of digits.

**Example 1**:
<pre><code><b>Input</b>: nums = [12,345,2,6,7896]
<b>Output</b>: 2
<b>Explanation</b>:
12 contains 2 digits (even number of digits).
345 contains 3 digits (odd number of digits).
2 contains 1 digit (odd number of digits).
6 contains 1 digit (odd number of digits).
7896 contains 4 digits (even number of digits).
Therefore only 12 and 7896 contain an even number of digits.
</code></pre>

**Example 2:**
<pre><code><b>Input</b>: nums = [555,901,482,1771]
<b>Output</b>: 1
<b>Explanation</b>:
Only 1771 contains an even number of digits.
</code></pre>

**Constraints:**
* `1<= nums.length <= 500`
*  <code>1 <= nums[i] <= 10<sup>5</sup><code>
