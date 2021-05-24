package abstractfactory

type ListLink struct {
	Caption string
	Url     string
}

func (l *ListLink) makeHTML() string {
	return " <li><a href=\\" + l.Url + "\\" + l.Caption + "</a></li>\n"
}

type ListPage struct {
	Title  string
	Author string
}

func (l *ListPage) makeHTML() string {
	return ""
}
