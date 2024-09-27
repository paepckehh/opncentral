package main

import (
	"fmt"
	"os"
	"time"

	"paepcke.de/opnborg"
)

const (
	_app     = "[OPNBORG-CLI]"
	_version = "[v0.0.3]"
)

func main() {

	// Startup
	t0 := time.Now()
	fmt.Println(_app + "[STARTUP]" + _version)

	// Configure
	config, err := opnborg.Setup()
	if err != nil {
		fmt.Printf(_app+"[ERROR][EXIT] %s\n", err)
		os.Exit(1)
	}

	// Perform Backup of all Appliances xml configuration
	err = opnborg.Start(config)
	if err != nil {
		fmt.Printf(_app+"[ERROR][EXIT] %s\n", err)
		os.Exit(1)
	}

	// Finish
	fmt.Println(_app + "[END][RUNTIME]:" + time.Now().Sub(t0).String())
}
