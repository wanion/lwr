package sysreport

import (
	"io"
	"strings"
	"testing"
)

func Test_parseUptime(t *testing.T) {
	type args struct {
		h io.Reader
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Uptime of 70831 seconds (almost 20 hours)",
			args: args{
				strings.NewReader("70831.65 141450.77\n"),
			},
			want: 70831,
		},
		{
			name: "Uptime of 1218771 seconds (approximately 14 days)",
			args: args{
				strings.NewReader("1218771.65 141450.77"),
			},
			want: 1218771,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseUptime(tt.args.h); got != tt.want {
				t.Errorf("parseUptime() = %v, want %v", got, tt.want)
			}
		})
	}
}
