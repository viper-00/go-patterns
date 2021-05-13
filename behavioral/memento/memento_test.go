package memento

import "testing"

func TestPattern(t *testing.T) {
	caretaker := &caretaker{
		mementoArray: make([]*memento, 0),
	}

	originator := &originator{state: "a"}

	t.Logf("originator current state: %s\n", originator.getState())

	caretaker.addMemento(originator.createMemento())

	originator.setState("b")
	t.Logf("originator current state: %s\n", originator.getState())
	caretaker.addMemento(originator.createMemento())

	originator.setState("c")
	t.Logf("originator current state: %s\n", originator.getState())
	caretaker.addMemento(originator.createMemento())

	originator.restoreMemento(caretaker.getMemento(1))
	t.Logf("restored to state: %s\n", originator.getState())

	originator.restoreMemento(caretaker.getMemento(0))
	t.Logf("restored to state: %s\n", originator.getState())
}
