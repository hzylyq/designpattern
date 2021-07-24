package state

type NightState struct {
}

func (s *NightState) doClock(c Context, hour int) {
	if hour >= 9 && hour <= 17 {
		c.changeState(s)
	}
}

func (s *NightState) doUse(c Context) {
	c.recordLog("紧急: 晚上使用金库!")
}

func (s *NightState) doAlarm(c Context) {
	c.callSecurityCenter("按下警铃(晚上)")
}
func (s *NightState) doPhone(c Context) {
	c.callSecurityCenter("晚上的通话录音")
}
