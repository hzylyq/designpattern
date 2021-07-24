package factory

type ListTray struct {
	Caption string
}

func NewListTray(caption string) Item {
	return &ListTray{
		Caption: caption,
	}
}

func (l *ListTray) add(item Item) {
	t.List = append(t.List, item)
}

func (l *ListTray) makeHTML() string {

}
