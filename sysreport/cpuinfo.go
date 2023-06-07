package sysreport

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// getCPUInfo parses /proc/cpuinfo and returns the type and number of processors.
func getCPUInfo() (model string, pcount int) {
	file, err := os.Open("/proc/cpuinfo")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	cpuInfo := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if line := scanner.Text(); strings.Contains(line, ":") {
			kv := strings.Split(line, ":")
			k := strings.TrimSpace(kv[0])
			v := strings.TrimSpace(kv[1])
			cpuInfo[k] = v
		}
	}

	proc, err := strconv.ParseInt(cpuInfo["processor"], 10, 8)
	if err != nil {
		log.Println("Couldn't parse processor count from /proc/cpuinfo.")
	}

	model = cpuInfo["model name"]
	pcount = int(proc + 1)

	return model, pcount
}
