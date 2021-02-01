package main

func rotate(m [][]int) [][]int {
	if len(m) == 0 {
		return m
	}

	result := [][]int{}
	for i := 0; i < len(m[0]); i++ {
		row := []int{}
		for j := len(m) - 1; j >= 0; j-- {
			row = append(row, m[j][i])
		}
		result = append(result, row)
	}
	return result
}
