package chainofresponsibility

import "strconv"

type trouble struct {
	number int
}

func (t *trouble) Number() int {
	return t.number
}

func (t *trouble) ToString() string {
	return "[Trouble " + strconv.Itoa(t.number) + "]"
}
