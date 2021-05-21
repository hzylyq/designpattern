package main

type MessageBox struct {
	decoChar byte
}

func NewMessageBox(b byte) *MessageBox {
	return &MessageBox{
		decoChar: b,
	}
}

func (m *MessageBox) use(s string) {
	length := len(s)
	for i := 0; i < length+4; i++ {
		print(string(m.decoChar))
	}

	println("")
	println(string(m.decoChar) + " " + s + " " + string(m.decoChar))
	for i := 0; i < length+4; i++ {
		print(string(m.decoChar))
	}
	println("")
}

func (m *MessageBox) Clone() Cloneable {
	return &MessageBox{
		decoChar: m.decoChar,
	}
}
