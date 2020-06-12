package dp

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		s       string
		results []string
	}{
		{"", []string{""}},
		{"a", []string{"a"}},
		{"aa", []string{"aa"}},
		{"ab", []string{"a", "b"}},
		{"ccd", []string{"cc"}},
		{"adam", []string{"ada"}},
		{"babad", []string{"bab", "aba"}},
		{"cbbd", []string{"bb"}},
		{"abccba", []string{"abccba"}},
		{"aabccbab", []string{"abccba"}},
		{"aabccba", []string{"abccba"}},
	}

	for i, tt := range tests {
		actual := LongestPalindrome(tt.s)
		var has bool
		for _, i := range tt.results {
			if actual == i {
				has = true
				break
			}
		}
		if !has {
			require.Fail(t, "", "the %d-th case: LongestPalindrome(%s), gets %s, expects in %v", i+1, tt.s, actual, tt.results)
		}
	}
}
