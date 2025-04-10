#### Two Sum II - Input array is sorted

Given an array of integers `numbers` that is already ***sorted in non-decreasing order***, find two numbers such that they add up to a specific `target` number.

Return *the indices of the two numbers (**1-indexed**) as an integer array `answer` of size `2`, where `1 <= answer[0] < answer[1] <= numbers.length`.*

The tests are generated such that there is **exactly one solution**. You **may not** use the same element twice.

**Example 1**:
<pre><b>Input</b>: numbers = [2,7,11,15], target = 9
<b>Output</b>: [1,2]
<b>Explanation</b>: The sum of 2 and 7 is 9. Therefore index1 = 1, index2 = 2.
</pre>

**Example 2**:
<pre><b>Input</b>: numbers = [2,3,4], target = 6
<b>Output</b>: [1,3]
</pre>

**Example 3**:
<pre><b>Input</b>: numbers = [-1,0], target = -1
<b>Output</b>: [1,2]
</pre>

**Constraints**:
* <code>2 <= numbers.length <= 3 * 10<sup>4</sup></code>
* `-1000 <= numbers[i] <= 1000`
* `numbers` is sorted in **non-decreasing order**.
* `-1000 <= target <= 1000`
* The tests are generated such that there is **exactly one solution**.
