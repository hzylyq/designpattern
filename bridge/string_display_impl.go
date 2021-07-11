package bridge

import "fmt"

type StringDisplayImpl struct {
	width int
	str   string
}

func NewStringDisplayImpl(str string) DisplayImpl {
	return &StringDisplayImpl{
		width: len(str),
		str:   str,
	}
}

func (s *StringDisplayImpl) rawOpen() {
	s.printLine()
}

func (s *StringDisplayImpl) rawPrint() {
	fmt.Println("|" + s.str + "|")
}

func (s *StringDisplayImpl) rawClose() {
	s.printLine()
}

func (s *StringDisplayImpl) printLine() {
	fmt.Print("+")
	for i := 0; i < s.width; i++ {
		fmt.Print("-")
	}
	fmt.Println("+")
}
