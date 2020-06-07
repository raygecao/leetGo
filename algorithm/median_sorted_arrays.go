package algorithm

// FindMedianSortedArrays find the median of two sorted array
// Two arrays are not all empty
// e.g. FindMedianSortedArrays([1, 3], [2]) = 2.0; FindMedianSortedArrays([1, 3], [2, 4]) = 2.5
// ref link
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// k-th num is what we want, if sum len is even, k+1-th num is also what we want
	k := (len(nums1) + len(nums2) + 1) / 2
	odd := (len(nums1)+len(nums2))%2 == 1
	for k > 1 {
		// make sure len(nums1) > len(nums2)
		if len(nums1) < len(nums2) {
			nums1, nums2 = nums2, nums1
		}
		if len(nums2) == 0 {
			if odd {
				return float64(nums1[k-1])
			}
			return float64(nums1[k-1]+nums1[k]) / 2
		}
		// sth is the shorter array element to be compared while lth is the longer array's
		sth := k / 2 // must be >= 0
		if len(nums2) < sth {
			sth = len(nums2)
		}
		lth := k - sth
		// index of x-th target is x - 1, which is guaranteed in this for loop
		if nums1[lth-1] < nums2[sth-1] {
			nums1 = nums1[lth:]
			k -= lth
		} else {
			nums2 = nums2[sth:]
			k -= sth
		}
	}
	var min int
	var second bool
	for {
		if len(nums1) < len(nums2) {
			nums1, nums2 = nums2, nums1
		}
		tmp := nums1[0]
		if len(nums2) > 0 && nums2[0] < tmp {
			// nums2 has the smallest value
			tmp = nums2[0]
			nums2 = nums2[1:]
		} else {
			// nums1 has the smallest value
			nums1 = nums1[1:]
		}
		if odd {
			return float64(tmp)
		}
		if second {
			return float64(tmp+min) / 2
		}
		second = true
		min = tmp
	}

}
