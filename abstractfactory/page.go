package abstractfactory

type Page interface {
	Item
}

type ListLink struct {
	Caption string
	Url     string
}

func (l *ListLink) makeHTML() string {
	return " <li><a href=\\" + l.Url + "\\" + l.Caption + "</a></li>\n"
}

type TablePage struct {
	Title  string
	Author string
}

func (l *TablePage) makeHTML() string {
	var b []byte
	b = append(b, []byte("<html><head><title>"+l.Title+"</title></head><html")...)
	b = append(b, []byte("<body>\n")...)
	b = append(b, []byte("<h1>"+l.Title+"</h1>\\n")...)
	b = append(b, []byte("table width=\"80%\" border=\"3\">\n")...)
	
	// todo range list
	
	b = append(b, []byte("</table>\n")...)
	b = append(b, []byte("<hr><address>"+l.Author+"</address>")...)
	b = append(b, []byte("</body></html>\n")...)
	
	return string(b)
}
