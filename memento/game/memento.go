package game

type Memento struct {
	money  int
	fruits []string
}

func NewMemento(money int) *Memento {
	return &Memento{
		money:  money,
		fruits: nil,
	}
}

func (m *Memento) AddFruits(fruit string) {
	m.fruits = append(m.fruits, fruit)
}

func (m *Memento) Fruits() []string {
	return m.fruits
}
