package adapter

import "fmt"

type windows struct {
}

func (w *windows) insertInSquarePort() {
	fmt.Println("Insert circle port into windows machine")
}
