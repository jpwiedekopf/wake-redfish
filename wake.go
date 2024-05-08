package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stmcginnis/gofish"
	"github.com/stmcginnis/gofish/redfish"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	clientConfig := gofish.ClientConfig{
		Insecure: true,
		Endpoint: os.Getenv("REDFISH_ENDPOINT"),
		Username: os.Getenv("REDFISH_USERNAME"),
		Password: os.Getenv("REDFISH_PASSWORD"),
	}
	connect, err := gofish.Connect(clientConfig)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Connected to %s\n", clientConfig.Endpoint)

	service := connect.Service
	systems, err := service.Systems()
	if err != nil {
		panic(err)
	}

	for _, sys := range systems {
		resetType := redfish.ResetType("On")
		powerState := sys.PowerState
		if powerState == ("On") {
			fmt.Printf("System %s is already on.\n", clientConfig.Endpoint)
			return
		}
		fmt.Printf("Current power state: %s\n", powerState)
		err := sys.Reset(resetType)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Sent reset type %s to %s\n", resetType, clientConfig.Endpoint)
	}
}
