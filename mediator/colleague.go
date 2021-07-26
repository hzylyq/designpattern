package mediator

type Colleague interface {
	setMediator(m Mediator)
	setColleagueEnable(enabled bool)
}
