package abstractFactory

import "fmt"

// wiki: https://en.wikipedia.org/wiki/Abstract_factory_pattern

/*
Abstract Factory is a creational design pattern, which solves the problem of
provide an interface for creating families of related or dependent objects
without specifiying their concrete classes.
*/

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

func GetGUIFactory(os string) (GUIFactroy, error) {
	if os == "linux" {
		return &LinuxFactory{}, nil
	} else if os == "window" {
		return &WindowFactory{}, nil
	} else if os == "macos" {
		return &MacosFactory{}, nil
	}

	return nil, fmt.Errorf("wrong os type passed")
}
