package command

type MacroCommand struct {
	Commands []Command
}

func (c *MacroCommand) execute() {
	for _, cmd := range c.Commands {
		cmd.execute()
	}
}

func (c *MacroCommand) append(cmd Command) {
	if c != cmd {
		c.Commands = append(c.Commands, cmd)
	}
}

func (c *MacroCommand) undo() {
	if len(c.Commands) > 0 {
		c.Commands = c.Commands[:len(c.Commands)-2]
	}
}

func (c *MacroCommand) clear() {
	c.Commands = nil
}
