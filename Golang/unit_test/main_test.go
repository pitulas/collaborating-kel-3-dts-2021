package unit_test

import "testing"

func Test_grade(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case return C",
			args: args{
				n: 30,
			},
			want: "C",
		},
		{
			name: "case return B",
			args: args{
				n: 60,
			},
			want: "D",
		},
		{
			name: "case return A",
			args: args{
				n: 80,
			},
			want: "A",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := status(tt.args.n); got != tt.want {
				t.Errorf("status() = %v, want %v", got, tt.want)
			}
		})
	}
}