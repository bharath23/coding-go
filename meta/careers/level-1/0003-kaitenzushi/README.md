#### Kaitenzushi

There are `N` dishes in a row on a _kaiten belt_, with the `i`th dish being of
type <code>D<sub>i</sub></code>. Some dishes may be of the same type as another.

You're hungry, but you'd also like to keep things interesting. The `N` dishes
will arrive in front of you, one after another in order, and each one you'll eat
it as long as it isn't the same type as an of the previous `K` dishes you've
eaten. You eat very fast, so you can consume a dish before the next one gets to
you. Any dishes you choose not eat as they pass will be eaten by others.

Determine how many dishes you'll end up eating.

**Constraints**:

- `1 <= N <= 500,000`
- `1 <= K <= N`
- <code>1 <= D<sub>i</sub> <= 1,000,000</code>
