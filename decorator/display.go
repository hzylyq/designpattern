package decorator

type IDisplay interface {
	Columns() int
	Rows() int
	RowText(row int) string
	show()
}
