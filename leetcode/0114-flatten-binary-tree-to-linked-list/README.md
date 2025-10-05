#### Flatten Binary Tree to Linked List

Given the `root` of a binary tree, flatten the tree into a "linked list":

The "linked list" should use the same `TreeNode` class where the `right` child
pointer points to the next node in the list and the `left` child pointer is
always null. The "linked list" should be in the same order as a pre-order
traversal of the binary tree.

**Example 1**:

![](example_1.jpg)

<pre><b>Input</b>: root = [1,2,5,3,4,null,6]
<b>Output</b>: [1,null,2,null,3,null,4,null,5,null,6]
</pre>

**Example 2**:

<pre><b>Input</b>: root = []
<b>Output</b>: []
</pre>

**Example 3**:

<pre><b>Input</b>: root = [0]
<b>Output</b>: [0]
</pre>

**Constraints**:

- The number of nodes in the tree is in the range `[0, 2000]`.
- `-100 <= Node.val <= 100`

Follow up: Can you flatten the tree in-place (with `O(1)` extra space)?
