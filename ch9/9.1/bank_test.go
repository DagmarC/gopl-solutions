// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package bank_test

import (
	"fmt"
	"testing"

	bank "github.com/DagmarC/gopl-solutions/ch9/9.1"
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
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

	// Dan
	go func() {
		ok := bank.Withdraw(123)
		if !ok {
			t.Errorf("Withdraw = %d should succeed, balance = %d", 123, bank.Balance())
		}
		done <- struct{}{}
	}()

	// Bob
	go func() {
		ok := bank.Withdraw(140)
		if !ok {
			t.Errorf("Withdraw = %d should succeed, balance = %d", 140, bank.Balance())
		}
		done <- struct{}{}
	}()

	// Alice
	go func() {
		ok := bank.Withdraw(500)
		if ok {
			t.Errorf("Withdraw = %d should NOT succeed, balance = %d", 500, bank.Balance())
		}
		done <- struct{}{}
	}()

	// Wait for withdrawals.
	<-done
	<-done
	<-done

	if got, want := bank.Balance(), 200+100-123-140; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
