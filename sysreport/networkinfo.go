package sysreport

import (
	"net"
	"strings"
)

// getNetworkInfo lists all the network intefaces and also selects the one true IP address
// for the IPAddress field of the report.
// The one true IP address will be the first non-loopback IPv4 address. If no addresses match
// criteria it will be the last IPv6 address.
func getNetworkInfo() (chosenIP string, macAddresses []string, interfaces []networkInterface) {
	var pickedIP bool

	ifaces, _ := net.Interfaces()
	for _, iface := range ifaces {
		var netface networkInterface
		netface.Name = iface.Name
		if iface.HardwareAddr != nil {
			macAddress := iface.HardwareAddr.String()
			macAddresses = append(macAddresses, macAddress)
			netface.Mac = &macAddress
		}

		addrs, _ := iface.Addrs()
		for _, addr := range addrs {
			ip := addr.(*net.IPNet).IP
			if !pickedIP {
				chosenIP = ip.String()
				if !ip.IsLoopback() && (ip.To4() != nil) {
					pickedIP = true
				}
			}

			if strings.Count(ip.String(), ":") < 2 {
				netface.IP4 = ip.String()
			} else {
				netface.IP6 = ip.String()
			}
		}
		interfaces = append(interfaces, netface)
	}
	return chosenIP, macAddresses, interfaces
}
