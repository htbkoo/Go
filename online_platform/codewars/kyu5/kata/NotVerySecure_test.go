package kata

// TODO: replace with your own tests (TDD). An example to get you started is included below.
// Ginkgo BDD Testing Framework <http://onsi.github.io/ginkgo></http:>
// Gomega Matcher Library <http://onsi.github.io/gomega></http:>

import (
	"testing"
)

func Test_alphanumeric(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"Example", args{".*?"}, false},
		{"Example", args{"a"}, true},
		{"Example", args{"Mazinkaiser"}, true},
		{"Example", args{"hello world_"}, false},
		{"Example", args{"PassW0rd"}, true},
		{"Example", args{"     "}, false},
		{"Example", args{""}, false},
		{"Example", args{"\n\t\n"}, false},
		{"Example", args{"ciao\n$$_"}, false},
		{"Example", args{"__ * __"}, false},
		{"Example", args{"&)))((("}, false},
		{"Example", args{"43534h56jmTHHF3k"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := alphanumeric(tt.args.str); got != tt.want {
				t.Errorf("alphanumeric() = %v, want %v", got, tt.want)
			}
		})
	}
}
