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
			name: "mask-# (any letter or digit)",
			args: args{
				mask:  "####-####-####-####",
				value: "qwer TYUI 0123-4567",
			},
			want: "qwer-TYUI-0123-4567",
		},
		{
			name: "mask-0 (only digit)",
			args: args{
				mask:  "+0 (000) 000-00-00",
				value: "A12345678900",
			},
			want: "+1 (234) 567-89-00",
		},
		{
			name: "mask-L (lower letter or digit)",
			args: args{
				mask:  "LLLL-LLLL",
				value: "QWER 1234",
			},
			want: "qwer-1234",
		},
		{
			name: "mask-l (lower letter)",
			args: args{
				mask:  "llll-llll",
				value: "QWER 1234 zxCV",
			},
			want: "qwer-zxcv",
		},
		{
			name: "mask-U (upper letter or digit)",
			args: args{
				mask:  "UUUU-UUUU",
				value: "qwer 1234",
			},
			want: "QWER-1234",
		},
		{
			name: "mask-u (upper letter)",
			args: args{
				mask:  "uuuu-uuuu",
				value: "qwer 1234 ASDV",
			},
			want: "QWER-ASDV",
		},
		{
			name: "error mask",
			args: args{
				mask:  "uuuu-uuuu",
				value: "123345",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "error mask with value",
			args: args{
				mask:  "uuuu-uuuu",
				value: "qwer",
			},
			want:    "QWER-",
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
