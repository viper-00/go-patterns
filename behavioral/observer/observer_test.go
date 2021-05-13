package observer

import "testing"

func TestPattern(t *testing.T) {
	shirtItem := newItem("nike")

	customerObserverFirst := &customer{id: "customer-first@gmail.com"}
	customerObserverSecond := &customer{id: "customer-second@gmail.com"}

	shirtItem.register(customerObserverFirst)
	shirtItem.register(customerObserverSecond)
	shirtItem.updateAvailability()

	shirtItem.deregister(customerObserverSecond)
	shirtItem.updateAvailability()

	employeeObserverFirst := &employee{id: "employee-first@gmail.com"}
	employeeObserverSecond := &employee{id: "employee-second@gmail.com"}

	shirtItem.register(employeeObserverFirst)
	shirtItem.register(employeeObserverSecond)
	shirtItem.updateAvailability()
}
