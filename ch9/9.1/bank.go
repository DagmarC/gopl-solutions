// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 261.
//!+

// Package bank provides a concurrency-safe bank with one account.
package bank

var deposits = make(chan int)    // send amount to deposit
var balances = make(chan int)    // receive balance
var withdraws = make(chan int)   // send amount to withdraw
var success = make(chan bool, 1) // inform about the success of withdrawal

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	withdraws <- amount
	return <-success // Wait for the success msg from the monitor gor
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case amount := <-withdraws:
			if balance < amount {
				success <- false // insufficient amount to withdraw
			} else {
				balance -= amount
				success <- true
			}
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

//!-
