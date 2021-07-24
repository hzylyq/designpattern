package state

type Context interface {
	setClock(hour int)
	changeState(s state)
	callSecurityCenter(msg string)
	recordLog(msg string)
}
