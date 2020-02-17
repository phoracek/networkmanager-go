package networkmanager

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	client, err := NewClient()
	if err != nil {
		t.Errorf("err: %v", err)
	}
	defer client.Close()
	if client == nil {
		t.Error("Client instance should not be nil")
	}
}

func TestNewClientPrivate(t *testing.T) {
	client, err := NewClientPrivate()
	if err != nil {
		t.Errorf("err: %v", err)
	}
	defer client.Close()
	if client == nil {
		t.Error("Client instance should not be nil")
	}
}
