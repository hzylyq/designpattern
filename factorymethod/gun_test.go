package factorymethod

import (
	"testing"
)

func Test_genGun(t *testing.T) {
	ak47, _ := genGun("ak47")
	maverick, _ := genGun("maverick")

	printDetails(ak47)
	printDetails(maverick)
}
