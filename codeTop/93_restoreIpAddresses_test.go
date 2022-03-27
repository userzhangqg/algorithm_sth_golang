package codeTop

import (
	"reflect"
	"testing"
)

func Test_restoreIpAddresses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		args args
		want []string
	}{
		{
			args{
				"25525511135",
			},
			[]string{"255.255.11.135", "255.255.111.35"},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := restoreIpAddresses(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("restoreIpAddresses() = %v, want %v", got, tt.want)
			}
		})
	}
}
