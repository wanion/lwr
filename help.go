package main

import (
	"fmt"
	"os"
	"path"
)

var subcommands = []struct{ Command, HelpText string }{
	{"configure [-h]", "display current or updated configuration"},
	{"help", "display this help"},
	{"print", "generate and print report to stdout (does not send)"},
	{"report [-h]", "generate and send a report to LanSweeper"},
	{"version", "display version"},
}

func printHelp() {
	appname := path.Base(os.Args[0])
	fmt.Printf("Subcommands for %s:\n", appname)
	for _, v := range subcommands {
		fmt.Printf("  %s\n  \t%s\n", v.Command, v.HelpText)
	}
	fmt.Println("Subcommands with a -h option can be run with -h to see what other flags they offer.")

	fmt.Printf("\nTo set a persistent asset identifier for this machine:\n  %s configure > /etc/lwr.conf\n", appname)
	fmt.Printf("\nTo modify configuration:\n  %s configure [-server <hostname>] [-port <port>] [-agentversion <version>] > /etc/lwr.conf\n", appname)
}
