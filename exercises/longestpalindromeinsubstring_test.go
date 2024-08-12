package exercises

import "testing"

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"racecar", "racecar"},
		{"babad", "bab"},
		{"cbbd", "bb"},
		{"a", "a"},
		{"ac", "a"},
		{"bb", "bb"},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := LongestPalindrome(tt.s)
			if got != tt.want {
				t.Errorf("longestPalindrome((%v) = %v; want %v", tt.s, got, tt.want)
			}
		})
	}
}
