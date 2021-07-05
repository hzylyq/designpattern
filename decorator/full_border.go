package decorator

type FullBorder struct {
	Border
}

func NewFullBorder(border Border) *FullBorder {
	return &FullBorder{
		Border: border,
	}
}

func (f *FullBorder) Columns() int {
	return 1 + f.Display.Columns() + 1
}

func (f *FullBorder) Rows() int {
	return 1 + f.Display.Rows() + 1
}

func (f *FullBorder) RowText(row int) string {
	if row == 0 {
		return "+" + f.makeLine('-', f.Display.Columns()) + "+"
	} else if row == f.Display.Rows()+1 {
		return "+" + f.makeLine('-', f.Display.Columns()) + "+"
	} else {
		return "|" + f.Display.RowText(row-1) + "|"
	}
}

func (f *FullBorder) show() {
	panic("implement me")
}

func (f *FullBorder) makeLine(ch byte, count int) string {
	var str []byte
	for i := 0; i < count; i++ {
		str = append(str, ch)
	}
	return string(str)
}
