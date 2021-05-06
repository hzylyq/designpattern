package adapter

type windowsAdapter struct {
	WindowsMachine *windows
}

func (w *windowsAdapter) insertInSquarePort() {
	w.WindowsMachine.insertInSquarePort()
}
