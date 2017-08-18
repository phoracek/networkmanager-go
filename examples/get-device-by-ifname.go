package main

import (
	"fmt"
	"os"

	networkmanager "github.com/phoracek/networkmanager-go/src"
)

func main() {
	deviceIface := os.Args[1]

	client := networkmanager.NewClient()
	defer client.Close()

	device := client.GetDeviceByIpIface(deviceIface)
	fmt.Printf("%s [%s]\n", device.Interface, device.Type)
}
