package main

import (
	"fmt"

	networkmanager "github.com/phoracek/networkmanager-go/src"
)

func main() {
	client, err := networkmanager.NewClientPrivate()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	devices, err := client.GetDevices()
	if err != nil {
		panic(err)
	}

	for _, device := range devices {
		fmt.Printf("%s [%s, %s]\n", device.Interface, device.Type, device.State)
	}
}
