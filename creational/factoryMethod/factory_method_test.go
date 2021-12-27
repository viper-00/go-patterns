package factorymethod

import "testing"

func TestPattern(t *testing.T) {
	window, err := GetDislogStyle(window)

	if err != nil {
		t.Error(err)
	} else {
		winButton := window.createButton()

		t.Log(winButton.render())

		t.Log(winButton.onClick())
	}
}
