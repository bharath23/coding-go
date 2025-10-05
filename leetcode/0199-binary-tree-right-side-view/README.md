#### Binary Tree Right Side View

Given the `root` of a binary tree, imagine yourself standing on the **right
side** of it, return _the values of the nodes you can see ordered from top to
bottom_.

**Example 1**:

<pre><b>Input</b>: root = [1,2,3,null,5,null,4]
<b>Output</b>: [1,3,4]
<b>Explanation</b>
</pre>

![](example_1.png)

**Example 2**:

<pre><b>Input</b>: root = [1,2,3,4,null,null,null,5]
<b>Output</b>: [1,3,4,5]
<b>Explanation</b>:
</pre>

![](example_2.png)

**Example 3**:

<pre><b>Input</b>: root = [1,null,3]
<b>Output</b>: [1,3]
</pre>

**Example 4**:

<pre><b>Input</b>: root = []
<b>Output</b>: []
</pre>

**Constraints**:

- The number of nodes in the tree is in the range `[0, 100]`.
- `-100 <= Node.val <= 100`
