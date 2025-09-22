#### N-Queens

The **n-queens** puzzle is the problem of placing `n` queens on an `n x n`
chessboard such that no two queens attach each other.

Given an integer `n`, return _all distinct solutions to the **n-quuens
puzzle**_. You may return the answer in **any order**.

Each soltion contains a distinct board configuration of the n-queens'
placement, where 'Q' and '.' both indicate a queen and an empty space,
respectively.

**Example 1**:

![](example_1.jpg)

<pre><b>Input</b>: n = 4
<b>Output</b>: [[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
<b>Explanation</b>: There exists two distinct solutions to the 4-queens puzzle as shown above
</pre>

**Example 2**:

<pre><b>Input</b>: n = 1
<b>Output</b>: [["Q"]]
</pre>

**Constraints**:

- `1 <= n <= 9`
