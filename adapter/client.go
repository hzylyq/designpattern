package adapter

type client struct {
}

func (c *client) insertInSquareUsbInComputer(com computer) {
	com.insertInSquarePort()
}
