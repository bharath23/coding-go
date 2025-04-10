#### Linked List
Given  `head`, the head of a linked list, determine if the linked list has a cycle in it.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the `next` pointer. Internally,  `pos` is used to denote the index of the node that tail's `next` pointer is connected to. **Note that `pos` is not passed as a parameter**.

**Do not modify**  the linked list.

**Example 1**:

![](example_1.png)
<pre><b>Input</b>: head = [3,2,0,-4], pos = 1
<b>Output</b>: true
<b>Explanation</b>: There is a cycle in the linked list, where the tail connects to the 1st node (0-indexed).
</pre>

**Example 2**:

![](example_2.png)
<pre><b>Input</b>: head = [3,2,0,-4], pos = 1
<b>Output</b>: true
<b>Explanation</b>: There is a cycle in the linked list, where the tail connects to the 1st node (0-indexed).
</pre>

**Example 3**:

![](example_3.png)
<pre><b>Input</b>: head = [3,2,0,-4], pos = 1
<b>Output</b>: true
<b>Explanation</b>: There is a cycle in the linked list, where the tail connects to the 1st node (0-indexed).
</pre>

**Constraints**:
* The number of the nodes in the list is in the range  <code>[0, 10<sup>4</sup>]</code>.
* <code>-10<sup>5</sup>  <= Node.val <= 110<sup>5</sup></code>
* `pos`  is  `-1`  or a  **valid index**  in the linked-list.
