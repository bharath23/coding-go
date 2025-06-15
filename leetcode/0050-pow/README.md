#### Pow(x,n)

Implement `pow(x, n)`, which calculates `x` raised to the power `n` (i.e.,
<code>x<sup>n</sup></code>).

**Example 1**:

<pre><b>Input</b>: x = 2.00000, n = 10
<b>Output</b>: 1024.00000
</pre>

**Example 2**:

<pre><b>Input</b>: x = 2.10000, n = 3
<b>Output</b>: 9.26100
</pre>

**Example 3**:

<pre><b>Input</b>: x = 2.00000, n = -2
<b>Output</b>: 0.25000
<b>Explanation</b>: 2<sup>-2</sup> = 1/2<sup>2</sup> = 1/4 = 0.25
</pre>

**Constraints**:

- `-100.0 < x < 100.0`
- <code>-2<sup>31</sup> <= n <= 2<sup>31</sup>-1</code>
- `n` is an integer
- Either `x` is not zero or `n > 0`
- <code>-10<sup>4</sup> <= x<sup>n</sup> <= 10<sup>4</sup></code>
