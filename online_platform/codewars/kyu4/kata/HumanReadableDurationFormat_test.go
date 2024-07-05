package kata

import "testing"

func TestFormatDuration(t *testing.T) {
	type args struct {
		seconds int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"Example", args{1}, "1 second"},
		{"Example", args{62}, "1 minute and 2 seconds"},
		{"Example", args{120}, "2 minutes"},
		{"Example", args{3600}, "1 hour"},
		{"Example", args{3662}, "1 hour, 1 minute and 2 seconds"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatDuration(tt.args.seconds); got != tt.want {
				t.Errorf("FormatDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}
