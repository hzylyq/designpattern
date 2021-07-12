package strategy

const (
	handValueGuu = 0 // 石头
	handValueCho = 1 // 剪刀
	handValuePaa = 2 // 布
)

var handMap = map[int]string{}

type Hand struct {
	handleValue int
}
