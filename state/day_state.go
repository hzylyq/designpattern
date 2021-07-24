package state

type DayState struct {
}

func (s *DayState) doClock(c Context, hour int) {
	if hour < 9 || hour >= 17 {
		c.changeState(s)
	}
}

func (s *DayState) doUse(c Context) {
	c.recordLog("使用金库(白天)")
}

func (s *DayState) doAlarm(c Context) {
	c.callSecurityCenter("按下警铃(白天)")
}
func (s *DayState) doPhone(c Context) {
	c.callSecurityCenter("正常通话(白天)")
}
