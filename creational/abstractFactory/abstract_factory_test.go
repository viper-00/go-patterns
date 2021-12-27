package abstractFactory

import "testing"

func TestPattern(t *testing.T) {
	linux, err := GetGUIFactory(Linux)

	if err != nil {
		t.Error(err)
	}

	_, ok := linux.(*LinuxFactory)
	if !ok {
		t.Error("linux factory is not exist.")
	} else {
		linuxButton := linux.createButton()

		t.Log(linuxButton.print())

		linuxCheckbox := linux.createCheckbox()

		t.Log(linuxCheckbox.print())
	}
}
