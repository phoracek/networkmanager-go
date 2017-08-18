package main

import (
	"fmt"
	"os"

	networkmanager "github.com/phoracek/networkmanager-go/src"
)

func main() {
	connectionIDToDeactivate := os.Args[1]

	client := networkmanager.NewClient()
	defer client.Close()

	if activeConnection := findActiveConnection(client, connectionIDToDeactivate); activeConnection != nil {
		err := client.DeactivateConnection(activeConnection)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed", err)
			os.Exit(1)
		}
	}
}

func findActiveConnection(client *networkmanager.Client, connectionID string) *networkmanager.ActiveConnection {
	activeConnections := client.ListActiveConnections()
	for _, activeConnection := range activeConnections {
		settings, _ := activeConnection.Connection.GetSettings()
		if settings["connection"]["id"] == connectionID {
			return activeConnection
		}
	}
	return nil
}
