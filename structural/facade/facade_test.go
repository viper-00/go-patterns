package facade

import "testing"

func TestPattern(t *testing.T) {
	var walletID string = "mywalletid"
	var securityCode int = 666

	walletFacade := newWalletFacade(walletID, securityCode)
	if err := walletFacade.addMoneyToWallet(walletID, securityCode, 250); err != nil {
		t.Errorf(err.Error())
	}

	if err := walletFacade.deductMoneyFromWallet(walletID, securityCode, 66); err != nil {
		t.Errorf(err.Error())
	}
}
