package main

type IProduct interface {
	Cloneable
	use(s string)
}

type MessageBox struct {
	decoChar byte
}

func (m *MessageBox) use(s string) {
	length := len(s)
	for i := 0; i < length+4; i++ {
		print(m.decoChar)
	}

	println("")
	println(string(m.decoChar) + " " + s + " " + string(m.decoChar))
	for i := 0; i < length+4; i++ {
		print(m.decoChar)
	}
	println("")
}

func (m *MessageBox) Clone() Cloneable {
	return &MessageBox{
		decoChar: m.decoChar,
	}
}
