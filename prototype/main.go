package main

func main() {
	manager := NewManager()

	uPen := NewUnderlinePen('~')
	mBox := NewMessageBox('*')
	sBox := NewMessageBox('/')

	manager.register("strong message", uPen)
	manager.register("warning box", mBox)
	manager.register("slash box", sBox)

	// 生成
	p1 := manager.create("strong message")
	pen := p1.(*UnderlinePen)
	pen.use("Hello world.")

	p2 := manager.create("warning box")
	mBox = p2.(*MessageBox)
	mBox.use("Hello world.")

	p3 := manager.create("slash box")
	sBox = p3.(*MessageBox)
	sBox.use("Hello world.")
}
