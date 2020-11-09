package utils

import "testing"

func TestTimeStamp(t *testing.T) {
	type args struct {
		val string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{"test1", args{"2020-02-11T14:06:42.520978+0800"}, 1581401202},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimeStamp(tt.args.val); got != tt.want {
				t.Errorf("TimeStamp() = %v, want %v", got, tt.want)
			}
		})
	}
}
