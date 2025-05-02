package meta0003

func getHitProbability(r, c int32, g [][]int32) float64 {
	var i, j, battleShipCells int32

	battleShipCells = 0
	for i = 0; i < r; i++ {
		for j = 0; j < c; j++ {
			battleShipCells += g[i][j]
		}
	}

	return float64(battleShipCells) / float64(r*c)
}
