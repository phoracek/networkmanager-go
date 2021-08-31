package networkmanager

import (
	"os/exec"
	"strings"
	"testing"
)

func getNetworkManagerVersion() string {
	cmd := exec.Command("NetworkManager", "--version")
	stdout, err := cmd.Output()
	if err != nil {
		return ""
	}
	return string(stdout)
}

func TestGetVersion(t *testing.T) {
	want := getNetworkManagerVersion()
	client, err := NewClientPrivate()
	if err != nil {
		panic(err)
	}
	defer client.Close()
	version, err := client.GetVersion()
	if err != nil {
		t.Errorf("expecting err to be nil, %q", err)
	}
	if !strings.Contains(want, version) {
		t.Errorf(`client.GetVersion() = %q, want match for %q`, version,  want)
	}
}
