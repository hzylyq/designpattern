package factorymethod

import "fmt"

func genGun(gunType string) (iGun, error) {
	switch gunType {
	case "ak47":
		return newAk47(), nil
	case "maverick":
		return newMaverick(), nil
	default:
		return nil, fmt.Errorf("wrong gun type passed")
	}
}

func printDetails(g iGun) {
	fmt.Printf("Gun: %s", g.Name())
	fmt.Println()
	fmt.Printf("Power: %d", g.Power())
	fmt.Println()
}
