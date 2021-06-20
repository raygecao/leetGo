package algorithm

func CountTriplets(arr []int) int {
	if len(arr) < 2 {
		return 0
	}
	counter := 0
	dp := make([][]int, len(arr))
	for i := 0; i < len(arr); i++ {
		dp[i] = make([]int, len(arr))
	}
	for i := 0; i < len(arr)-1; i++ {
		dp[i][i] = arr[i]
		for j := i + 1; j < len(arr); j++ {
			dp[i][j] = dp[i][j-1] ^ arr[j]
			if dp[i][j] == 0 {
				counter += j - i
			}
		}
	}
	return counter
}
