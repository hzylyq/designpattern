package visit

import "log"

type Visitor interface {
	Visit(entry)
}

type ListVisitor struct {
	currentDir string
}

func NewListVisitor() *ListVisitor {
	return &ListVisitor{}
}

func (v *ListVisitor) Visit(e entry) {
	switch c := e.(type) {
	case *File:
		log.Print(v.currentDir + "/" + c.Name())
	case *Directory:
		log.Println(v.currentDir + "/" + c.Name())
		saveDir := v.currentDir
		v.currentDir = v.currentDir + "/" + c.Name()
		for _, ent := range c.children {
			ent.Accept(v)
		}
		v.currentDir = saveDir
	}
}
