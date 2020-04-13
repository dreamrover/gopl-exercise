// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package bank_test

import (
	"fmt"
	"testing"

	"exercise/9.1"
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

	if got, want := bank.Balance(), 300; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}

	failures := 0

	// Carl
	go func() {
		if !bank.Withdraw(200) {
			failures++
		}
		done <- struct{}{}
	}()

	// David
	go func() {
		if !bank.Withdraw(200) {
			failures++
		}
		done <- struct{}{}
	}()

	// Ellen
	go func() {
		if !bank.Withdraw(100) {
			failures++
		}
		done <- struct{}{}
	}()

	<-done
	<-done
	<-done

	if got, want := bank.Balance(), 0; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}

	if failures != 1 {
		t.Errorf("failures = %d, want %d", failures, 1)
	}
}
