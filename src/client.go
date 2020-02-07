package networkmanager

import (
	"github.com/godbus/dbus"
)

const (
	InterfacePath = "org.freedesktop.NetworkManager"
	ObjectPath    = "/org/freedesktop/NetworkManager"
)

type Client struct {
	conn *dbus.Conn
}

func NewClient() (*Client, error) {
	client := new(Client)

	dbusConn, err := dbus.SystemBus()
	if err != nil {
		return client, err
	}

	client.conn = dbusConn
	return client, nil
}

func NewClientPrivate() (*Client, error) {
	client := new(Client)

	dbusConn, err := dbus.SystemBusPrivate()
	if err != nil {
		return client, err
	}

	client.conn = dbusConn
	return client, nil
}

func (client *Client) Close() {
	client.conn.Close()
	client.conn = nil
}
