package main

import networkmanager "github.com/phoracek/networkmanager-go/src"

func main() {
	client, err := networkmanager.NewClient()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	checkpoint, err := client.CheckpointCreate(nil, networkmanager.NmCheckpointNoTimeout, networkmanager.NmCheckpointCreateFlagNone)
	if err != nil {
		panic(err)
	}

	client.CheckpointDestroy(checkpoint)

	checkpoint, err = client.CheckpointCreate(nil, networkmanager.NmCheckpointNoTimeout, networkmanager.NmCheckpointCreateFlagNone)
	if err != nil {
		panic(err)
	}

	client.CheckpointRollback(checkpoint)
}
