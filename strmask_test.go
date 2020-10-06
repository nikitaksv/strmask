package strmask

import (
	"testing"
)

func TestApply(t *testing.T) {
	type args struct {
		mask  string
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				mask:  "LLL-LLL-LLL",
				value: "A23456ABC",
			},
			want: "A23-456-ABC",
		},
		{
			name: "success",
			args: args{
				mask:  "0000-000-0",
				value: "123456789",
			},
			want: "1234-567-8",
		},
		{
			name: "success",
			args: args{
				mask:  "0000-000-0",
				value: "9999-123-4",
			},
			want: "9999-123-4",
		},
		{
			name: "error",
			args: args{
				mask:  "0000-000-0",
				value: "9999-123-",
			},
			wantErr: true,
		},
		{
			name: "error",
			args: args{
				mask:  "0000-000-0",
				value: "asdasdasdas",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Apply(tt.args.mask, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Apply() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Apply() got = %v, want %v", got, tt.want)
			}
		})
	}
}
