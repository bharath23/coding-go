#### Shortest Path in Binary Matrix

Given an `n x n` binary matrix `grid`, return the _length of the shortest **clear path** in the matrix_. If there is no clear path, return `-1`.
A **clear path** in a binary matrix is a path from the **top-left** cell
(i.e., `(0, 0)`) to the **bottom-right** cell (i.e., `(n-1, m-1)`) such that:

- All the visited cells of the path are `0`
- All the adjacent cells of the path are **8-directionally** connected (i.e., they are different and they share and edge or a corner).

The **length of a clear path** is the number of visisted cells of this path.

**Example 1**:

![](example_1.png)

<pre><b>Input</b>: grid = [[0, 1], [1, 0]]
<b>Output</b>: 2
</pre>

**Example 2**:

![](example_2.png)

<pre><b>Input</b>: grid = [[0, 0, 0], [1, 1, 0], [1, 1, 0]]
<b>Output</b>: 4
</pre>

**Example 3**:

<pre><b>Input</b>: grid = [[1, 0, 0], [1, 1, 0], [1, 1, 0]]
<b>Output</b>: -1
</pre>

**Constraints**:

- `n == grid.length`
- `n == grid[i].length`
- `1 <= n <= 100`
