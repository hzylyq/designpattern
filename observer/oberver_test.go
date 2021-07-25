package observer

import "testing"

func TestDigitObserver_Update(t *testing.T) {
	generator := &RandomGenerator{}
	ob1 := &DigitObserver{}
	ob2 := &GraphObserver{}
	generator.addObserver(ob1)
	generator.addObserver(ob2)
	generator.Execute()
}
