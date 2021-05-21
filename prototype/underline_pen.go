package main

type UnderlinePen struct {
	ulChar byte
}

func NewUnderlinePen(b byte) *UnderlinePen {
	return &UnderlinePen{
		ulChar: b,
	}
}

func (u *UnderlinePen) use(s string) {
	length := len(s)
	println("\\" + s + "\\")
	print(" ")
	for i := 0; i < length; i++ {
		print(string(u.ulChar))
	}
	println(" ")
}

func (u *UnderlinePen) Clone() Cloneable {
	return &UnderlinePen{
		ulChar: u.ulChar,
	}
}
