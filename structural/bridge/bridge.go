package bridge

import (
	"fmt"
)

// wiki: https://en.wikipedia.org/wiki/Bridge_pattern

/**
 * Bridge is a structural design pattern that split a large class or a set of closely related classes into two separate
 * hierarchies-abstraction and implementation, which can be developed independently of each other.
 */

// computer
type computer interface {
	print()
	setPrinter(printer)
}

// mac
type mac struct {
	printer printer
}

func (mac *mac) print() {
	fmt.Println("Print request for mac")
	mac.printer.printFile()
}

func (mac *mac) setPrinter(p printer) {
	mac.printer = p
}

// windows
type windows struct {
	printer printer
}

func (w *windows) print() {
	fmt.Println("Print request for windows")
	w.printer.printFile()
}

func (w *windows) setPrinter(p printer) {
	w.printer = p
}

type printer interface {
	printFile()
}

type epson struct{}

func (e *epson) printFile() {
	fmt.Println("Printing by a EPSON Printer")
}

type hp struct{}

func (hp *hp) printFile() {
	fmt.Println("Printing by a HP Printer")
}
