#### Serialize and Deserialize Binary tree

Serialize a given binary tree to a file and deserialize it back to a tree.
Make sure that the original and the deserialized trees are identical.

- Serialize: Write the tree to a file.
- Deserialize: Read from a file and reconstruct the tree in memory.

Serialize the tree into an array of integers, and then, deserialize it back
from the array to a tree. For simplicity’s sake, there’s no need to write the
array to the files.

**Constraints**:

- The number of nodes in the tree is in the range `[0, 500]`
- `-1000 <= Node.value <= 1000`
