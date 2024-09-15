package main

import (
	"fmt"
	"os"

	"paepcke.de/opnborg"
)

const (
	_app     = "[OPNBORG-CLI]"
	_version = "[v0.0.1]"
)

func main() {

	fmt.Println(_app + "[STARTUP]" + _version)

	// Read Application Env
	fmt.Println(_app + "[STARTING][READ-CONFIG-FROM-ENV]")
	config, err := opnborg.ReadConfig()
	if err != nil {
		fmt.Printf(_app+"[EXIT]%s\n", err)
		os.Exit(1)
	}
	config.AppName = _app
	config.Log = false
	fmt.Println(_app + "[SUCCESS][READ-CONFIG-FROM-ENV]")

	// Perform Backup of all Appliances xml configuration
	err = opnborg.Backup(config)
	if err != nil {
		fmt.Printf(_app+"[EXIT]%s\n", err)
		os.Exit(1)
	}
	fmt.Println(_app + "[END]")
}
