package command

import "fmt"

// wiki: https://en.wikipedia.org/wiki/Command_pattern
//
// Bundles a command and arguments to call later.

type command interface {
	execute()
}

type device interface {
	on()
	off()
}

type button struct {
	command command
}

func (b *button) press() {
	b.command.execute()
}

type onCommand struct {
	device device
}

func (c *onCommand) execute() {
	c.device.on()
}

type offCommand struct {
	device device
}

func (c *offCommand) execute() {
	c.device.off()
}

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
