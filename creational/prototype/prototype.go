package prototype

import "fmt"

// wiki: https://en.wikipedia.org/wiki/Prototype_pattern
//
// When the type of objects to create is determined by a prototypical instance, which is cloned to produce new objects.

type Node interface {
	print(string)
	clone() Node
}

type File struct {
	name string
}

func (f *File) print(s string) {
	fmt.Println(s + f.name)
}

func (f *File) clone() Node {
	return &File{name: f.name + "_clone"}
}

type Folder struct {
	children []Node
	name     string
}

func (f *Folder) print(s string) {
	fmt.Println(s + f.name)
	for _, v := range f.children {
		v.print(s + s)
	}
}

func (f *Folder) clone() Node {
	cloneFolder := &Folder{name: f.name + "_clone"}
	var tempChildren []Node
	for _, v := range f.children {
		folder_copy := v.clone()
		tempChildren = append(tempChildren, folder_copy)
	}
	cloneFolder.children = tempChildren
	return cloneFolder
}
