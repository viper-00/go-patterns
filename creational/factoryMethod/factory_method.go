package factorymethod

import "fmt"

// wiki: https://en.wikipedia.org/wiki/Factory_method_pattern
//
// Defers instantiation of an object to a specialized function for creating instances.

type platform int

var (
	window platform = 1
	web    platform = 2
)

type Dislog interface {
	createButton() Button
}

type WindowDialog struct{}

func (window *WindowDialog) createButton() Button {
	return new(WindowButton)
}

type WebDialog struct{}

func (web *WebDialog) createButton() Button {
	return new(WebButton)
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

func GetDislogStyle(desktop platform) (Dislog, error) {
	switch desktop {
	case window:
		return new(WindowDialog), nil
	case web:
		return new(WebDialog), nil
	default:
		return nil, fmt.Errorf(fmt.Sprintf("Platform of type %d is not exist", desktop))
	}
}
