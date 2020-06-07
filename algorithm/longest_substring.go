package algorithm

// lengthOfLongestSubstring gets the longest substring's length which is non-repeatable
// e.g. LengthOfLongestSubstring("abcabcbb") = 3
// ref link: https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
// leetcode-3
func LengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	bitMap := make([]int32, 8)
	var i, j, longest int
	for j < len(s) {
		// s[j] is repeated, move i
		if inBitMap(bitMap, s[j]) {
			if longest < j-i {
				longest = j - i
			}
			for i < j {
				if s[i] == s[j] {
					// move to next begin loc
					i++
					break
				}
				resetBitMap(bitMap, s[i])
				i++
			}
		}
		setBitMap(bitMap, s[j])
		j++
	}
	if longest < j-i {
		longest = j - i
	}
	return longest
}

func setBitMap(bm []int32, v uint8) {
	block, subIndex := v/32, uint(v%32)
	bm[block] |= int32(1 << subIndex)
}

func resetBitMap(bm []int32, v uint8) {
	block, subIndex := v/32, uint(v%32)
	bm[block] &= ^int32(1 << subIndex)
}

func inBitMap(bm []int32, v uint8) bool {
	block, subIndex := v/32, uint(v%32)
	return bm[block]&int32(1<<subIndex) != 0
}
