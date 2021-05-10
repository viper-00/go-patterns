package composite

import "testing"

func TestPattern(t *testing.T) {
	f1 := &file{name: "f1"}
	f2 := &file{name: "f2"}
	f3 := &file{name: "f3"}

	folder1 := &folder{name: "folder1"}
	folder2 := &folder{name: "folder2"}
	folder1.add(f1)

	folder2.add(f2)
	folder2.add(f3)
	folder2.add(folder1)

	folder2.search("f1")
}
