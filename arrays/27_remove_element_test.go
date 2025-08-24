package arrays

import (
	"reflect"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		nums   []int
		val    int
		want   int
		new_nums []int
	}{
		{
			nums:   []int{3, 2, 2, 3},
			val:    3,
			want:   2,
			new_nums: []int{2, 2},
		},
		{
			nums:   []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:    2,
			want:   5,
			new_nums: []int{0, 1, 3, 0, 4},
		},
	}

	for _, tt := range tests {
		got := RemoveElement(tt.nums, tt.val)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("RemoveElement(%v, %d) = %d, want %d", tt.nums, tt.val, got, tt.want)
		}

		if !reflect.DeepEqual(tt.nums[:got], tt.new_nums) {
			t.Errorf("RemoveElement(%v, %d) = %v, want %v", tt.nums, tt.val, tt.nums[:got], tt.new_nums)
		}
	}
}
