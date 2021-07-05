package decorator

type SideBorder struct {
	borderChar byte
	Border
}

func NewSideBorder(disPlay *Display, ch byte) *SideBorder {
	return &SideBorder{Border: Border{
		Display: disPlay,
	},
		borderChar: ch,
	}
}

func (s *SideBorder) Columns() int {
	return 1 + s.Display.Columns() + 1
}

func (s *SideBorder) Rows() int {
	return s.Display.Rows()
}

func (s *SideBorder) RowText(row int) string {
	return string(s.borderChar) + s.Display.RowText(row) + string(s.borderChar)
}
