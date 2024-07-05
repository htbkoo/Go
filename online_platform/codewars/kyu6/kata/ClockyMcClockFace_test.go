package kata

import "testing"

func TestWhatTimeIsIt(t *testing.T) {
	type args struct {
		angle int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"Example", args{0}, "12:00"},
		{"Example", args{90}, "03:00"},
		{"Example", args{180}, "06:00"},
		{"Example", args{270}, "09:00"},
		{"Example", args{360}, "12:00"},
		{"Example", args{30}, "01:00"},
		{"Example", args{40}, "01:20"},
		{"Example", args{45}, "01:30"},
		{"Example", args{50}, "01:40"},
		{"Example", args{60}, "02:00"},

		{"Actual Test Case", args{5}, "12:10"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WhatTimeIsIt(tt.args.angle); got != tt.want {
				t.Errorf("WhatTimeIsIt() = %v, want %v", got, tt.want)
			}
		})
	}
}
