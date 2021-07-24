package factory

type Item interface {
	makeHTML() string
}

type IFactory interface {
	createLink(caption, url string) Item
	createTray(caption string) Item
	createPage(title, author string) Item
}
