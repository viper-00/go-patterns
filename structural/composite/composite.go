package composite

import "fmt"

// wiki: https://en.wikipedia.org/wiki/Composite_pattern

/**
 * Composite is a structural design pattern that compose objects into tree structures and
 * then work with these structures as if they were individual objects.
 */

// Leaf
type node interface {
	search(string) bool
}

// file
type file struct {
	name string
}

func (f *file) search(keyword string) bool {
	fmt.Printf("searching for keyword %s in file %s\n", keyword, f.name)
	if keyword == f.name {
		fmt.Println("get it")
		return true
	}
	return false
}

func (f *file) getName() string {
	return f.name
}

// folder
type folder struct {
	node []node
	name string
}

func (f *folder) search(keyword string) bool {
	fmt.Printf("searching recursively for keyword %s in folder %s\n", keyword, f.name)
	for _, composite := range f.node {
		if isTrue := composite.search(keyword); isTrue {
			return true
		}
	}

	return false
}

func (f *folder) add(node node) {
	f.node = append(f.node, node)
}
