package state

type state interface {
	doClock(c Context, hour int)
	doUse(c Context)
	doAlarm(c Context)
	doPhone(c Context)
}
