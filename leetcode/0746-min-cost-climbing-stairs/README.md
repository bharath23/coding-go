#### Min Cost Climbing Stairs

You are given an integer array `cos`t where `cost[i]` is the cost of <code>i<sup>th</sup></code> step on a staircase. Once you pay the cost, you can either climb one or two steps.
You can either start from the step with index `0`, or the step with index `1`.
Return _the minimum cost to reach the top of the floor_.

**Example 1**:
<pre><b>Input</b>: cost = [10, <u>15</u>, 20]
<b>Output</b>: 15
<b>Explanation</b>: You will start at index 1.
— Pay 15 and climb two steps to reach the top.
Total cost is 15.
</pre>

**Example 2**:
<pre><b>Input</b>: cost = [<u>1</u>, 100, <u>1</u>, 1, <u>1</u>, 100, <u>1</u>, <u>1</u>, 100, <u>1</u>]
<b>Output</b>: 6
<b>Explanation</b>: You will start at index 0.
— Pay 1 and climb two steps to reach index 2.
— Pay 1 and climb two steps to reach index 4.
— Pay 1 and climb two steps to reach index 6.
— Pay 1 and climb one step to reach index 7.
— Pay 1 and climb two steps to reach index 9.
— Pay 1 and climb one step to reach the top.
Total cost is 6.
</pre>

**Constraints**:
* `2 <= cost.length <= 1000`
* `0 <= cost[i] <= 999`
