package main

import (
	"fmt"

	"github.com/wanion/lwr/sysreport"
)

func cmdReport(conf configuration) {
	lSServer := fmt.Sprintf("https://%s:%d/lsagent", conf.Server, conf.Port)

	printVerbose("Preparing report...")
	scan := sysreport.NewReport(conf.AgentVersion)
	printVerbose("Report assembled.")

	printVerbose("Sending report...")
	err := scan.SendReport(lSServer, conf.AssetID)
	if err == nil {
		fmt.Println("Report successfully submitted.")
	}
}
