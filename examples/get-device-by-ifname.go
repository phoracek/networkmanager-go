package main

import (
	"fmt"
	"os"

	networkmanager "github.com/phoracek/networkmanager-go/src"
)

func main() {
	deviceIface := os.Args[1]

	client, err := networkmanager.NewClient()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	device, err := client.GetDeviceByIpIface(deviceIface)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s [%s %s]\n", device.Interface, device.Type, device.State)
}
