package chainOfResponsibility

import "fmt"

// wiki: https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern

/**
 * Chain of responsibility is a behavioral design pattern that pass requests along a chain of
 * handlers. Upon receiving a request, each handler decides either to process the request or
 * to pass it to the next handler in the chain.
 */

// handler interface - department
type department interface {
	execute(*patient)
	setNext(department)
}

type patient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}

// concrete handler - reception
type reception struct {
	next department
}

func (r *reception) execute(p *patient) {
	if p.registrationDone {
		fmt.Println("Patient registration already done")
		r.next.execute(p)
		return
	}
	fmt.Println("Reception registering patient")
	p.registrationDone = true
	r.next.execute(p)
}

func (r *reception) setNext(next department) {
	r.next = next
}

// concrete handler - doctor
type doctor struct {
	next department
}

func (d *doctor) execute(p *patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		d.next.execute(p)
		return
	}
	fmt.Println("Doctor checking patient")
	p.doctorCheckUpDone = true
	d.next.execute(p)
}

func (d *doctor) setNext(next department) {
	d.next = next
}

// concrete handler - medical
type medical struct {
	next department
}

func (m *medical) execute(p *patient) {
	if p.medicineDone {
		fmt.Println("Medicine already given to patient")
		m.next.execute(p)
		return
	}

	fmt.Println("Medical giving medicine to patient")
	p.medicineDone = true
	m.next.execute(p)
}

func (m *medical) setNext(next department) {
	m.next = next
}

// concrete handler - cashier
type cashier struct {
	next department
}

func (c *cashier) execute(p *patient) {
	if p.paymentDone {
		fmt.Println("Payment already done")
		return
	}
	p.paymentDone = true
	fmt.Println("Cashier getting money from patient")
}

func (c *cashier) setNext(next department) {
	c.next = next
}
