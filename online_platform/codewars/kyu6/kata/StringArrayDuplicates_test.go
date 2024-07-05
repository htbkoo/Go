package kata

import (
	"reflect"
	"testing"
)

func TestDup(t *testing.T) {
	type args struct {
		arr []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{"Example", args{[]string{"ccooddddddewwwaaaaarrrrsssss", "piccaninny", "hubbubbubboo"}}, []string{"codewars", "picaniny", "hubububo"}},
		{"Example", args{[]string{"abracadabra", "allottee", "assessee"}}, []string{"abracadabra", "alote", "asese"}},
		{"Example", args{[]string{"kelless", "keenness"}}, []string{"keles", "kenes"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Dup(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Dup() = %v, want %v", got, tt.want)
			}
		})
	}
}
