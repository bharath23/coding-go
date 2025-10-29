#### Buildings With an Ocean View

There are `n` buildings in a line. You are given an integer array `heights` of
size `n` that represents the heights of the buildings in the line.

The ocean is to the right of the buildings. A building has an ocean view if the
building can see the ocean without obstructions. Formally, a building has an
ocean view if all the buildings to its right have a **smaller** height.

Return a list of indices **(0-indexed)** of buildings that have an ocean view,
sorted in increasing order.

**Example 1**:

<pre><b>Input</b>: heights = [4,2,3,1]
<b>Output</b>: [0,2,3]
<b>Explanation</b>: Building 1 (0-indexed) does not have an ocean view because building 2 is taller.
</pre>

**Example 2**:

<pre><b>Input</b>: heights = [4,3,2,1]
<b>Output</b>: [0,1,2,3]
<b>Explanation</b>: All the buildings have an ocean view.
</pre>

**Example 3**:

<pre><b>Input</b>: heights = [1,3,2,4]
<b>Output</b>: [3]
<b>Explanation</b>: Only building 3 has an ocean view.
</pre>

**Constraints**:

- <code>1 <= heights.length <= 10<sup>5</sup></code>
- <code>1 <= heights[i] <= 10<sup>9</sup></code>
