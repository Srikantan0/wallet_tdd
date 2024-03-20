package wallet2

import (
	"errors"
)

var errNegativeOrZeroCredit = errors.New("value is greater than 1 USD")
var errDebitMoreThanBalanceOrNegativeDebit = errors.New("trying to debit more than the balance or negative money")

type token struct {
	amt  float64
	curr currency
}

type Wallet struct {
	balance token
}

func ConvertOneTypeToAnother(c1, c2 currency) currency {
	return c1 / c2
}

func (w *Wallet) Balance(c currency) float64 {
	return w.balance.amt * float64(ConvertOneTypeToAnother(w.balance.curr, c))
}

func (w *Wallet) CreditToAccount(t token) error {

	if t.amt <= 0 {
		return errNegativeOrZeroCredit
	}
	w.balance.amt += t.amt * float64(ConvertOneTypeToAnother(t.curr, w.balance.curr))
	return nil
}

func (w *Wallet) DebitFromAccount(t token) error {
	if t.amt > w.balance.amt || t.amt <= 0 {
		return errDebitMoreThanBalanceOrNegativeDebit
	}
	w.balance.amt -= t.amt * float64(ConvertOneTypeToAnother(t.curr, w.balance.curr))
	return nil
}
