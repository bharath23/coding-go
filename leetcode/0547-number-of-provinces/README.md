#### Number of Provinces
There are `n` cities. Some of them are connected, while some are not. If city `a` is connected directly with city `b`, and city `b` is connected directly with city `c`, then city `a is connected indirectly with city c`.

A **province** is a group of directly or indirectly connected cities and no other cities outside of the group.

You are given an `n x n` matrix `isConnected` where `isConnected[i][j] = 1` if the <code>i<sup>th</sup></code> city and the <code>j<sup>th</sup></code> city are directly connected, and `isConnected[i][j] = 0` otherwise.

Return _the total number of **provinces**_.

**Example 1**:

![](example_1.jpg)
<pre><b>Input</b>: isConnected = [[1,1,0],[1,1,0],[0,0,1]]
<b>Output</b>: 2
</pre>

**Example 2**:

![](example_2.jpg)
<pre><b>Input</b>:  isConnected = [[1,0,0],[0,1,0],[0,0,1]]
<b>Output</b>: 3
</pre>

**Constraints**:
* `1 <= n <= 200`
* `n == isConnected.length`
* `n == isConnected[i].length`
* `isConnected[i][j]`  is  `1`  or  `0`.
* `isConnected[i][i] == 1`
* `isConnected[i][j] == isConnected[j][i]`
