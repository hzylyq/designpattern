package main

// Cloneable 是原型对象需要实现的接口
type Cloneable interface {
	Clone() Cloneable
}

type Manager struct {
	protoTypes map[string]Cloneable
}

func NewManager() *Manager {
	return &Manager{
		protoTypes: make(map[string]Cloneable),
	}
}

func (m *Manager) register(name string, proto Cloneable) {
	m.protoTypes[name] = proto
}

func (m *Manager) create(name string) Cloneable {
	return m.protoTypes[name]
}
