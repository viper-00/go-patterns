package composite

import "fmt"

// wiki: https://en.wikipedia.org/wiki/Composite_pattern
//
// Encapsulates and provides access to a number of different objects.

type node interface {
	search(string) bool
}

type file struct {
	name string
}

func (f *file) search(keyword string) bool {
	fmt.Printf("searching for keyword %s in file %s\n", keyword, f.getName())
	if keyword == f.getName() {
		fmt.Println("got it")
		return true
	}
	return false
}

func (f *file) getName() string {
	return f.name
}

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
