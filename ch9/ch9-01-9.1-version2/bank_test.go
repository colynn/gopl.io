package bank_test

import (
	"fmt"
	"testing"

	bank "gopl.io/ch9/ch9-01-9.1-version2"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	// Alice
	go func() {
		bank.Deposit(200)
		fmt.Println("=", bank.Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		bank.Deposit(100)
		fmt.Println("=", bank.Balance())
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

	if got, want := bank.Balance(), 300; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	} else {
		fmt.Printf("current : %v\n", 300)
	}

	// Withdraw
	// Colynn
	result := bank.Withdraw(100)
	if got, want := result, true; got != want {
		t.Errorf("WithdRawResult = %v, want %v", got, want)
	} else {
		fmt.Printf("Withdraw %d success\n", 100)
	}

	// Alex
	Alexresult := bank.Withdraw(210)

	if got, want := Alexresult, false; got != want {
		t.Errorf("WithdRawResult = %v, want %v", got, want)
	} else {
		fmt.Printf("Withdraw %d failed\n", 210)
	}
}
