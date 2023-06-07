package sysreport

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

// getMemInfo parses the output of /proc/meminfo and turns it into a map.
func getMemInfo(path string) map[string]string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	return parseMemInfo(file)
}

// parseMemInfo extracts key/value pairs from meminfo.
func parseMemInfo(h io.Reader) map[string]string {
	memInfo := make(map[string]string)
	scanner := bufio.NewScanner(h)
	for scanner.Scan() {
		if line := scanner.Text(); strings.Contains(line, ":") {
			kv := strings.Split(line, ":")
			k := strings.TrimSpace(kv[0])
			v := strings.TrimSpace(kv[1])
			memInfo[k] = v
		}
	}

	return memInfo
}

// getTotalMemory returns just the total memory from getMemInfo().
func getTotalMemory() string {
	memInfo := getMemInfo("/proc/meminfo")
	if _, ok := memInfo["MemTotal"]; !ok {
		log.Fatal("Couldn't determine total memory.")
	}
	return memInfo["MemTotal"]
}
