package visit

import (
	"log"
	"testing"
)

func TestVisit(t *testing.T) {
	log.Println("Making root entries...")
	rootDir := NewDirectory("root")
	binDir := NewDirectory("bin")
	tmpDir := NewDirectory("tmp")
	userDir := NewDirectory("user")
	rootDir.Add(binDir)
	rootDir.Add(tmpDir)
	rootDir.Add(userDir)
	binDir.Add(NewFile("vi", 10000))
	binDir.Add(NewFile("latex", 20000))
	rootDir.Accept(NewListVisitor())

	log.Println("")
	log.Println("Making user entries...")
	yuki := NewDirectory("yuki")
	hanako := NewDirectory("hanako")
	tomura := NewDirectory("tomura")
	userDir.Add(yuki)
	userDir.Add(hanako)
	userDir.Add(tomura)
	yuki.Add(NewFile("diary.html", 100))
	yuki.Add(NewFile("Composite.go", 200))
	hanako.Add(NewFile("memo.tex", 300))
	tomura.Add(NewFile("game.doc", 400))
	tomura.Add(NewFile("junk.mail", 500))
	rootDir.Accept(NewListVisitor())
}
