package factorymethod

type iGun interface {
	Name() string
	Power() int
	SetName(string)
	SetPower(int)
}
