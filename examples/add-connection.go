package main

import (
	"fmt"
	"os"

	networkmanager "github.com/phoracek/networkmanager-go/src"
	uuid "github.com/satori/go.uuid"
)

func main() {
	client, err := networkmanager.NewClient()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	bridgeSettings := map[string]map[string]interface{}{
		"connection": map[string]interface{}{
			"uuid":           uuid.NewV4().String(),
			"type":           "bridge",
			"interface-name": "br-test",
			"id":             "bridge-br-test",
			"autoconnect":    true,
		},
		"ipv4": map[string]interface{}{
			"method": "manual",
			"address-data": []map[string]interface{}{
				map[string]interface{}{
					"address": "10.11.11.11",
					"prefix":  24,
				},
			},
		},
	}

	err = client.AddConnection(bridgeSettings)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed", err)
		os.Exit(1)
	}
}
