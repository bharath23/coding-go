#### Valid Parentheses

Given a string s containing just the characters `'('`, `')'`, `'{'`, `'}'`,
`'['` and `']'`, determine if the input string is valid.

An input string is valid if:

1. Open brackets must be closed by the same type of brackets.
2. Open brackets must be closed in the correct order.
3. Every close bracket has a corresponding open bracket of the same type.

**Example 1**:

<pre><b>Input</b>: s = "()"
<b>Output</b>: true
</pre>

**Example 2**:

<pre><b>Input</b>: s = "()[]{}"
<b>Output</b>: true
</pre>

**Example 3**:

<pre><b>Input</b>: s = "(]"
<b>Output</b>: false
</pre>

**Example 4**:

<pre><b>Input</b>: s = "([])"
<b>Output</b>: true
</pre>

**Constraints**:

- <code>1 <= s.length <= 10<sup>4</sup></code>
- `s` consists of parntheses only `'()[]{}'`.
