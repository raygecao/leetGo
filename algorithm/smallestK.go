package algorithm

//设计一个算法，找出数组中最小的k个数。以任意顺序返回这k个数均可。
//
//示例：
//
//输入： arr = [1,3,5,7,2,4,6,8], k = 4
//输出： [1,2,3,4]
//提示：
//
//0 <= len(arr) <= 100000
//0 <= k <= min(100000, len(arr))

func SmallestK(arr []int, k int) []int {
	if k == 0 {
		return nil
	}
	karr := arr[:k]
	buildHeap(karr)
	for _, val := range arr[k:] {
		if val < karr[0] {
			karr[0] = val
			sink(karr, 0)
		}
	}
	return karr
}

func buildHeap(arr []int) {
	if len(arr) <= 1 {
		return
	}
	for i := len(arr)/2 - 1; i >= 0; i-- {
		sink(arr, i)
	}
}

func sink(arr []int, i int) {
	chid := 2*i + 1
	for chid < len(arr) {
		// get max index of childs
		if chid+1 < len(arr) && arr[chid+1] > arr[chid] {
			chid++
		}
		if arr[i] >= arr[chid] {
			return
		}
		arr[i], arr[chid] = arr[chid], arr[i]
		i, chid = chid, 2*chid+1
	}
}
