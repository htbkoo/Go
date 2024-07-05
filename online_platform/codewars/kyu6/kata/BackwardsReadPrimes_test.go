package kata

import (
	"reflect"
	"testing"
)

func TestBackwardsPrime(t *testing.T) {
	type args struct {
		start int
		stop  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"Example", args{1, 100}, []int{13, 17, 31, 37, 71, 73, 79, 97}},
		{"Example", args{1, 31}, []int{13, 17, 31}},
		{"Example", args{501, 599}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BackwardsPrime(tt.args.start, tt.args.stop); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BackwardsPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}
