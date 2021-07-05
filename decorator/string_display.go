package decorator

import "log"

type Display struct {
	str string
}

func NewDisplay(str string) *Display {
	return &Display{str: str}
}

func (d *Display) Columns() int {
	return len(d.str)
}

func (d *Display) Rows() int {
	return 1
}

func (d *Display) RowText(row int) string {
	if row == 0 {
		return d.str
	}

	return ""
}

func (d *Display) show() {
	for i := 0; i < d.Rows(); i++ {
		log.Println(d.RowText(i))
	}
}
