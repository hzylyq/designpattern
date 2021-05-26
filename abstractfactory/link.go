package abstractfactory

type Link struct {
	Url string
	Item
}

type ListLink struct {
	Caption string
	Url     string
}

func (l *ListLink) makeHTML() string {
	return " <li><a href=\\" + l.Url + "\\" + l.Caption + "</a></li>\n"
}
