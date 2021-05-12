package command

import "fmt"

// wiki: https://en.wikipedia.org/wiki/Command_pattern

/**
 * Command is a behavioral design pattern that turns a request into a stand-alone object that contains all information
 * about the request. This transformation pass requests as a method arguments, delay or queue a request's excution,
 * and support undoable operation.
 */

// invoke - button
type button struct {
	command command
}

func (b *button) press() {
	b.command.execute()
}

// command interface - command
type command interface {
	execute()
}

// concrete command - onCommand
type onCommand struct {
	device device
}

func (c *onCommand) execute() {
	c.device.on()
}

// concrete command - offCommand
type offCommand struct {
	device device
}

func (c *offCommand) execute() {
	c.device.off()
}

// receiver interface
type device interface {
	on()
	off()
}

// concrete receiver
type tv struct {
	isRunning bool
}

func (t *tv) on() {
	t.isRunning = true
	fmt.Println("Turning tv on")
}

func (t *tv) off() {
	t.isRunning = false
	fmt.Println("Turning tv off")
}
