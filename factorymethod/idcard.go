package factorymethod

import "fmt"

type IdCard struct {
	Owner string
}

func New(owner string) *IdCard {
	 fmt.Println()
	return &IdCard{
		Owner: owner,
	}
}

func (card *IdCard) use() {
	fmt.Println("")
}
