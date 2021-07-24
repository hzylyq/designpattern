package factory

type ListLink struct {
	Caption string
	Url     string
}

func (l *ListLink) makeHTML() string {
	return " <li><a href=\\" + l.Url + "\\" + l.Caption + "</a></li>\n"
}

func NewListLink(caption, url string) Item {
	return &ListLink{
		Caption: caption,
		Url:     url,
	}
}
