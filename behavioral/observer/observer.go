package observer

import "fmt"

// wiki: https://en.wikipedia.org/wiki/Observer_pattern

/**
 * Observer is a behavioral design pattern that define a subscription mechanism
 * to notify multiple objects about any events that happen to the object
 * they're observing.
 */

// subject
type subject interface {
	register(Observer observer)
	deregister(Observer observer)
	notifyAll()
}

// concrete subject - item
type item struct {
	observerList []observer
	name         string
	inStock      bool
}

func newItem(name string) *item {
	return &item{
		name: name,
	}
}

func (i *item) updateAvailability() {
	fmt.Printf("Item %s is now is stock\n", i.name)
	i.inStock = true
	i.notifyAll()
}

func (i *item) register(o observer) {
	i.observerList = append(i.observerList, o)
}

func (i *item) deregister(o observer) {
	i.observerList = removeFromSlice(i.observerList, o)
}

func (i *item) notifyAll() {
	for _, observer := range i.observerList {
		observer.update(i.name)
	}
}

func removeFromSlice(observerList []observer, observerToRemove observer) []observer {
	len := len(observerList)
	for i, observer := range observerList {
		if observer.getID() == observerToRemove.getID() {
			// remove item from index position, item is disorder
			lastItem := observerList[len-1]
			observerList[len-1] = observerList[i]
			observerList[i] = lastItem
			return observerList[:len-1]
		}
	}
	return observerList
}

// observer
type observer interface {
	update(string)
	getID() string
}

// concrete observer - customer
type customer struct {
	id string
}

func (c *customer) update(itemName string) {
	fmt.Printf("Sending email to customer %s for item %s\n", c.id, itemName)
}

func (c *customer) getID() string {
	return c.id
}

// concrete observer - employee
type employee struct {
	id string
}

func (e *employee) update(itemName string) {
	fmt.Printf("Sending email to employee %s for item %s\n", e.id, itemName)
}

func (e *employee) getID() string {
	return e.id
}
