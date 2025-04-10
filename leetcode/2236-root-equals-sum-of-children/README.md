#### Root Equals Sum of Children
You are given the `root` of a **binary tree** that consists of exactly `3` nodes: the root, its left child, and its right child.

Return `true` _if the value of the root is equal to the **sum** of the values of its two children_, or `false` _otherwise_.

**Example 1**:

![](example_1.png)
<pre><b>Input</b>: root = [10,4,6]
<b>Output</b>: true
<b>Explanation</b>: The values of the root, its left child, and its right child are 10, 4, and 6, respectively.
10 is equal to 4 + 6, so we return true.
</pre>

**Example 2**:

![](example_2.png)
<pre><b>Input</b>: root = [5,3,1]
<b>Output</b>: false
<b>Explanation</b>: The values of the root, its left child, and its right child are 5, 3, and 1, respectively.
5 is not equal to 3 + 1, so we return false.
</pre>

**Constraints**:
* The tree consists only of the root, its left child, and its right child.
* `-100 <= Node.val <= 100`
