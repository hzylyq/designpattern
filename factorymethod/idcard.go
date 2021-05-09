package factorymethod

import "fmt"

type IdCard struct {
	Owner string
}

func New(owner string) *IdCard {
	fmt.Println("制作" + owner + "的ID卡。")
	return &IdCard{
		Owner: owner,
	}
}

func (card *IdCard) use() {
	fmt.Println("使用" + card.Owner + "的ID卡。")
}
