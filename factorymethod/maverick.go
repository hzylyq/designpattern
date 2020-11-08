package factorymethod

type maverick struct {
	gun
}

func newMaverick() iGun {
	return &maverick{
		gun: gun{
			name:  "Maverick gun",
			power: 5,
		},
	}
}
