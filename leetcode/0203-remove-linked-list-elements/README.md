#### Remove Linked List Elements
Given the head of a linked list and an integer val, remove all the nodes of the linked list that has Node.val == val, and return the new head.

**Example 1**:

![](example_1.jpg)
<pre><code><b>Input</b>: head = [1,2,6,3,4,5,6], val = 6
<b>Output</b>: [1,2,3,4,5]
</code></pre>

**Example 2**:
<pre><code><b>Input</b>: head = [], val = 1
<b>Output</b>: []
</code></pre>

**Example 3**:
<pre><code><b>Input</b>: head = [7,7,7,7], val = 7
<b>Output</b>: []
</code></pre>

**Constraints**:
* The number of nodes in the list is in the range  <code>[0, 10<sup>4</sup>]</code>.
* `1 <= Node.val <= 50`
* `0 <= val <= 50`
