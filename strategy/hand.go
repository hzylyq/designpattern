package strategy

const (
	handValueGuu = 0 // 石头
	handValueCho = 1 // 剪刀
	handValuePaa = 2 // 布
)

var handList = []*Hand{{
	handleValue: handValueGuu,
}, {
	handleValue: handValueCho,
}, {
	handleValue: handValuePaa,
}}

type Hand struct {
	handleValue int
}

func (h *Hand) isStrongThan(t *Hand) bool {
	return h.fight(t) == 1
}

func (h *Hand) isWeakThan(t *Hand) bool {
	return h.fight(t) == -1
}

func (h *Hand) fight(t *Hand) int {
	if h == t {
		return 0
	}
	if (h.handleValue+1)%3 == t.handleValue {
		return 1
	}
	return -1
}

func NewHand(handVal int) *Hand {
	return &Hand{
		handleValue: handVal,
	}
}

func getHand(val int) *Hand {
	return handList[val]
}
