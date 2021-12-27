package abstractFactory

import (
	"fmt"
)

// wiki: https://en.wikipedia.org/wiki/Abstract_factory_pattern
//
// Provides an interface for creating families of related objects.

type osType int

const (
	Linux  osType = 1
	Mac    osType = 2
	Window osType = 3
)

type GUIFactroy interface {
	createButton() Button
	createCheckbox() Checkbox
}

type LinuxFactory struct{}

func (linux *LinuxFactory) createButton() Button {
	return &LinuxButton{}
}

func (linux *LinuxFactory) createCheckbox() Checkbox {
	return &LinuxCheckbox{}
}

type WindowFactory struct{}

func (window *WindowFactory) createButton() Button {
	return &WindowButton{}
}

func (window *WindowFactory) createCheckbox() Checkbox {
	return &WindowCheckbox{}
}

type MacosFactory struct{}

func (macos *MacosFactory) createButton() Button {
	return &MacosButton{}
}

func (macos *MacosFactory) createCheckbox() Checkbox {
	return &MacosCheckbox{}
}

type Button interface {
	print() string
}

type LinuxButton struct{}

func (linux *LinuxButton) print() string {
	return "Render a button in a Linux style"
}

type WindowButton struct{}

func (window *WindowButton) print() string {
	return "Render a button in a Window style"
}

type MacosButton struct{}

func (macos *MacosButton) print() string {
	return "Render a button in a Macos style"
}

type Checkbox interface {
	print() string
}

type LinuxCheckbox struct{}

func (linux *LinuxCheckbox) print() string {
	return "Render a checkbox in a Linux style"
}

type WindowCheckbox struct{}

func (window *WindowCheckbox) print() string {
	return "Render a checkbox in a Window style"
}

type MacosCheckbox struct{}

func (macos *MacosCheckbox) print() string {
	return "Render a checkbox in a Macos style"
}

func GetGUIFactory(os osType) (GUIFactroy, error) {
	switch os {
	case Linux:
		return new(LinuxFactory), nil
	case Window:
		return new(WindowFactory), nil
	case Mac:
		return new(MacosFactory), nil
	default:
		return nil, fmt.Errorf(fmt.Sprintf("Factory of type %d is not exist", os))
	}
}
