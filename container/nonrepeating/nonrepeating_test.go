package nonrepeating

import "testing"

func TestSubstr(t *testing.T)  {
	tests := []struct{
		s string
		ans int
	}{
		// normal cases
		{"abcabcbb", 3 },
		{"pwwkew", 3 },

		// edge cases
		{"", 0},
		{"b", 1},
		{"bbbbbb", 1},
		{"abcabcd", 4},

		// Chinese support
		{"一二三一二", 1},
	}

	for _, tt := range tests {
		actual := lengthOfNonRepeatingSubStr(tt.s)
		if actual != tt.ans {
			t.Errorf("got %d for input %s; "+"expected %d", actual, tt.s, tt.ans)
		}
	}

}


