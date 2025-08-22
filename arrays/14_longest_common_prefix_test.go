package arrays

import (
	"reflect"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		strs []string
		want string
	}{
		{
			strs: []string{"flower", "flow", "flight"},
			want: "fl",
		},
		{
			strs: []string{"dog", "racecar", "car"},
			want: "",
		},
		{
			strs: []string{"", "flow", "flight"},
			want: "",
		},
		{
			strs: []string{"flower", "flower", "flower"},
			want: "flower",
		},
		{
			strs: []string{"dog", "do", "d"},
			want: "d",
		},
		{
			strs: []string{"flower"},
			want: "flower",
		},
		{
			strs: []string{},
			want: "",
		},
	}

	for _, tt := range tests {
		got := LongestCommonPrefix(tt.strs)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("LongestCommonPrefix(%v) = %v, want %v", tt.strs, got, tt.want)
		}
	}
}
