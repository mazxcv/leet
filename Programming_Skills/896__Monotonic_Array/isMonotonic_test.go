package _96__Monotonic_Array

import "testing"

func Test_isMonotonic(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "{1, 2, 3}",
			args: args{nums: []int{1, 2, 3}},
			want: true,
		},
		{
			name: "{1,2,2,3}",
			args: args{nums: []int{1, 2, 2, 3}},
			want: true,
		},
		{
			name: "{6,5,4,4}",
			args: args{nums: []int{6, 5, 4, 4}},
			want: true,
		},
		{
			name: "{1,3,2}",
			args: args{nums: []int{1, 3, 2}},
			want: false,
		},
		{
			name: "{3,1,9}",
			args: args{nums: []int{3, 1, 9}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMonotonic(tt.args.nums); got != tt.want {
				t.Errorf("isMonotonic() = %v, want %v", got, tt.want)
			}
		})
	}
}
