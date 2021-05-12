package chainOfResponsibility

import "testing"

func TestPattern(t *testing.T) {
	cashier := &cashier{}

	medical := &medical{}
	medical.setNext(cashier)

	doctor := &doctor{}
	doctor.setNext(medical)

	reception := &reception{}
	reception.setNext(doctor)

	patient := &patient{name: "zmy"}

	t.Log("patient visiting first")
	reception.execute(patient)

	t.Log("patient visiting second")
	reception.execute(patient)
}
