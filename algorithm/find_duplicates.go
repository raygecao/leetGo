package algorithm

//给定一个整数数组 a，其中1 ≤ a[i] ≤ n （n为数组长度）, 其中有些元素出现两次而其他元素出现一次。
//
//找到所有出现两次的元素。
//
//你可以不用到任何额外空间并在O(n)时间复杂度内解决这个问题吗？
//
//示例：
//
//输入:
//[4,3,2,7,8,2,3,1]
//
//输出:
//[2,3]

func FindDuplicates(nums []int) []int {
	var result []int
	for i := 0; i < len(nums); i++ {
		v := abs(abs(nums[i]))
		if nums[v-1] < 0 {
			result = append(result, v)
		}
		nums[v-1] = -1 * nums[v-1]
	}
	return result
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -1 * a
}
