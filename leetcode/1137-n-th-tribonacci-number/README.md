#### N<sup>th</sup> Tribonacci Number

The Tribonacci sequence <code>T<sub>n</sub></code> is defined as follows:
<code>T<sub>0</sub> = 0, T<sub>1</sub> = 1, T<sub>2</sub> = 1 and T<sub>n+3</sub> = T<sub>n</sub> + T<sub>n+1</sub> + T<sub>n+2</sub></code> for `n >= 0`.
Given `n`, return the value of <code>T<sub>n</sub></code>.

**Example 1**:
<pre><b>Input</b>: 4
<b>Output</b>: 4
Explanation:
T<sub>3</sub> = 0 + 1 + 1 = 2
T<sub>4</sub> = 1 + 1 + 2 = 4
</pre>

**Example 2**:
<pre><b>Input</b>: 25
<b>Output</b>: 1389537
</pre>

**Constraints**
* `0 <= n <= 37`
* The answer is guaranteed to fit within a 32-bit integer, i.e. <code>answer <= 2<sup>31</sup> - 1</code>.
