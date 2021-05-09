package templatemethod

import "fmt"

type stringDisplay struct {
	str   string
}

func (s *stringDisplay) open() {
	s.printLine()
}

func (s *stringDisplay) print() {
	fmt.Println("|" + s.str + "|")
}

func (s *stringDisplay) close() {
	s.printLine()
}

func (s *stringDisplay) printLine() {
	fmt.Print("+")
	for i := 0; i < len(s.str); i++ {
		fmt.Print("-")
	}
	fmt.Println("+")
}
