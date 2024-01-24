package main

import "testing"

func TestWallet(t *testing.T) {
	assertWallets := func(t testing.TB, w Wallet, want Bitcoin) {
		t.Helper()
		got := w.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertNoError := func(t testing.TB, err error) {
		t.Helper()

		if err != nil {
			t.Fatal("Wanted no error but got one")
		}
	}

	assertError := func(t testing.TB, err error, want error) {
		t.Helper()
		if err == nil {
			t.Fatal("wanted an error but didn't get one")
		}

		if err != want {
			t.Errorf("got %q, want %q", err, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertWallets(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw with funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		assertNoError(t, err)
		assertWallets(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrorInsufficientFunds)
		assertWallets(t, wallet, startingBalance)
	})
}
