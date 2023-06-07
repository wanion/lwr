package sysreport

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os/exec"
	"strings"
)

// getRpmPackages returns a list of installed RPMs retrieved using the `rpm` command.
func getRpmPackages() (packages []softwarePackage) {
	cmd := exec.Command("rpm", "-qa", "--qf", `%{NAME}\t%{VERSION}\t%{RELEASE}\t%{ARCH}\t%{INSTALLTIME:date}\t%{SUMMARY}\n`)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		fmt.Println("Error retrieving installed packages.")
		return packages
	}

	packages = parseRpmPackages(stdout)

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}

	return packages
}

func parseRpmPackages(h io.Reader) (packages []softwarePackage) {
	scanner := bufio.NewScanner(h)
	for scanner.Scan() {
		var pkg softwarePackage
		line := scanner.Text()
		if strings.Contains(line, "\t") {
			pkg = newSoftwarePackageFromRpm(line)
			packages = append(packages, pkg)
		}
	}

	return packages
}

// newSoftwarePackageFromRpm takes a line of correctly formatted output from `rpm` and turns
// it into a nice softwarePackage.
func newSoftwarePackageFromRpm(rpmString string) softwarePackage {
	var pkg softwarePackage
	p := strings.Split(rpmString, "\t")
	pkg.Name = p[0]
	pkg.Version = p[1]
	pkg.Release = p[2]
	pkg.Architecture = p[3]
	pkg.InstallDate = p[4]
	pkg.Description = p[5]
	return pkg
}
