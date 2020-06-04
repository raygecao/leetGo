package algorithm

// TwoSum find two integer with the sum of `target` of an array and return those index
// e.g: TwoSum([]int{1,2,3}, 5) = int{1,2}
// ref link: https://leetcode-cn.com/problems/two-sum
// leetcode-1
func TwoSum(nums []int, target int) []int {
	// visited means map[num]index
	visited := make(map[int]int, len(nums))
	for i, num := range nums {
		if v, ok := visited[target-num]; ok {
			return []int{v, i}
		}
		// if the diff-peer value is not in visited, add this num in visited
		visited[num] = i
	}
	return nil
}
