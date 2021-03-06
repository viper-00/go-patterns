package adapter

import "fmt"

// wiki: https://en.wikipedia.org/wiki/Adapter_pattern
//
// Allows objects with incompatible interfaces to collaborate.

type client struct{}

func (c *client) insertLightningConnectorIntoComputer(computer computer) {
	fmt.Println("Client inserts Lightning connector into computer")
	computer.insertLightningPort()
}

type computer interface {
	insertLightningPort()
}

type mac struct{}

func (mac *mac) insertLightningPort() {
	fmt.Println("Lightning connector is plugged into mac machine.")
}

type windows struct{}

func (w *windows) insertUSBPort() {
	fmt.Println("USB connector is plugged into windows machine.")
}

type windowsAdapter struct {
	windowsMachine *windows
}

func (w *windowsAdapter) insertLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.windowsMachine.insertUSBPort()
}
