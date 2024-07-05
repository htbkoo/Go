package kata

import "testing"

func TestLargestDifference(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"Example", args{[]int{9, 4, 1, 10, 3, 4, 0, -1, -2}}, 4},
		{"Example", args{[]int{3, 2, 1}}, 0},
		{"Example", args{[]int{1, 2, 3}}, 2},
		{"Example", args{[]int{9, 4, 1, 2, 3, 0, -1, -2}}, 2},
		{"Example", args{[]int{9, 4, 1, 2, 3, 4, 0, -1, -2}}, 4},
		{"Example", args{[]int{78, 88, 64, 94, 17, 91, 57, 69, 38, 62, 13, 17, 35, 15, 20}}, 10},
		{"Example", args{[]int{4, 3, 3, 1, 5, 2}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LargestDifference(tt.args.data); got != tt.want {
				t.Errorf("LargestDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
