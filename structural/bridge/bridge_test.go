package bridge

import "testing"

func TestPattern(t *testing.T) {
	hp := &hp{}
	epson := &epson{}

	mac := &mac{}
	mac.setPrinter(hp)
	mac.print()
	mac.setPrinter(epson)
	mac.print()

	windows := &windows{}
	windows.setPrinter(hp)
	windows.print()
	windows.setPrinter(epson)
	windows.print()
}
