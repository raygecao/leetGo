package dp

// LongestPalindrome find the longest palindrome substring in string s
// e.g. LongestPalindrome("babad") = "aba"
// ref link: https://leetcode-cn.com/problems/longest-palindromic-substring/
// leetcode-5
func LongestPalindrome(s string) string {
	// corner cases
	if len(s) <= 1 {
		return s
	}

	longest := -1
	// record the start index and end index
	var start, end int

	// init the dp map for every char and its neighbors
	dict := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dict[i] = make([]bool, len(s))
		dict[i][i] = true
		if i < len(s)-1 && s[i] == s[i+1] {
			dict[i][i+1] = true
			longest, start, end = 1, i, i+1
		}
	}

	for j := 0; j < len(s)-1; j++ {
		for i := 1; i <= j; i++ {
			// the core recursive formular
			dict[i-1][j+1] = dict[i][j] && s[i-1] == s[j+1]
			if dict[i-1][j+1] && j-i+2 > longest {
				longest, start, end = j-i+2, i-1, j+1
			}
		}
	}
	return s[start : end+1]
}
