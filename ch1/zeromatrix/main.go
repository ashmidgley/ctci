package main

func setZeroes(m [][]int) [][]int {
	if len(m) == 0 {
		return m
	}

	xZeroes := make(map[int]struct{})
	yZeroes := make(map[int]struct{})
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if m[i][j] == 0 {
				xZeroes[j] = struct{}{}
				yZeroes[i] = struct{}{}
			}
		}
	}

	for y := 0; y < len(m); y++ {
		for x := range xZeroes {
			m[y][x] = 0
		}
	}

	for y := range yZeroes {
		for x := 0; x < len(m[0]); x++ {
			m[y][x] = 0
		}
	}
	return m
}
