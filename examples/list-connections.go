package main

import (
	"fmt"

	networkmanager "github.com/phoracek/networkmanager-go/src"
)

func main() {
	client, err := networkmanager.NewClient()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	connections, err := client.ListConnections()
	if err != nil {
		panic(err)
	}

	for _, conn := range connections {
		settings, _ := conn.GetSettings()
		fmt.Printf("%s (%s)\n", settings["connection"]["id"], settings["connection"]["uuid"])
	}
}
