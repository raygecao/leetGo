package dfs

func Permute(nums []int) [][]int {
	var res [][]int
	res = dfs(0, nums, res)
	return res
}

func dfs(pos int, nums []int, res [][]int) [][]int {
	if pos == len(nums)-1 {
		result := make([]int, len(nums))
		copy(result, nums)
		res = append(res, result)
		return res
	}
	m := make(map[int]struct{})
	for i := pos; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			continue
		}
		m[nums[i]] = struct{}{}
		nums[i], nums[pos] = nums[pos], nums[i]
		res = dfs(pos+1, nums, res)
		nums[pos], nums[i] = nums[i], nums[pos]
	}
	return res
}

func Permute2(nums []int) [][]int {
	visited := make(map[int]struct{})
	return dfs2(nil, nums, visited, nil)
}

func dfs2(tmp []int, nums []int, visited map[int]struct{}, res [][]int) [][]int {
	if len(tmp) == len(nums) {
		result := make([]int, len(nums))
		copy(result, tmp)
		res = append(res, result)
		return res
	}
	for _, num := range nums {
		if _, ok := visited[num]; !ok {
			tmp = append(tmp, num)
			visited[num] = struct{}{}
			res = dfs2(tmp, nums, visited, res)
			delete(visited, num)
			tmp = tmp[:len(tmp)-1]
		}
	}
	return res
}
