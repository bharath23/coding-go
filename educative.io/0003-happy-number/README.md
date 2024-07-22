#### Happy Number

Write an algorithm to determine if a number `num` is a happy number.

We use the following process to check if a given number is a happy number:
* Starting with the given number `num`, replace the number with the sum of the squares of its digits.
* Repeat the process until:
- The number equals `1`, which will depict that the given number `num` is a happy number.
- It enters a cycle, which will depict that the given number `num` is not a happy number.

Return TRUE if `num` is a happy number, and FALSE if not.

**Constraints**:
* <code>1 ≤ num ≤ 2<sup>31</sup> −1</code>
