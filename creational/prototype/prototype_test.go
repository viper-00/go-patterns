package prototype

import "testing"

func TestPattern(t *testing.T) {
	d1 := &File{name: "d1"}
	d2 := &File{name: "d2"}
	d3 := &File{name: "d3"}
	d4 := &File{name: "d4"}

	d1Folder := &Folder{
		children: []Node{d1},
		name:     "d1Folder",
	}

	superFolder := &Folder{
		children: []Node{d1Folder, d2, d3},
		name:     "superFolder",
	}

	supersuperFolder := &Folder{
		children: []Node{superFolder, d4},
		name:     "supersuperFolder",
	}

	if superFolder != nil {
		t.Log("printing hierarchy for super Folder:")
		superFolder.print(" ")
	}

	if supersuperFolder != nil {
		t.Log("printing hierarchy for super super Folder:")
		supersuperFolder.print(" ")
	}

	cloneFolder := superFolder.clone()
	if cloneFolder != nil {
		t.Log("printing hierarchy for clone Folder:")
		cloneFolder.print(" ")
	}
}
