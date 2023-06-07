/*
Package sysreport implements routines to inspect the system and produce
a report compatible with LanSweeper Agent.

It does not require elevated privileges for the information currently reported.

It runs the following commands:
	rpm -qa --qf "%{NAME}\t%{VERSION}\t%{RELEASE}\t%{ARCH}\t%{INSTALLTIME:date}\t%{SUMMARY}\n"
	df -k

Implemented:
	Hard disks (mounts)
	Packages
	SystemInfo

Not implemented:
	SMBIOS information
	Devices
		Graphics cards
		PCI cards
		Optical drives

Blank in captured output from official Linux LSAgent:
	Volumes
	Sound cards
	SystemInfo — PCManufacturer, SystemSku, Firmware

This package does not faithfully replicate bugs where they are not required for the report
to be processed correctly.

Bugs fixed:
	Erroneous "split" entry in SmBios.
	Linebreak ("\n") after hostname in AssetName
	Duplicated/null MAC addresses

Consider:
	https://www.dmtf.org/sites/default/files/standards/documents/DSP0134_3.1.1.pdf
	https://github.com/digitalocean/go-smbios

*/
package sysreport
