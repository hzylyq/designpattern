package composite

import (
	"log"
	"testing"
)

func TestComposite(t *testing.T) {
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
	rootDir.PrintList()

}
