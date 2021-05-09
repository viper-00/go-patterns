package bridge

import "testing"

func TestPattern(t *testing.T) {
	hp := &hp{}
	epson := &epson{}

	mac := &mac{}

	// hp printer on mac
	mac.setPrinter(hp)
	mac.print()

	// epson printer on mac
	mac.setPrinter(epson)
	mac.print()

	windows := &windows{}

	// hp printer on windows
	windows.setPrinter(hp)
	windows.print()

	// epson printer on windows
	windows.setPrinter(epson)
	windows.print()
}
