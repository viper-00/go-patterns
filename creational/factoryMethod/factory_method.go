package factorymethod

import "fmt"

// wiki: https://en.wikipedia.org/wiki/Factory_method_pattern

/**
 * Factory Method is a creational design pattern that provides
 * an interface for creating objects in a superclass, but allows
 * subclasses to alter the type of objects that will be created.
 */

type Dislog interface {
	createButton() Button
}

type WindowDialog struct{}

func (window *WindowDialog) createButton() Button {
	return &WindowButton{}
}

type WebDialog struct{}

func (web *WebDialog) createButton() Button {
	return &WebButton{}
}

type Button interface {
	render() string
	onClick() bool
}

type WindowButton struct{}

func (window *WindowButton) render() string {
	return "Render a button in a window style"
}

func (window *WindowButton) onClick() bool {
	return true
}

type WebButton struct{}

func (web *WebButton) render() string {
	return "Render a button in a web style"
}

func (web *WebButton) onClick() bool {
	return true
}

func GetDislogStyle(platform string) (Dislog, error) {
	if platform == "window" {
		return &WindowDialog{}, nil
	} else if platform == "web" {
		return &WebDialog{}, nil
	}

	return nil, fmt.Errorf("wrong platform type passed")
}
