package observer

import "math/rand"

type RandomGenerator struct {
	NumberGenerator
}

func (r *RandomGenerator) Number() int {
	return r.number
}

func (r *RandomGenerator) Execute() {
	for i := 0; i < 20; i++ {
		r.number = rand.Intn(50)
		r.notify()
	}
}
