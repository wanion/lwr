package sysreport

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

// getDpkgPackages retrieves a list of installed packages from `/var/lib/dpkg/status` and then
// attempts to determine their install dates from `/var/log/dpkg.log`.
func getDpkgPackages(statusPath, logPath string) (packages []softwarePackage) {
	// get installed packages from /var/lib/dpkg/status
	f, err := os.Open(statusPath)
	if err == nil {
		packages = processDpkgStatus(f)
	}
	defer f.Close()

	// try to get install timestamp from /var/log/dpkg.log
	f, err = os.Open(logPath)
	if err == nil {
		packages = getDpkgInstallTime(f, packages)
	}
	defer f.Close()

	return packages
}

// processDpkgStatus determines which packages currently have an installed status according to
// the contents of dpkg/status.
func processDpkgStatus(h io.Reader) (packages []softwarePackage) {
	scanner := bufio.NewScanner(h)
	var pkg softwarePackage
	for scanner.Scan() {
		if line := scanner.Text(); strings.Contains(line, ": ") {
			kv := strings.SplitN(line, ": ", 2)
			k := kv[0]
			v := kv[1]
			switch k {
			case "Package":
				if pkg.Name != "" {
					packages = append(packages, pkg)
				}
				pkg = softwarePackage{}
				pkg.Name = v
			case "Version":
				pkg.Version = v
			case "Architecture":
				pkg.Architecture = v
			case "Status":
				// discard packages that aren't installed
				if v != "install ok installed" {
					pkg.Name = ""
				}
			}
		}
	}

	return packages
}

// getDpkgInstallTime inspects dpkg.log to find the original installation timestamp for all
// packages and then populates those values in the list of supplied packages.
func getDpkgInstallTime(h io.Reader, packages []softwarePackage) []softwarePackage {
	scanner := bufio.NewScanner(h)
	installs := make(map[string]time.Time)
	for scanner.Scan() {
		if line := scanner.Text(); strings.Contains(line, " install ") {
			parts := strings.SplitN(line, " ", 6)
			idx := strings.LastIndex(parts[3], ":")
			name := string(parts[3][:idx])
			ts, err := time.Parse("2006-01-02 15:04:05", parts[0]+" "+parts[1])
			if err != nil {
				log.Println(err)
			}
			if _, ok := installs[name]; !ok {
				installs[name] = ts
			}
		}
	}

	loc, _ := time.LoadLocation("Pacific/Auckland")

	for i := range packages {
		if v, ok := installs[packages[i].Name]; ok {
			packages[i].InstallDate = v.In(loc).Format("Mon 2 Jan 2006 15:04:05 MST")
		}
	}
	return packages
}
