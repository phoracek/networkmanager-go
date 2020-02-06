package main

import (
	"fmt"

	networkmanager "github.com/phoracek/networkmanager-go/src"
)

func main() {
	client := networkmanager.NewClient()
	defer client.Close()

	devices := client.GetDevices()
	for _, device := range devices {
		fmt.Printf("%s [%s, %s]\n", device.Interface, device.Type, device.State)
	}
}
