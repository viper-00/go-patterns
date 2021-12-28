package facade

import "fmt"

// wiki: https://en.wikipedia.org/wiki/Facade_pattern
//
// Uses one type as an API to a number of others.

type walletFacade struct {
	account      *account
	wallet       *wallet
	securityCode *securityCode
	notification *notification
	ledger       *ledger
}

func newWalletFacade(accountID string, code int) *walletFacade {
	fmt.Println("Start to create an account")
	walletFacade := &walletFacade{
		account:      newAccount(accountID),
		wallet:       newWallet(),
		securityCode: newSecurityCode(code),
		notification: &notification{},
		ledger:       &ledger{},
	}
	fmt.Println("Account already created")
	return walletFacade
}

func (w *walletFacade) addMoneyToWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Start to add some money to wallet")
	if err := w.account.checkAccount(accountID); err != nil {
		return err
	}
	if err := w.securityCode.checkCode(securityCode); err != nil {
		return err
	}
	w.wallet.creditBalance(amount)
	w.notification.sendWalletCreditNotification()
	w.ledger.makeEntry(accountID, "credit", amount)
	return nil
}

func (w *walletFacade) deductMoneyFromWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Start to debit money from wallet")

	if err := w.account.checkAccount(accountID); err != nil {
		return err
	}

	if err := w.securityCode.checkCode(securityCode); err != nil {
		return err
	}

	if err := w.wallet.debitBalance(amount); err != nil {
		return err
	}

	w.notification.sendWalletDebitNotification()
	w.ledger.makeEntry(accountID, "debit", amount)
	return nil
}

type account struct {
	name string
}

func newAccount(accountName string) *account {
	return &account{name: accountName}
}

func (a *account) checkAccount(accountName string) error {
	if a.name != accountName {
		return fmt.Errorf("name of account is incorrect")
	}
	fmt.Println("Account verified")
	return nil
}

type securityCode struct {
	code int
}

func newSecurityCode(code int) *securityCode {
	return &securityCode{code: code}
}

func (s *securityCode) checkCode(code int) error {
	if s.code != code {
		return fmt.Errorf("security code is incorrect")
	}
	fmt.Println("SecurityCode verified")
	return nil
}

type wallet struct {
	balance int
}

func newWallet() *wallet {
	return &wallet{balance: 0}
}

func (w *wallet) creditBalance(amount int) {
	w.balance += amount
	fmt.Println("Wallet balance added successfully")
}

func (w *wallet) debitBalance(amount int) error {
	if w.balance < amount {
		return fmt.Errorf("balance is not sufficient")
	}

	fmt.Println("Wallet balance is sufficient")
	w.balance -= amount
	return nil
}

type ledger struct{}

func (l *ledger) makeEntry(accountID string, payType string, amount int) {
	fmt.Printf("Make a ledger entry for accountID %s with payType %s for amount %d\n", accountID, payType, amount)
}

type notification struct{}

func (n *notification) sendWalletCreditNotification() {
	fmt.Println("Send a wallet credit notification")
}

func (n *notification) sendWalletDebitNotification() {
	fmt.Println("Send a wallet debit notification")
}
