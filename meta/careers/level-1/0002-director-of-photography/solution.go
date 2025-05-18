package meta0002

func cartesianProduct(sets ...[]int32) [][]int32 {
	if len(sets) == 0 {
		return [][]int32{}
	}

	product := [][]int32{{}}
	for _, set := range sets {
		updatedPorduct := [][]int32{}
		for _, p := range product {
			for _, s := range set {
				combined := append([]int32{}, p...)
				combined = append(combined, s)
				updatedPorduct = append(updatedPorduct, combined)
			}
		}
		product = updatedPorduct
	}

	return product
}

func distance(i, j int32) int32 {
	dist := i - j
	if dist < 0 {
		return -dist
	}
	return dist
}

func getArtisticPhotographCount(N int32, C string, X, Y int32) int32 {
	pIndex := []int32{}
	aIndex := []int32{}
	bIndex := []int32{}

	for i, v := range C {
		switch v {
		case 'P':
			pIndex = append(pIndex, int32(i))
		case 'A':
			aIndex = append(aIndex, int32(i))
		case 'B':
			bIndex = append(bIndex, int32(i))
		}
	}

	photographs := cartesianProduct(pIndex, aIndex, bIndex)
	artistic := 0
	for _, photo := range photographs {
		p := photo[0]
		a := photo[1]
		b := photo[2]
		distPA := distance(p, a)
		distAB := distance(a, b)
		if ((p < a && a < b) || (b < a && a < p)) &&
			(X <= distAB && distAB <= Y) &&
			(X <= distPA && distPA <= Y) {
			artistic++
		}
	}
	return int32(artistic)
}
