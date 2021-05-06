package abstractFactory

import "testing"

func TestPattern(t *testing.T) {
	linux, err := GetGUIFactory("linux")

	if err != nil {
		t.Error("test pattern has mistake!")
	}

	linuxButton := linux.createButton()

	t.Log(linuxButton.print())
}
