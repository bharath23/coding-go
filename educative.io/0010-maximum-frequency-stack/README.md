#### Maximum Frequency Stack

Design a stack-like data structure. You should be able to push elements to this
data structure and pop elements with maximum frequency.

You'll need to implement **FreqStack** struct that should consist of the
following:

- **FreqStack**: This a struct used to declare a frequency stack.
- **Push(value)**: This is used to push an integer data onto top of the stack.
- **Pop()**: This is used to remove and return the most frequent element in the
  stack.

> **Note**: If there is a tie for the most frequent element, then most
> recently pushed element is removed and returned.

**Constraints**:

- <code>0 <= value <= 10<sup>4</sup></code>
- At most, <code>2 x 10<sup>3</sup><code> calls will be made to **Push()** and
  **Pop()**.
- It is guaranteed that there will be at least one element in the stack before
  calling **Pop()**.
