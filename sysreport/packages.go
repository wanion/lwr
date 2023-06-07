package sysreport

import (
	"os"
)

// getPackages guesses whether a system uses rpm or dpkg and returns a list of installed packages.
func getPackages() (packages []softwarePackage) {
	if _, err := os.Stat("/var/lib/dpkg/status"); err == nil {
		packages = getDpkgPackages("/var/lib/dpkg/status", "/var/log/dpkg.log")
	} else {
		packages = getRpmPackages()
	}
	return packages
}
