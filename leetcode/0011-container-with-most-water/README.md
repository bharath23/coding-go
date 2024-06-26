#### Container With Most Water

Given `n` non-negative integers `a1, a2, ..., an`, where each represents a point
at coordinate `(i, ai)`. `n` vertical lines are drawn such that the two
endpoints of the line `i` is at `(i, ai)` and `(i, 0)`. Find two lines, which,
together with the x-axis forms a container, such that the container contains the
most water.

**Notice** that you may not slant the container.

**Example 1**:

![](example_1.jpg)
<pre><code><b>Input</b>: height = [1,8,6,2,5,4,8,3,7]<b>Output</b>: 49
<b>Explanation</b>: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7].
             In this case, the max area of water (blue section) the container can contain
	         is 49.
</code></pre>

**Example 2**:
<pre><code><b>Input</b>: height = [1,1]
<b>Output</b>: 1
</code></pre>

**Example 3**:
<pre><code><b>Input</b>: height = [4,3,2,1,4]
<b>Output</b>: 16
</code></pre>

**Example 4**:
<pre><code><b>Input</b>: height = [1,2,1]
<b>Output</b>: 2
</code></pre> 

**Constraints**:
* `n == height.length`
* `2 <= n <= 105`
* `0 <= height[i] <= 104`
