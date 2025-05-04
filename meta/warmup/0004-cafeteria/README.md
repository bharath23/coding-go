#### Cafeteria

A cafeteria table consists of a row of `N` seats, numbered from 1 to N from
left to right. Social distancing guidelines require that every diner be seated
such that `K` seats to their left and `K` seats to their right (or all the
remaining seats to that side if there are fewer than `K`) remain empty.

There are currently `M` diners seated at the table, the `i`th of whom is in
seat <code>S<sub>i</sub></code>. No two diners are sitting in the same seat, and the social distancing guidelines are satisfied.

Determine the maximum number of additional diners who can potentially sit at the table without social distancing guidelines being violated for any new or existing diners, assuming that the existing diners cannot move and that the additional diners will cooperate to maximize how many of them can sit down.

_Please take care to write a solution which runs within the time limit._

**Constraints**

- <code>1 ≤ N ≤ 10<sup>15</sup></code>
- `1 ≤ K ≤ N`
- `1 ≤ M ≤ 500,000`
- `M ≤ N`
- <code>1 ≤ S<sub>i</sub> ≤ N</code>
