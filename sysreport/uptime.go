package sysreport

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// getUptime uses `/proc/uptime` to return the current uptime.
func getUptime() (uptime int) {
	f, err := os.Open("/proc/uptime")
	if err == nil {
		uptime = parseUptime(f)
	}
	defer f.Close()
	return uptime
}

// parseUptime parses the uptime truncating to the nearest second.
func parseUptime(h io.Reader) int {
	b, err := ioutil.ReadAll(h)
	if err != nil {
		log.Fatal(err)
	}
	line := strings.Split(string(b), ".")
	uptime64, _ := strconv.ParseInt(line[0], 10, 32)
	return int(uptime64)
}
