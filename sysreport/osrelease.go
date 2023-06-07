package sysreport

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// getLinuxRelease checks what distribution OS release is installed.
func getLinuxRelease() (release string) {
	release = getRedHatRelease("/etc/redhat-release")
	if release != "" {
		return release
	}

	return getOSRelease("/etc/os-release")
}

// getRedHatRelease tries to get the OS release from /etc/redhat-release.
func getRedHatRelease(path string) (release string) {
	content, err := ioutil.ReadFile(path)
	if err == nil {
		release = strings.TrimSpace(string(content))
	}

	return release
}

// getOSRelease gets the release from /etc/os-release.
func getOSRelease(path string) (release string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	etcOsRelease := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if line := scanner.Text(); strings.Contains(line, "=") {
			kv := strings.Split(line, "=")
			k := kv[0]
			v := strings.Trim(kv[1], "\"")
			etcOsRelease[k] = v
		}
	}

	if prettyName, ok := etcOsRelease["PRETTY_NAME"]; ok {
		release = prettyName
	} else {
		name, nok := etcOsRelease["NAME"]
		version, vok := etcOsRelease["VERSION"]
		if vok && nok {
			release = name + " " + version
		} else if nok {
			release = name
		} else if vok {
			release = version
		}
	}

	return release
}
