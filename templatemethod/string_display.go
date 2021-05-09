package templatemethod

import "fmt"

type stringDisplay struct {
	str   string
	width int
}

func (s *stringDisplay) open() {
	s.printLine()
}

func (s *stringDisplay) print() {
	fmt.Print("|" + s.str + "|\n")
}

func (s *stringDisplay) close() {
	s.printLine()
}

func (s *stringDisplay) printLine() {
	fmt.Print("+")
	for i := 0; i < s.width; i++ {
		fmt.Print("-")
	}
	fmt.Print("+")
}
