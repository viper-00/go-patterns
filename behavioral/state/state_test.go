package state

import "testing"

func TestPattern(t *testing.T) {
	vendingMachine := newVendingMachine(1, 10)

	err := vendingMachine.requestItem()
	if err != nil {
		t.Fatalf(err.Error())
	}

	err = vendingMachine.insertMoney(10)
	if err != nil {
		t.Fatalf(err.Error())
	}

	err = vendingMachine.dispenseItem()
	if err != nil {
		t.Fatalf(err.Error())
	}

	err = vendingMachine.addItem(2)
	if err != nil {
		t.Fatalf(err.Error())
	}

	err = vendingMachine.requestItem()
	if err != nil {
		t.Fatalf(err.Error())
	}

	err = vendingMachine.insertMoney(10)
	if err != nil {
		t.Fatalf(err.Error())
	}

	err = vendingMachine.dispenseItem()
	if err != nil {
		t.Fatalf(err.Error())
	}
}
