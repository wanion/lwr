package sysreport

import (
	"reflect"
	"testing"
)

func Test_getMemInfo(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{
			name: "Test parsing /proc/meminfo from Red Hat 8.4 (kernel 4.18.0)",
			args: args{"testdata/meminfo-4.18"},
			want: map[string]string{
				"MemTotal":          "4022908 kB",
				"MemFree":           "582972 kB",
				"MemAvailable":      "1231396 kB",
				"Buffers":           "140 kB",
				"Cached":            "679780 kB",
				"SwapCached":        "108132 kB",
				"Active":            "889336 kB",
				"Inactive":          "1103660 kB",
				"Active(anon)":      "398044 kB",
				"Inactive(anon)":    "918292 kB",
				"Active(file)":      "491292 kB",
				"Inactive(file)":    "185368 kB",
				"Unevictable":       "24012 kB",
				"Mlocked":           "24012 kB",
				"SwapTotal":         "2097148 kB",
				"SwapFree":          "1442276 kB",
				"Dirty":             "16 kB",
				"Writeback":         "0 kB",
				"AnonPages":         "1322404 kB",
				"Mapped":            "121968 kB",
				"Shmem":             "1208 kB",
				"KReclaimable":      "74252 kB",
				"Slab":              "176968 kB",
				"SReclaimable":      "74252 kB",
				"SUnreclaim":        "102716 kB",
				"KernelStack":       "14364 kB",
				"PageTables":        "78576 kB",
				"NFS_Unstable":      "0 kB",
				"Bounce":            "0 kB",
				"WritebackTmp":      "0 kB",
				"CommitLimit":       "3584312 kB",
				"Committed_AS":      "6060700 kB",
				"VmallocTotal":      "34359738367 kB",
				"VmallocUsed":       "0 kB",
				"VmallocChunk":      "0 kB",
				"Percpu":            "2072 kB",
				"HardwareCorrupted": "0 kB",
				"AnonHugePages":     "0 kB",
				"ShmemHugePages":    "0 kB",
				"ShmemPmdMapped":    "0 kB",
				"FileHugePages":     "0 kB",
				"FilePmdMapped":     "0 kB",
				"HugePages_Total":   "1",
				"HugePages_Free":    "1",
				"HugePages_Rsvd":    "0",
				"HugePages_Surp":    "0",
				"Hugepagesize":      "1048576 kB",
				"Hugetlb":           "1048576 kB",
				"DirectMap4k":       "439664 kB",
				"DirectMap2M":       "2705408 kB",
				"DirectMap1G":       "3145728 kB",
			},
		}, {
			name: "Test parsing /proc/meminfo from Ubuntu 20.04 (kernel 5.4.0)",
			args: args{"testdata/meminfo-5.4"},
			want: map[string]string{
				"MemTotal":          "4030564",
				"MemFree":           "363740",
				"MemAvailable":      "3514636",
				"Buffers":           "223056",
				"Cached":            "2999396",
				"SwapCached":        "52",
				"Active":            "1489532",
				"Inactive":          "1792728",
				"Active(anon)":      "29768",
				"Inactive(anon)":    "39448",
				"Active(file)":      "1459764",
				"Inactive(file)":    "1753280",
				"Unevictable":       "18788",
				"Mlocked":           "18788",
				"SwapTotal":         "4030460",
				"SwapFree":          "4029936",
				"Dirty":             "0",
				"Writeback":         "0",
				"AnonPages":         "78480",
				"Mapped":            "166932",
				"Shmem":             "1120",
				"KReclaimable":      "227288",
				"Slab":              "298052",
				"SReclaimable":      "227288",
				"SUnreclaim":        "70764",
				"KernelStack":       "3808",
				"PageTables":        "2116",
				"NFS_Unstable":      "0",
				"Bounce":            "0",
				"WritebackTmp":      "0",
				"CommitLimit":       "6045740",
				"Committed_AS":      "402244",
				"VmallocTotal":      "34359738367",
				"VmallocUsed":       "24988",
				"VmallocChunk":      "0",
				"Percpu":            "1648",
				"HardwareCorrupted": "0",
				"AnonHugePages":     "0",
				"ShmemHugePages":    "0",
				"ShmemPmdMapped":    "0",
				"FileHugePages":     "0",
				"FilePmdMapped":     "0",
				"CmaTotal":          "0",
				"CmaFree":           "0",
				"HugePages_Total":   "0",
				"HugePages_Free":    "0",
				"HugePages_Rsvd":    "0",
				"HugePages_Surp":    "0",
				"Hugepagesize":      "2048",
				"Hugetlb":           "0",
				"DirectMap4k":       "182144",
				"DirectMap2M":       "4012032",
				"DirectMap1G":       "2097152",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMemInfo(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getMemInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
