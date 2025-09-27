#### Simplify Path

You are given an _absolute_ path for a Unix-style file system, which always
begins with a slash `'/'`. Your task is to transform this absolute path into
its **simplified canonical path**.

The _rules_ of a Unix-style file system are as follows:

- A single period `'.'` represents the current directory.
- A double period `'..'` represents the previous/parent directory.
- Multiple consecutive slashes such as `'//'` and `'///'` are treated as a
  single slash `'/'`.
- Any sequence of periods that **does not match** the rules above should be
  treated as a **valid directory or file name**. For example, `'...'` and
  `'....'` are valid directory or file names.

The simplified canonical path should follow these _rules_:

- The path must start with a single slash `'/'`.
- Directories within the path must be separated by exactly one slash `'/'`.
- The path must not end with a slash `'/'`, unless it is the root directory.
- The path must not have any single or double periods (`'.'` and `'..'`) used
  to denote current or parent directories.

Return the **simplified canonical path**.

**Example 1**:

<pre><b>Input</b>: path = "/home/"
<b>Output</b>: "/home"
<b>Explanation</b>: The trailing slash should be removed.
</pre>

**Example 2**:

<pre><b>Input</b>: path = "/home//foo/"
<b>Output</b>: "/home/foo"
<b>Explanation</b>: Multiple consecutive slashes are replaced by a single one.
</pre>

**Example 3**:

<pre><b>Input</b>: path = "/home/user/Documents/../Pictures"
<b>Output</b>: "/home/user/Pictures"
<b>Explanation</b>: A double period `".."` refers to the directory up a level
(the parent directory).
</pre>

**Example 4**:

<pre><b>Input</b>: path = "/../"
<b>Output</b>: "/"
<b>Explanation</b>: Going one level up from the root directory is not possible.
</pre>

**Example 5**:

<pre><b>Input</b>: path = "/.../a/../b/c/../d/./"
<b>Output</b>: "/"
<b>Explanation</b>: "..." is a valid name for a directory in this problem.
</pre>

**Constraints**:

- `1 <= path.length <= 3000`
- `path` consists of English letters, digits, period `.`, slash `/` or `_`.
- `path` is a valid absolute Unix path.
