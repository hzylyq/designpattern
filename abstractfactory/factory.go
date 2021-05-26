package abstractfactory

type Item interface {
	makeHTML() string
}

type IFactory interface {
	createLink(caption, url string) Link
	createTray(caption string) Tray
	createPage(title, author string) Page
}
