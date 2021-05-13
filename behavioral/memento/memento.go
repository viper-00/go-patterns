package memento

// wiki: https://en.wikipedia.org/wiki/Memento_pattern

/**
 * Memento is a behavioral design pattern that save and restore the previous state
 * of an object without revealing the details of its implementation.
 */

// originator
type originator struct {
	state string
}

func (o *originator) createMemento() *memento {
	return &memento{state: o.state}
}

func (o *originator) restoreMemento(m *memento) {
	o.state = m.getSavedState()
}

func (o *originator) setState(state string) {
	o.state = state
}

func (o *originator) getState() string {
	return o.state
}

// memento
type memento struct {
	state string
}

func (m *memento) getSavedState() string {
	return m.state
}

// caretaker
type caretaker struct {
	mementoArray []*memento
}

func (c *caretaker) addMemento(m *memento) {
	c.mementoArray = append(c.mementoArray, m)
}

func (c *caretaker) getMemento(index int) *memento {
	if index < len(c.mementoArray) {
		return c.mementoArray[index]
	}
	return nil
}
