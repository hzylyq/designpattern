package decorator

import "log"

type SideBorder struct {
	borderChar byte
	IDisplay
}

func NewSideBorder(d IDisplay, ch byte) *SideBorder {
	return &SideBorder{
		IDisplay:   d,
		borderChar: ch,
	}
}

func (s *SideBorder) Columns() int {
	return 1 + s.IDisplay.Columns() + 1
}

func (s *SideBorder) Rows() int {
	return s.IDisplay.Rows()
}

func (s *SideBorder) RowText(row int) string {
	return string(s.borderChar) + s.IDisplay.RowText(row) + string(s.borderChar)
}

func (s *SideBorder) show() {
	for i := 0; i < s.Rows(); i++ {
		log.Println(s.RowText(i))
	}
}
