package wallet2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckIf82dot47RupeeIsOneDollar(t *testing.T) {
	x := 1 * ConvertOneTypeToAnother(dollar, rupee)
	y := 82.470000
	assert.Equal(t, float64(x), y)
}

func TestIfICanCreditMoneyToMyWallet(t *testing.T) {
	w := Wallet{token{2.0, dollar}}
	err := w.CreditToAccount(token{2.0, dollar})
	assert.Nil(t, err)
}

func TestIfICanDebitMoneyFromMyWallet(t *testing.T) {
	w := Wallet{token{2.0, dollar}}
	err := w.DebitFromAccount(token{1.0, dollar})
	assert.Nil(t, err)
}

func TestBalance(t *testing.T) {
	w := Wallet{token{3.0, dollar}}
	got := w.Balance(euro)
	want := 2.4741
	assert.Equal(t, got, want)
}

func TestCredit50INRThen1USDHavingBalance132dot47INR(t *testing.T) {
	w := Wallet{token{50.00, rupee}}
	_ = w.CreditToAccount(token{1.0, dollar})
	got := w.Balance(rupee)
	want := 132.47
	assert.Equal(t, got, want)

}

func TestCrediting82INRFollowedBy1USDThen164INRToCheckBalanceIs4USD(t *testing.T) {
	w := Wallet{token{82.47, rupee}}
	_ = w.CreditToAccount(token{1.0, dollar})
	_ = w.CreditToAccount(token{164.94, rupee})
	got := w.Balance(dollar)
	want := 4.0
	assert.Equal(t, got, want)
}

func TestCredit50INRThen1USDThen1EuroHavingBalance(t *testing.T) {
	w := Wallet{token{50.00, rupee}}
	_ = w.CreditToAccount(token{1.0, dollar})
	_ = w.CreditToAccount(token{1.0, euro})
	got := w.Balance(dollar)
	want := 2.8188432157148053
	assert.Equal(t, got, want)

}
