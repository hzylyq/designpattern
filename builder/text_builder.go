package builder

type TextBuilder struct {
	Buffer string
}

func (b *TextBuilder) makeTitle(title string) {
	b.Buffer += "===================\n"
	b.Buffer += "[" + title + "]\n"
	b.Buffer += "\n"
}

func (b *TextBuilder) makeString(str string) {
	b.Buffer += "-" + str + "\n"
}

func (b *TextBuilder) makeItems(items []string) {
	for _, val := range items {
		b.Buffer += val
		b.Buffer += "\n"
	}
}

func (b *TextBuilder) close() {
	b.Buffer += "=========================\n"
}
