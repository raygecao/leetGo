package algorithm

func SetZeroes(matrix [][]int) {
	if matrix == nil {
		return
	}
	var zRow, zColumn bool
	for i, line := range matrix {
		for j, ele := range line {
			if i == 0 && ele == 0 {
				zRow = true
			}
			if j == 0 && ele == 0 {
				zColumn = true
			}
			if ele == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i, line := range matrix {
		if i == 0 {
			continue
		}
		if line[0] == 0 {
			for j := range line {
				line[j] = 0
			}
		}
	}
	for j, v := range matrix[0] {
		if j == 0 {
			continue
		}
		if v == 0 {
			for i := 1; i < len(matrix); i++ {
				matrix[i][j] = 0
			}
		}
	}
	if zRow {
		for j := range matrix[0] {
			matrix[0][j] = 0
		}
	}
	if zColumn {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
}
