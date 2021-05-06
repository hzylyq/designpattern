package adapter

import "testing"

func Test(t *testing.T) {
	client := &client{}
	mac := &mac{}
	client.insertSquareUsbInComputer(mac)
	windows := &windows{}
	client.insertSquareUsbInComputer(windows)
	
	windowsAdapter := &windowsAdapter{
		WindowsMachine: windows,
	}
	client.insertSquareUsbInComputer(windowsAdapter)
}
