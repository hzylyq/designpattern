package abstractfactory

type Item interface {
	makeHTML() string
}

type Link interface {
	makeHTML()
}

type Tray interface {
	makeHTML()
}

type IFactory interface {
	createLink(caption, url string) Link
	createTray(caption string) Tray
	createPage(title, author string)
}
