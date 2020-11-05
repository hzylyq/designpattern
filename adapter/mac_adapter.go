package adapter

type macAdapter struct {
	macMachine *mac
}

func (m *macAdapter) insertInSquarePort() {
	m.macMachine.insertInSquarePort()
}
