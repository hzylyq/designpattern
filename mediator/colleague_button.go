package mediator

type colleagueButton struct {
	mediator Mediator
}

func (c *colleagueButton) setMediator(m Mediator) {
	c.mediator = m
}

func (c *colleagueButton) setColleagueEnable(enabled bool) {

}
