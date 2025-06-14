#### Merge k Sorted Lists

You are given an array of `k` linked-lists `lists`, each linked-list is sorted
in ascending order.

_Merge all the linked-lists into one sorted linked-list and return it_

**Example 1**:

<pre><b>Input</b>: lists = [[1, 4, 5], [1, 3, 4], [2, 6]]
<b>Output</b>: [1, 1, 2, 3, 4, 4, 5, 6]
<b>Explanation</b>: The linked-list are:
[
    1->4->5,
    1->3->4,
    2->6
]
merging them into sroted list:
1->1->2->3->4->4->5->6
</pre>

**Example 2**:

<pre><b>Input</b>: lists = []
<b>Output</b>: []
</pre>

**Example 3**:

<pre><b>Input</b>: lists = [[]]
<b>Output</b>: []
</pre>

**Constraints**:

- `k == lists.length`
- <code>0 <= k <= 10<sup>4</sup></code>
- `0 <= lists[i].length <= 500`
- <code>-10<sup>4</sup> <= lists[i][j] <= 10<sup>4</sup><code>
- The sum of the `lists[i].length` will not exceed <code>10<sup>4</sup></code>
