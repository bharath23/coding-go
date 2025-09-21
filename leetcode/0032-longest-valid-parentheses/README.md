#### Longest Valid Parentheses

Given a string containing just the characters `'('` and `')'`, return _the
length of the longest valid (well-formed) parentheses substring_.

**Example 1**:

<pre><b>Input</b>: s = "(()"
<b>Output</b>: 2
<b>Explanation</b>: The longest valid parentheses substring is "()"
</pre>

**Example 2**:

<pre><b>Input</b>: s = ")()())"
<b>Output</b>: 4
<b>Explanation</b>: The longest valid parentheses substring is "()()"
</pre>

**Example 3**:

<pre><b>Input</b>: s = ""
<b>Output</b>: 0
</pre>

**Constraints**:

- <code>0 <= s.length <= 3 \* 10<sup>4</sup></code>
- `s[i]` is `'('`, or `')'`.
