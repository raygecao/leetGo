package dfs

//给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
//
//示例:
//
//输入: n = 4, k = 2
//输出:
//[
//[2,4],
//[3,4],
//[2,3],
//[1,2],
//[1,3],
//[1,4],
//]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/combinations
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func Combine(n int, k int) [][]int {
	var results [][]int
	var dfs func(pos int, tmp []int)
	dfs = func(pos int, tmp []int) {
		if len(tmp) == k {
			r := make([]int, len(tmp))
			copy(r, tmp)
			results = append(results, r)
			return
		}
		if pos > n {
			return
		}

		dfs(pos+1, tmp)
		dfs(pos+1, append(tmp, pos))
	}
	dfs(1, nil)
	return results
}
