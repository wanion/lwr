package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func configure(conf configuration) string {
	var config string
	if conf.Server != "" {
		config += fmt.Sprintf("Server=%s\n", conf.Server)
	}
	if conf.Port != 0 {
		config += fmt.Sprintf("Port=%d\n", conf.Port)
	}
	if conf.AssetID == "" {
		conf.AssetID = generateAssetID()
	}
	config += fmt.Sprintf("AssetId=%s\n", conf.AssetID)
	if conf.AgentVersion != "" {
		config += fmt.Sprintf("Version=%s\n", conf.AgentVersion)
	}
	return config
}

func generateAssetID() string {
	id := uuid.NewString()
	return id
}

func getConfiguration(configurationPath string) (conf configuration) {
	h, err := os.Open(configurationPath)
	defer h.Close()
	if err != nil {
		return conf
	}
	return loadConfiguration(h)
}

func loadConfiguration(h io.Reader) (conf configuration) {
	scanner := bufio.NewScanner(h)
	for scanner.Scan() {
		if line := scanner.Text(); strings.Contains(line, "=") {
			kv := strings.SplitN(line, "=", 2)
			k := kv[0]
			v := kv[1]
			switch k {
			case "Server":
				conf.Server = v
			case "Port":
				port, err := strconv.ParseInt(v, 10, 32)
				if err != nil {
					log.Println("couldn't parse port in configuration")
					log.Fatalln(err)
				}
				conf.Port = int(port)
			case "AssetId":
				conf.AssetID = v
			case "Version":
				conf.AgentVersion = v
			default:
				log.Println("Unknown setting", k, "with value", v, "ignored.")
			}
		}
	}

	return conf
}
