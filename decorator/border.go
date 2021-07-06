package decorator

import "log"

type Border struct {
	IDisplay
}

func (d *Border) show() {
	for i := 0; i < d.Rows(); i++ {
		log.Println(d.RowText(i))
	}
}
