#### Valid Mountain Array
Given an array of integers `arr`, return _`true`  if and only if it is a valid mountain array_.

Recall that arr is a mountain array if and only if:
* `arr.length >= 3`
* There exists some `i` with `0 < i < arr.length - 1` such that:
  * `arr[0] < arr[1] < ... < arr[i - 1] < arr[i]`
  * `arr[i] > arr[i + 1] > ... > arr[arr.length - 1]`

**Example 1**:
<pre><code><b>Input</b>: arr = [2,1]
<b>Output</b>: false
</code></pre>

**Example 2**:
<pre><code><b>Input</b>: arr = [3,5,5]
<b>Output</b>: false
</code></pre>

**Example 3**:
<pre><code><b>Input</b>: arr = [0,3,2,1]
<b>Output</b>: true
</code></pre>

**Constraints:**
* <code>1 <= arr.length <= 10<sup>4</sup></code>
* <code>0 <= arr[i] <= 10<sup>4</sup></code>
