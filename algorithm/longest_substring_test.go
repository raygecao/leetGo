package algorithm

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		s string
		v int
	}{
		{"", 0},
		{"a", 1},
		{"aab", 2},
		{"dvdf", 3},
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
	}

	for i, tt := range tests {
		actual := LengthOfLongestSubstring(tt.s)
		require.Equalf(t, actual, tt.v, "The %d-th case: LengthOfLongestSubstring(%s), got %d, expected %d",
			i+1, tt.s, actual, tt.v)
	}
}
