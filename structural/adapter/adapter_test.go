package adapter

import (
	"testing"
)

func TestPattern(t *testing.T) {
	client := &client{}

	mac := &mac{}

	client.insertLightningConnectorIntoComputer(mac)

	windows := &windows{}
	windowsAdapter := &windowsAdapter{windows}
	client.insertLightningConnectorIntoComputer(windowsAdapter)
}
