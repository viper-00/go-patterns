package prototype

import "testing"

func TestPattern(t *testing.T) {
	golang := &File{name: "go"}
	java := &File{name: "java"}
	py := &File{name: "python"}
	sub := &File{name: "sub"}

	goFolder := &Folder{
		children: []Node{golang},
		name:     "goFolder",
	}

	superFolder := &Folder{
		children: []Node{goFolder, java, py},
		name:     "superFolder",
	}

	supersuperFolder := &Folder{
		children: []Node{superFolder, sub},
		name:     "supersuperFolder",
	}

	t.Log("printing hierarchy for super Folder:")
	superFolder.print(" ")

	cloneFolder := superFolder.clone()
	t.Log("printing hierarchy for clone Folder:")
	cloneFolder.print(" ")

	t.Log("printing hierarchy for super super Folder:")
	supersuperFolder.print(" ")
}
