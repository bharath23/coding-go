There's a multiple-choice test with `N` questions, numbered from `1` to `N`.
Each question has `2` answer options, labelled _A_ and _B_. You know that the
correct answer for the `i`th question is the `i`th character in the string `C`,
which is either "A" or "B", but you want to get a score of `0` on this test by
answering every question incorrectly.

Your task is to implement the function `getWrongAnswers(N, C)` which returns a
string with `N` characters, the `i`th of which is the answer you should give
for question `i` in order to get it wrong (either "A" or "B").

**Constraints**

- `1 ≤ N ≤ 100`
- <code>C<sub>i</sub> ∈ { "A", "B" }</code>
