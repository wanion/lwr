package main

import (
	"io"
	"reflect"
	"regexp"
	"strings"
	"testing"
)

func Test_generateAssetID(t *testing.T) {
	re := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
	got := generateAssetID()
	if re.FindString(got) == "" {
		t.Errorf("generateAssetID() = %v, want %v", got, re.String())
	}
}

func Test_configure(t *testing.T) {
	type args struct {
		conf configuration
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Full configuration",
			args: args{
				configuration{
					Server:       "abc",
					Port:         123,
					AssetID:      "xxx",
					AgentVersion: "7.2.110.19",
				},
			},
			want: "Server=abc\nPort=123\nAssetId=xxx\nVersion=7.2.110.19\n",
		},
		{
			name: "Configuration without server",
			args: args{
				configuration{
					Port:         123,
					AssetID:      "xxx",
					AgentVersion: "7.2.110.19",
				},
			},
			want: "Port=123\nAssetId=xxx\nVersion=7.2.110.19\n",
		},
		{
			name: "Configuration without port",
			args: args{
				configuration{
					Server:       "abc",
					AssetID:      "xxx",
					AgentVersion: "7.2.110.19",
				},
			},
			want: "Server=abc\nAssetId=xxx\nVersion=7.2.110.19\n",
		},
		{
			name: "Configuration without port or server",
			args: args{
				configuration{
					AssetID:      "xxx",
					AgentVersion: "7.2.110.19",
				},
			},
			want: "AssetId=xxx\nVersion=7.2.110.19\n",
		},
		{
			name: "Configuration without agent version",
			args: args{
				configuration{
					Server:  "abc",
					Port:    123,
					AssetID: "xxx",
				},
			},
			want: "Server=abc\nPort=123\nAssetId=xxx\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := configure(tt.args.conf); got != tt.want {
				t.Errorf("configure() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loadConfiguration(t *testing.T) {
	type args struct {
		h io.Reader
	}
	tests := []struct {
		name     string
		args     args
		wantConf configuration
	}{
		{
			name: "Complete configuration file",
			args: args{strings.NewReader("Server=server\nPort=1234\nAssetId=f0caa043-df30-4b3d-9451-e03dc289f866\nVersion=1.2.3")},
			wantConf: configuration{
				Server:       "server",
				Port:         1234,
				AssetID:      "f0caa043-df30-4b3d-9451-e03dc289f866",
				AgentVersion: "1.2.3",
			},
		},
		{
			name: "AssetID only in configuration file",
			args: args{strings.NewReader("AssetId=f0caa043-df30-4b3d-9451-e03dc289f866")},
			wantConf: configuration{
				AssetID: "f0caa043-df30-4b3d-9451-e03dc289f866",
			},
		},
		{
			name:     "Configuration file with no recognised settings",
			args:     args{strings.NewReader("a\nb\nc=d\ne:f")},
			wantConf: configuration{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotConf := loadConfiguration(tt.args.h); !reflect.DeepEqual(gotConf, tt.wantConf) {
				t.Errorf("loadConfiguration() = %v, want %v", gotConf, tt.wantConf)
			}
		})
	}
}

func Test_getConfiguration(t *testing.T) {
	type args struct {
		configurationPath string
	}
	tests := []struct {
		name     string
		args     args
		wantConf configuration
	}{
		{
			name: "Parse golden configuration file",
			args: args{"testdata/lwr.conf"},
			wantConf: configuration{
				Server:       "server.example.com",
				Port:         321,
				AssetID:      "640f05b6-9ef4-45f8-b376-ab189ed48082",
				AgentVersion: "1.2.3",
			},
		},
		{
			name:     "Attempt to read missing configuration file",
			args:     args{"testdata/missing.conf"},
			wantConf: configuration{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotConf := getConfiguration(tt.args.configurationPath); !reflect.DeepEqual(gotConf, tt.wantConf) {
				t.Errorf("getConfiguration() = %v, want %v", gotConf, tt.wantConf)
			}
		})
	}
}
