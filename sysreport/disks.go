package sysreport

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

// getDisks checks what volumes `df -k` returns.
func getDisks() (disks []hardDisk) {
	cmd := exec.Command("df", "-k")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		fmt.Println("Error retrieving disk usage.")
		return disks
	}
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		var disk hardDisk
		line := scanner.Text()
		if strings.Contains(line, "/") {
			d := strings.Fields(line)
			disk.Filesystem = d[0]
			disk.Size = d[1]
			disk.Used = d[2]
			disk.Available = d[3]
			disk.Percentage = d[4]
			disk.MountedOn = d[5]
			disks = append(disks, disk)
		}
	}
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
	return disks
}
