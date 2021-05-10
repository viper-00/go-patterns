package facade

import "testing"

func TestPattern(t *testing.T) {

	walletFacade := newWalletFacade("mywalletid", 666)

	if err := walletFacade.addMoneyToWallet("mywalletid", 666, 250); err != nil {
		t.Fatalf("Error: %s\n", err.Error())
	}

	if err := walletFacade.deductMoneyFromWallet("mywalletid", 666, 66); err != nil {
		t.Fatalf("Error: %s\n", err.Error())
	}
}
