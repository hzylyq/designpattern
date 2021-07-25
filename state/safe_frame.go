package state

import (
	"fmt"
	"strconv"
)

type SafeFrame struct {
	State state
}

func NewSafeFrame() *SafeFrame {
	return &SafeFrame{
		State: &DayState{},
	}
}

func (s *SafeFrame) setClock(hour int) {
	clockStr := "现在时间是"
	if hour < 10 {
		clockStr = clockStr + "0" + strconv.Itoa(hour) + ":00"
	} else {
		clockStr = clockStr + strconv.Itoa(hour) + ":00"
	}
	fmt.Println(clockStr)
	s.State.doClock(s, hour)
}

func (s *SafeFrame) changeState(ste state) {
	fmt.Println("状态改变")
	s.State = ste
}
func (s *SafeFrame) callSecurityCenter(msg string) {
	fmt.Println("call!", msg)
}

func (s *SafeFrame) recordLog(msg string) {
	fmt.Println("record ...", msg)
}
