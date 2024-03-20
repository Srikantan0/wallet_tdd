package wallet

import (
	"errors"
)

var RupeesGreaterThanOneDollarError = errors.New("Value is greater than 1 USD")
var RupeesLesserThanOneDollarError = errors.New("Value is lesser than 1 USD")
var NegativeValueAddedToWallet = errors.New("Value added to wallet is negative")
var ZeroAddedToWallet = errors.New("Value added to wallet is zero")
var DebitingMoreThanBalanceError = errors.New("Debitign more than balance in account")

var mywallet = MyWallet{Rupee{0, 0}, "Rupee"}

func ConvertRupeesToDollars(r Rupee) (*Dollar, error) {
	ru, p := r.rupees, r.paise
	if ru == 82 && p == 47 {
		return &Dollar{1, 0}, nil
	} else if ru < 82 || p < 47 {
		return nil, RupeesLesserThanOneDollarError
	}
	return nil, RupeesGreaterThanOneDollarError
}

func CreditToWallet(r Rupee) (*MyWallet, error) {
	if r.rupees < 0 {
		return nil, NegativeValueAddedToWallet
	} else if r.rupees == 0 && r.paise == 0 {
		return nil, ZeroAddedToWallet
	}
	mywallet.balance.rupees += r.rupees
	mywallet.balance.paise += r.paise
	mywallet.curr = "Rupees"
	return &mywallet, nil
}

func DebitINRFromMyWallet(a MyWallet, r Rupee) (*MyWallet, error) {
	if a.balance.rupees < r.rupees || a.balance.paise < r.paise {
		return nil, DebitingMoreThanBalanceError
	}
	a.balance.rupees -= r.rupees
	a.balance.paise -= r.paise
	return &a, nil
}

func DebitUSDFromMyWallet(a MyWallet, d Dollar) (*MyWallet, error) {
	var r Rupee
	r.rupees = d.dollar * 85
	r.paise = d.cent * 85
	r.paise /= 100
	if a.balance.rupees < d.dollar || a.balance.paise < d.cent {
		return nil, DebitingMoreThanBalanceError
	}
	a.balance.rupees -= r.rupees
	a.balance.paise -= r.paise
	return &a, nil
}
