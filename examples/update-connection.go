package main

import (
	"fmt"
	"os"

	networkmanager "github.com/phoracek/networkmanager-go/src"
)

func main() {
	client, err := networkmanager.NewClient()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	if connection := findConnection(client, "bridge-br-test"); connection != nil {
		settings, _ := connection.GetSettings()
		newSettings := map[string]map[string]interface{}{
			"connection": map[string]interface{}{
				"uuid":           settings["connection"]["uuid"],
				"type":           "bridge",
				"interface-name": "br-test",
				"id":             "bridge-br-test",
				"autoconnect":    true,
			},
			"ipv4": map[string]interface{}{
				"method": "manual",
				"address-data": []map[string]interface{}{
					map[string]interface{}{
						"address": "11.11.11.11",
						"prefix":  24,
					},
				},
			},
		}
		err := connection.Update(newSettings)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed", err)
			os.Exit(1)
		}
	}
}

func findConnection(client *networkmanager.Client, connectionID string) *networkmanager.Connection {
	connections, err := client.ListConnections()
	if err != nil {
		panic(err)
	}

	for _, connection := range connections {
		settings, _ := connection.GetSettings()
		if settings["connection"]["id"] == connectionID {
			return connection
		}
	}

	return nil
}
