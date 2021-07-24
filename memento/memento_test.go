package memento

import (
	"fmt"
	"github.com/hzylyq/designpattern/memento/game"
	"testing"
)

func TestMemento(t *testing.T) {
	gamer := game.Gamer{
		Memento: game.NewMemento(100),
	}
	memento := gamer.CreateMemento()
	for i := 0; i < 20; i++ {
		fmt.Println("==== ", i)
		fmt.Println("当前状态:", gamer)

		gamer.Bet()
		
	}
}
