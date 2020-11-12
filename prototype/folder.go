package prototype

import "fmt"

type folder struct {
	children []inode
	name     string
}

func (f *folder) print(indentation string) {
	fmt.Println(indentation + f.name + "_clone")
	for _, i := range f.children {
		i.print(indentation)
	}
}

func (f *folder) clone() inode {
	cloneFolder := &folder{
		name: f.name,
	}
	var tempChildren []inode
	for _, item := range f.children {
		itemCopy := item.clone()
		tempChildren = append(tempChildren, itemCopy)
	}
	cloneFolder.children = tempChildren
	return cloneFolder
}
