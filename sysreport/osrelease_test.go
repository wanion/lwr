package sysreport

import (
	"testing"
)

func Test_getRedHatRelease(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name        string
		args        args
		wantRelease string
	}{
		{
			name:        "Parse RHEL 8.4 redhat-release file",
			args:        args{"testdata/redhat-release-84"},
			wantRelease: "Red Hat Enterprise Linux release 8.4 (Ootpa)",
		}, {
			name:        "Parse RHEL 8.4 redhat-release file",
			args:        args{"testdata/redhat-release-79"},
			wantRelease: "Red Hat Enterprise Linux Server release 7.9 (Maipo)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRelease := getRedHatRelease(tt.args.path); gotRelease != tt.wantRelease {
				t.Errorf("getRedHatRelease() = %v, want %v", gotRelease, tt.wantRelease)
			}
		})
	}
}

func Test_getOSRelease(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name        string
		args        args
		wantRelease string
	}{
		{
			name:        "Parse Ubuntu 20.04 os-release file",
			args:        args{"testdata/os-release-ubuntu2004"},
			wantRelease: "Ubuntu 20.04.2 LTS",
		}, {
			name:        "Parse Ubuntu 20.04 os-release with no PRETTY_NAME",
			args:        args{"testdata/os-release-nopretty"},
			wantRelease: "Ubuntu 20.04.2 LTS (Focal Fossa)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRelease := getOSRelease(tt.args.path); gotRelease != tt.wantRelease {
				t.Errorf("getOSRelease() = %v, want %v", gotRelease, tt.wantRelease)
			}
		})
	}
}
