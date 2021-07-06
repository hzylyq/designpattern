package decorator

type SideBorder struct {
	borderChar byte
	Border
}

func NewSideBorder(d IDisplay, ch byte) *SideBorder {
	return &SideBorder{
		Border: Border{
			IDisplay: d,
		},
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
