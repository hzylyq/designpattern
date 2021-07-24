package game

import (
	"fmt"
	"math/rand"
	"strings"
)

var fruitsName = []string{"苹果", "葡萄", "香蕉", "橘子"}

type Gamer struct {
	*Memento
}

func (g *Gamer) Bet() {
	dice := rand.Intn(6) + 1
	if dice == 1 {
		g.money += 100
		fmt.Println("所持金钱增加了。")
	} else if dice == 2 {
		g.money /= 2
		fmt.Println("所持金钱减半了。")
	} else if dice == 6 {
		f := g.Fruits()
		fmt.Println("所持金钱减半了。")
		g.fruits = append(g.fruits, f...)
	} else {
		fmt.Println("什么都没有发生。")
	}
}

func (g *Gamer) Fruit() string {
	prefix := ""
	if rand.Intn(2) == 1 {
		prefix = "好吃的"
	}
	return prefix + fruitsName[rand.Intn(len(fruitsName))]
}

func (g *Gamer) CreateMemento() *Memento {
	m := &Memento{money: g.money}
	for _, fruit := range g.fruits {
		if strings.HasPrefix("好吃的", fruit) {
			m.AddFruits(fruit)
		}
	}
	return m
}

func (g *Gamer) RestoreMemento(m *Memento) {
	g.money = m.money
	g.fruits = m.fruits
}

func (g *Gamer) ToString() {
	fmt.Printf("[money = %d, , fruits = %v", g.money, g.fruits)
}
