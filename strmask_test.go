package strmask

import (
	"testing"
)

func TestApply(t *testing.T) {
	type args struct {
		mask string
		in   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				mask: "LLL-LLL-LLL",
				in:   "A23456ABC",
			},
			want: "A23-456-ABC",
		},
		{
			name: "2",
			args: args{
				mask: "0000-000-0",
				in:   "123456789",
			},
			want: "1234-567-8",
		},
		{
			name: "3",
			args: args{
				mask: "0000-000-0",
				in:   "9999-123-4",
			},
			want: "9999-123-4",
		},
		{
			name: "4",
			args: args{
				mask: "0000-000-0",
				in:   "9999-123-",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := Apply(tt.args.mask, tt.args.in); got != tt.want {
				if err != nil {
					t.Error(err)
				} else {
					t.Errorf("Apply() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
