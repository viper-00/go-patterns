package command

import "testing"

func TestPattern(t *testing.T) {
	tv := new(tv)

	onCommand := &onCommand{device: tv}

	offCommand := &offCommand{device: tv}

	onButton := &button{
		command: onCommand,
	}
	onButton.press()

	offButton := &button{
		command: offCommand,
	}
	offButton.press()
}
