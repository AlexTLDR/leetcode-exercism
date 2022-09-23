package main

import "testing"

func Test_maximumScore(t *testing.T) {
	tests := []struct {
		name        string
		want        int
		nums        []int
		multipliers []int
	}{
		// TODO: Add test cases.
		{
			name:        "happyPath",
			want:        14,
			nums:        []int{1, 2, 3},
			multipliers: []int{3, 2, 1},
		},

		{
			name:        "1Path",
			want:        1,
			nums:        []int{1},
			multipliers: []int{1},
		},
		{
			name:        "2NumsPath",
			want:        -6,
			nums:        []int{10, 2},
			multipliers: []int{-3},
		}, {
			name:        "leetcodeExamplePath",
			want:        102,
			nums:        []int{-5, -3, -3, -2, 7, 1},
			multipliers: []int{-10, -5, 3, 4, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumScore(tt.nums, tt.multipliers); got != tt.want {
				t.Errorf("maximumScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
