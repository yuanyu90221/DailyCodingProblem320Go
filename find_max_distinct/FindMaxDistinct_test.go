package find_max_distinct

import "testing"

func TestFindMaxDistinct(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "jiujitsu",
			args: args{input: "jiujitsu"},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMaxDistinct(tt.args.input); got != tt.want {
				t.Errorf("FindMaxDistinct() = %v, want %v", got, tt.want)
			}
		})
	}
}
