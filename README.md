# NetworkManager Go wrapper

Golang wrapper for NetworkManager implemented through D-Bus.

List all available connections:

``` go
client, err := networkmanager.NewClient()
if err != nil {
	panic(err)
}
defer client.Close()

connections, err := client.ListConnections()
if err != nil {
	panic(err)
}

for _, conn := range connections {
	settings, err := conn.GetSettings()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s (%s)\n", settings["connection"]["id"], settings["connection"]["uuid"])
}
```

Find more example under the `examples/` directory.
