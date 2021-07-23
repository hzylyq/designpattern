package observer

type NumberGenerator struct {
	number    int
	observers []Observer
}

func (n *NumberGenerator) addObserver(o Observer) {
	n.observers = append(n.observers, o)
}

func (n *NumberGenerator) deleteObserver(o Observer) {
	var observers []Observer
	for _, ob := range n.observers {
		if o == ob {
			continue
		}
		observers = append(observers, ob)
	}
	n.observers = observers
}

func (n *NumberGenerator) notify() {
	for _, o := range n.observers {
		o.Update(n)
	}
}
