package algorithm

func CanBeEqual(target []int, arr []int) bool {
	if len(target) != len(arr) {
		return false
	}
	counters := make(map[int]int)
	for _, ele := range arr {
		counters[ele] += 1
	}
	for _, ele := range target {
		counters[ele] -= 1
	}
	for _, counter := range counters {
		if counter != 0 {
			return false
		}
	}
	return true
}
