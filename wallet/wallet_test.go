package wallet

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfRupeeConvertedIsOneDollar(t *testing.T) {
	got, _ := ConvertRupeesToDollars(Rupee{82, 47})
	want := Dollar{1, 0}
	assert.Equal(t, *got, want)
}

func TestToCheckIfICanCreditSomeMoneyToMyWallet(t *testing.T) {
	got, _ := CreditToWallet(Rupee{82, 47})
	want := MyWallet{Rupee{82, 47}, "Rupees"}
	assert.Equal(t, *got, want)
	fmt.Printf("Your Balance is %d.%d %s", *&got.balance.rupees, *&got.balance.paise, *&got.curr)
}

func TestToCheckIfICanCreditZeroMoneyToMyWallet(t *testing.T) {
	got, _ := CreditToWallet(Rupee{0, 0})
	assert.NotEqual(t, got, nil)

}

func TestToCheckIfICanCreditNegativeMoneyToMyWallet(t *testing.T) {
	got, _ := CreditToWallet(Rupee{-1, 47})
	assert.NotEqual(t, got, nil)
}

func TestDebitINRFromMyWallet(t *testing.T) {
	a := MyWallet{Rupee{82, 47}, "Rupees"}
	got, _ := DebitINRFromMyWallet(a, Rupee{82, 47})
	want := MyWallet{Rupee{0, 0}, "Rupees"}
	assert.Equal(t, *got, want)
	fmt.Printf("Balace is %d.%d", got.balance.rupees, got.balance.paise)
}

func TestDebitUSDFromMyWalletAtConstantExchangeRateOf_85INR_Is_1USD(t *testing.T) {
	a := MyWallet{Rupee{90, 0}, "Rupees"}
	got, _ := DebitUSDFromMyWallet(a, Dollar{1, 0})
	want := MyWallet{Rupee{5, 0}, "Rupees"}
	assert.Equal(t, *got, want)
	fmt.Printf("Balace is %d.%d", got.balance.rupees, got.balance.paise)
}
