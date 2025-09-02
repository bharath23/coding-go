#### Divide Two Integers

Given two integers `dividend` and `divisor`, divide two integers **without**
using multiplication, division, and mod operator.

The integer division should truncate toward zero, which means losing its
fractional part. For example, `8.345` would be truncated to `8`, and `-2.7335`
would be truncated to `-2`.

Return the quotient after dividing dividend by divisor.

Note: Assume we are dealing with an environment that could only store integers
within the **32-bit** signed integer range:
<code>[−2<sup>31</sup>, 2<sup>31</sup> − 1]</code>. For this problem, if the
quotient is **strictly greater than** <code>2<sup>31</sup> - 1</code>, then
return <code>2<sup>31</sup> - 1</code>, and if the quotient is **strictly less
than** <code>-2<sup>31</sup></code>, then return <code>-2<sup>31</sup></code>.

**Example 1**:

<pre><b>Input</b>: dividend = 10, divisor = 3
<b>Output</b>: 3
<b>Explanation</b>: 10/3 = 3.33333.. which is truncated to 3.
</pre>

**Example 2**:

<pre><b>Input</b>: dividend = 7, divisor = -3
<b>Output</b>: -2
<b>Explanation</b>: 7/-3 = -2.33333.. which is truncated to -2.
</pre>

**Constraints**:

- <code>-2<sup>31</sup> <= dividend, divisor <= 2<sup>31</sup> - 1</code>
- `divisor != 0`
