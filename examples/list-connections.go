package main

import (
	"fmt"

	networkmanager "github.com/phoracek/networkmanager-go/src"
)

func main() {
	client := networkmanager.NewClient()
	defer client.Close()

	connections := client.ListConnections()
	for _, conn := range connections {
		settings, _ := conn.GetSettings()
		fmt.Printf("%s (%s)\n", settings["connection"]["id"], settings["connection"]["uuid"])
	}
}
