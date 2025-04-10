#### Height Checker

A school is trying to take an annual photo of all the students. The students are asked to stand in a single file line in  **non-decreasing order**  by height. Let this ordering be represented by the integer array  `expected`  where  `expected[i]`  is the expected height of the  `ith`  student in line.

You are given an integer array  `heights`  representing the  **current order**  that the students are standing in. Each  `heights[i]`  is the height of the  `ith`  student in line (**0-indexed**).

Return  _the  **number of indices**  where_ `heights[i] != expected[i]`.

**Example 1**:
<pre><b>Input</b>: heights = [1,1,4,2,1,3]
<b>Output</b>: 3
<b>Explanation</b>:
heights:  [1,1,4,2,1,3]
expected: [1,1,1,2,3,4]
Indices 2, 4, and 5 do not match.
</pre>

**Example 2**:
<pre><b>Input</b>: heights = [5,1,2,3,4]
<b>Output</b>: 5
<b>Explanation</b>:
heights:  [5,1,2,3,4]
expected: [1,2,3,4,5]
All indices do not match.
</pre>

**Example 3**:
<pre><b>Input</b>: heights = [1,2,3,4,5]
<b>Output</b>: 0
<b>Explanation</b>:
heights:  [1,2,3,4,5]
expected: [1,2,3,4,5]
All indices match.
</pre>

**Constraints:**
* `1 <= heights.length <= 100`
* `1 <= heights[i] <= 100`
