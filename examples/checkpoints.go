package main

import networkmanager "github.com/phoracek/networkmanager-go/src"

func main() {
	client := networkmanager.NewClient()
	defer client.Close()

	checkpoint := client.CheckpointCreate(nil, networkmanager.NmCheckpointNoTimeout, networkmanager.NmCheckpointCreateFlagNone)
	client.CheckpointDestroy(checkpoint)

	checkpoint = client.CheckpointCreate(nil, networkmanager.NmCheckpointNoTimeout, networkmanager.NmCheckpointCreateFlagNone)
	client.CheckpointRollback(checkpoint)
}
