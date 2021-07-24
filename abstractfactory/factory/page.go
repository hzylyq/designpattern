package factory

import (
	"io/ioutil"
	"os"
)

type Page struct {
	Title   string
	Author  string
	Content []Item
	Item
}

func (l *Page) Add(item Item) {
	l.Content = append(l.Content, item)
}

func (l *Page) OutPut() error {
	fileName := l.Title + ".html"
	return ioutil.WriteFile(fileName, []byte(l.makeHTML()), os.ModePerm)
}

func (l *Page) makeHTML() string {
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
