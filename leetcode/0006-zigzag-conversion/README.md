#### ZigZag Conversion

The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of
rows like this: (you may want to display this pattern in a fixed font for
better legibility)

<pre><code>P   A   H   N
A P L S I I G
Y   I   R
</code></pre>

And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number
of rows:

<pre><code>string convert(string s, int numRows);
</code></pre>

**Example 1**:
<pre><code><b>Input</b>: s = "PAYPALISHIRING", numRows = 3
<b>Output</b>: "PAHNAPLSIIGYIR"
</code></pre>

**Example 2**:
<pre><code><b>Input</b>: s = "PAYPALISHIRING", numRows = 4
<b>Output</b>: "PINALSIGYAHRPI"
<b>Explanation</b>:

P     I    N
A   L S  I G
Y A   H R
P     I
</code></pre>
