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

	version, err := client.GetVersion()
	if err != nil {
		panic(err)
	}

	fmt.Println(version)
}
