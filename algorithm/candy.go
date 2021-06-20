package algorithm

func Candy(ratings []int) int {
	result := make([]int, len(ratings))

	var sum, score int
	for i := range ratings {
		if i == 0 {
			result[i] = 1
			continue
		}
		if ratings[i] > ratings[i-1] {
			result[i] = result[i-1] + 1
		} else {
			result[i] = 1
		}
	}

	for i := len(ratings) - 1; i >= 0; i-- {
		if i == len(ratings)-1 {
			sum += result[i]
			score = 1
			continue
		}
		if ratings[i] > ratings[i+1] {
			score += 1
		} else {
			score = 1
		}
		sum += max(score, result[i])
	}
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
