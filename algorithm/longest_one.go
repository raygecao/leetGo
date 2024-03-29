package algorithm

//给定一个由若干 0 和 1 组成的数组 A，我们最多可以将 K 个值从 0 变成 1 。
//
//返回仅包含 1 的最长（连续）子数组的长度。
//
//
//
//示例 1：
//
//输入：A = [1,1,1,0,0,0,1,1,1,1,0], K = 2
//输出：6
//解释：
//[1,1,1,0,0,1,1,1,1,1,1]
//粗体数字从 0 翻转到 1，最长的子数组长度为 6。
//示例 2：
//
//输入：A = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], K = 3
//输出：10
//解释：
//[0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
//粗体数字从 0 翻转到 1，最长的子数组长度为 10。
//
//
//提示：
//
//1 <= A.length <= 20000
//0 <= K <= A.length
//A[i] 为 0 或 1
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/max-consecutive-ones-iii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func LongestOnes(A []int, K int) int {
	if K >= len(A) {
		return len(A)
	}
	if K == 0 {
		return continue1(A)
	}
	i, j := 0, K-1
	var num0 int
	max := K
	for k := i; k <= j; k++ {
		if A[k] == 0 {
			num0++
		}
	}
	remind := K - num0
	j++
	for j < len(A) {
		if A[j] == 0 {
			remind--
		}
		if remind < 0 {
			for A[i] != 0 {
				i++
			}
			remind++
			i++
		}
		if j-i+1 > max {
			max = j - i + 1
		}
		j++
	}
	return max
}

func continue1(A []int) int {
	sum, max := 0, 0
	for _, i := range A {
		if i == 1 {
			sum++
			continue
		}
		if sum > max {
			max = sum
		}
		sum = 0
	}
	if sum > max {
		max = sum
	}
	return max
}
