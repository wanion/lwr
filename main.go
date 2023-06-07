package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/wanion/lwr/sysreport"
)

var (
	appName           = "lwr"
	appVersion        = "21.05.27"
	verbose           = false
	configurationPath = "/etc/lwr.conf"
)

type configuration struct {
	Server       string
	Port         int
	AssetID      string
	AgentVersion string
}

func main() {
	conf := getConfiguration(configurationPath)
	if conf.AssetID == "" {
		conf.AssetID = generateAssetID()
	}

	configureCmd := flag.NewFlagSet("configure", flag.ExitOnError)
	configureCmd.StringVar(&conf.Server, "server", "lansweeper.example.org", "specify the server to send reports")
	configureCmd.IntVar(&conf.Port, "port", 9524, "port send report on")
	configureCmd.StringVar(&conf.AgentVersion, "agentversion", "7.2.110.19", "agent version to report")

	reportCmd := flag.NewFlagSet("report", flag.ExitOnError)
	reportCmd.BoolVar(&verbose, "verbose", false, "verbose output")
	reportCmd.StringVar(&conf.Server, "server", "lansweeper.example.org", "specify the server to send reports")
	reportCmd.IntVar(&conf.Port, "port", 9524, "port send report on")
	reportCmd.StringVar(&conf.AgentVersion, "agentversion", "7.2.110.19", "agent version to report")

	var subcommand string
	switch l := len(os.Args); {
	case l > 1 && os.Args[1][0] == '-':
		fmt.Printf("Please call %s with `report` subcommand if you want to pass options for a scan.\nUse `help` for help.\n", os.Args[0])
		os.Exit(1)
	case l == 1:
		subcommand = "report"
	case l >= 2:
		subcommand = os.Args[1]
	default:
		log.Fatalln("Error processing commandline.")
	}

	switch subcommand {
	case "configure":
		if len(os.Args) > 2 {
			configureCmd.Parse(os.Args[2:])
		}
		fmt.Print(configure(conf))
		os.Exit(0)

	case "version":
		fmt.Println(appName, "version", appVersion)
		os.Exit(0)

	case "report":
		if len(os.Args) >= 2 {
			reportCmd.Parse(os.Args[2:])
		}
		cmdReport(conf)

	case "print", "dump":
		scan := sysreport.NewReport(conf.AgentVersion)
		fmt.Println(string(scan.GetJSON()))

	case "help":
		printHelp()

	default:
		fmt.Printf("unrecognised subcommand %s\n", subcommand)
		printHelp()
		os.Exit(1)
	}
}

func printVerbose(str string) {
	if verbose {
		fmt.Println(str)
	}
}
