#### Insert into a Sorted Circular Linked List
Given a Circular Linked List node, which is sorted in ascending order, write a function to insert a value  `insertVal`  into the list such that it remains a sorted circular list. The given node can be a reference to any single node in the list and may not necessarily be the smallest value in the circular list.

If there are multiple suitable places for insertion, you may choose any place to insert the new value. After the insertion, the circular list should remain sorted.

If the list is empty (i.e., the given node is  `null`), you should create a new single circular list and return the reference to that single node. Otherwise, you should return the originally given node.

**Example 1**:

![](example_1.jpg)
<pre><code><b>Input</b>: head = [3,4,1], insertVal = 2
<b>Output</b>: [3,4,1,2]
<b>Explanation</b>: In the figure above, there is a sorted circular list of three elements. You are given a reference to the node with value 3, and we need to insert 2 into the list. The new node should be inserted between node 1 and node 3. After the insertion, the list should look like this, and we should still return node 3.
</code>
<img src="example_1_after.jpg" /></pre>

**Example 2**:
<pre><code><b>Input</b>: head = [], insertVal = 1
<b>Output</b>: [1]
<b>Explanation</b>: The list is empty (given head is null). We create a new single circular list and return the reference to that single node.
</code></pre>

**Example 3**:
<pre><code><b>Input</b>: head = [1], insertVal = 0
<b>Output</b>: [1,0]
</code></pre>

**Constraints**:
* The number of nodes in the list is in the range <code>[0, 5 * 10<sup>4</sup>]</code>.
* <code>-10<sup>6</sup> <= Node.val, insertVal <= 10<sup>6</sup></code>
