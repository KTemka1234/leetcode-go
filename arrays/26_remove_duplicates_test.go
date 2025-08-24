package arrays

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		nums     []int
		want     int
		new_nums []int
	}{
		{
			nums:     []int{1, 1, 2},
			want:     2,
			new_nums: []int{1, 2},
		},
		{
			nums:     []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			want:     5,
			new_nums: []int{0, 1, 2, 3, 4},
		},
	}

	for _, tt := range tests {
		got := RemoveDuplicates(tt.nums)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("RemoveDuplicates(%v) = %d, want %d", tt.nums, got, tt.want)
		}

		if !reflect.DeepEqual(tt.nums[:got], tt.new_nums) {
			t.Errorf("RemoveDuplicates(%v) = %v, want %v", tt.nums, tt.nums[:got], tt.new_nums)
		}
	}
}
