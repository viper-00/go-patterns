package abstractFactory

import "fmt"

// wiki: https://en.wikipedia.org/wiki/Abstract_factory_pattern

/*
Abstract Factory is a creational design pattern, which solves the problem of
provide an interface for creating families of related or dependent objects
without specifiying their concrete classes.
*/

type Button interface {
	print() string
}

type LinuxButton struct{}

func (linux *LinuxButton) print() string {
	return "Render a button in a Linux style"
}

type WindowButton struct{}

func (window *WindowButton) print() string {
	return "Render a button in a window style"
}

type MacOsButton struct{}

func (macos *MacOsButton) print() string {
	return "Render a button in a macos style"
}

type GUIFactroy interface {
	createButton() Button
}

type Linux struct{}

func (linux *Linux) createButton() Button {
	return &LinuxButton{}
}

type Window struct{}

func (window *Window) createButton() Button {
	return &WindowButton{}
}

type Macos struct{}

func (macos *Macos) createButton() Button {
	return &MacOsButton{}
}

func GetGUIFactory(os string) (GUIFactroy, error) {
	if os == "linux" {
		return &Linux{}, nil
	} else if os == "window" {
		return &Window{}, nil
	} else if os == "macos" {
		return &Macos{}, nil
	}

	return nil, fmt.Errorf("wrong os type passed")
}
