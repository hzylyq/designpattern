package decorator

import "log"

type FullBorder struct {
	Border
}

func NewFullBorder(d IDisplay) *FullBorder {
	return &FullBorder{
		Border: Border{IDisplay: d},
	}
}

func (f *FullBorder) Columns() int {
	return 1 + f.Border.Columns() + 1
}

func (f *FullBorder) Rows() int {
	return 1 + f.Border.Rows() + 1
}

func (f *FullBorder) RowText(row int) string {
	if row == 0 {
		return "+" + f.makeLine('-', f.Border.Columns()) + "+"
	} else if row == f.Border.Rows()+1 {
		return "+" + f.makeLine('-', f.Border.Columns()) + "+"
	} else {
		return "|" + f.Border.RowText(row-1) + "|"
	}
}

func (f *FullBorder) makeLine(ch byte, count int) string {
	var str []byte
	for i := 0; i < count; i++ {
		str = append(str, ch)
	}
	return string(str)
}

func (f *FullBorder) show() {
	for i := 0; i < f.Rows(); i++ {
		log.Println(f.RowText(i))
	}
}
