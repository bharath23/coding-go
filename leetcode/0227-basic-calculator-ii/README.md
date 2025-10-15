#### Basic Calculator II

Given a string `s` which represents an expression, _evaluate this expression
and return its value_.

The integer division should truncate toward zero.

You may assume that the given expression is always valid. All intermediate
results will be in the range of <code>[-2<sup>31</sup>, 2<sup>31</sup> - 1]</code>.

Note: You are not allowed to use any built-in function which evaluates strings
as mathematical expressions, such as `eval()`.

**Example 1**:

<pre><b>Input</b>: s = "3+2*2"
<b>Output</b>: 7
</pre>

**Example 2**:

<pre><b>Input</b>: s = " 3/2 "
<b>Output</b>: 1
</pre>

**Example 3**:

<pre><b>Input</b>: s = " 3+5 / 2 "
<b>Output</b>: 5
</pre>

**Constraints**:

- <code>1 <= s.length <= 3 \* 10<sup>5</sup></code>
- `s` consists of integers and operators `('+', '-', '*', '/')` separated by
  some number of spaces.
- `s` represents a valid expression.
- All the integers in the expression are non-negative integers in the range
  <code>[0, 2<sup>31</sup> - 1]</code>.
- The answer is **guaranteed** to fit in a **32-bit integer**.
