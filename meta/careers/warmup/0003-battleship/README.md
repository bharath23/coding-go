#### Battleship

You're playing Battleship on a grid of cells with `R` rows and `C` columns.
There are `0` or more battleships on the grid, each occupying a single distinct
cell. The cell in the `i`th row from the top and `j`th column from the left
either contains a battleship (<code>G<sub>i,j</sub> = 1</code>) or doesn't (
<code>G<sub>i,j</sub> = 0</code>).

You're going to fire a single shot at a random cell in the grid. You'll choose
this cell uniformly at random from the `R*C` possible cells. You're interested
in the probability that the cell hit by your shot contains a battleship.

Your task is to implement the function `getHitProbability(R, C, G)` which
returns this probability.

_Note: Your return value must have an absolute or relative error of at most
<code>10<sup>−6</sup></code> to be considered correct._

**Constraints**

- `1 <= R, C <= 100`
- <code>0 <= G<sub>i,j</sub> <= 1</code>
