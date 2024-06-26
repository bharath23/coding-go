#### Remove N<sup>th</sup> Node From End of List
Given the  `head`  of a linked list, remove the  <code>n<sup>th</sup></code>  node from the end of the list and return its head.

**Example 1**:

![](example_1.jpg)
<pre><code><b>Input</b>: head = [1,2,3,4,5], n = 2
<b>Output</b>: [1,2,3,5]
</code></pre>

**Example 2**:
<pre><code><b>Input</b>: head = [1], n = 1
<b>Output</b>: []
</code></pre>

**Example 3**:
<pre><code><b>Input</b>: head = [1,2], n = 1
<b>Output</b>: [1]
</code></pre>

**Constraints**:
* The number of nodes in the list is  `sz`.
* `1 <= sz <= 30`
* `0 <= Node.val <= 100`
* `1 <= n <= sz`
