#### Number of Steps to Reduce a Number to Zero

Given an integer `num`, return _the number of steps to reduce it to zero_.

In one step, if the current number is even, you have to divide it by `2`, otherwise, you have to subtract `1` from it.

**Example** 1:
<pre><b>Input</b>: num = 14
<b>Output</b>: 6
<b>Explanation</b>: 
Step 1) 14 is even; divide by 2 and obtain 7. 
Step 2) 7 is odd; subtract 1 and obtain 6.
Step 3) 6 is even; divide by 2 and obtain 3. 
Step 4) 3 is odd; subtract 1 and obtain 2. 
Step 5) 2 is even; divide by 2 and obtain 1. 
Step 6) 1 is odd; subtract 1 and obtain 0.
</pre>

**Example** 2:
<pre><b>Input</b>: num = 8
<b>Output</b>: 4
<b>Explanation</b>: 
Step 1) 8 is even; divide by 2 and obtain 4. 
Step 2) 4 is even; divide by 2 and obtain 2. 
Step 3) 2 is even; divide by 2 and obtain 1. 
Step 4) 1 is odd; subtract 1 and obtain 0.
</pre>

**Example 3**:
<pre><b>Input</b>: num = 123
<b>Output</b>: 12
</pre>

**Constraints**:
* <code>0 <= num <= 10<sup>6</sup></code>
