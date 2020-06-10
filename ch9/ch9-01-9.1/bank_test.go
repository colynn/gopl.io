package bank_test

import (
	"fmt"
	"testing"

	bank "gopl.io/ch9/ch9-01-9.1"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})
	withdRawResult := bank.Withdraws{}
	withdRawResult01 := make(chan bank.Withdraws)

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
	}

	// Withdraw
	// Colynn
	result := bank.Withdraw(100)
	withdRawResult = <-result
	if got, want := withdRawResult.Success, true; got != want {
		t.Errorf("WithdRawResult = %v, want %v", got, want)
	} else {
		fmt.Printf("Withdraw %d success\n", withdRawResult.Amount)
	}

	// Alex
	go func() {
		result := bank.Withdraw(210)
		item := <-result
		withdRawResult01 <- item
	}()

	result01 := <-withdRawResult01

	if got, want := result01.Success, false; got != want {
		t.Errorf("WithdRawResult = %v, want %v", got, want)
	} else {
		fmt.Printf("Withdraw %d failed\n", result01.Amount)
	}
}
