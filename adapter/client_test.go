package adapter

import "testing"

func Test(t *testing.T) {
	client := &client{}
	mac := &mac{}
	windows := &windows{}

	client.insertInSquareUsbInComputer(mac)
	client.insertInSquareUsbInComputer(windows)

	windowsAdapter := &windowsAdapter{
		windowMachine: windows,
	}
	client.insertInSquareUsbInComputer(windowsAdapter)
	macAdapter := &macAdapter{
		macMachine: mac,
	}
	client.insertInSquareUsbInComputer(macAdapter)
}
