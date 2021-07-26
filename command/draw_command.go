package command

type DrawCommand struct {
	drawable  Drawable
	positionX int
	positionY int
}

func (c *DrawCommand) execute() {
	c.drawable.draw(c.positionX, c.positionY)
}

func NewDrawCommand(drawable Drawable, x, y int) Command {
	return &DrawCommand{
		drawable:  drawable,
		positionX: x,
		positionY: y,
	}
}
