package abstractfactory

type Tray struct {
	List []Item
}

func (t *Tray) add(item Item) {
	t.List = append(t.List, item)
}

type ListTray struct {
}


