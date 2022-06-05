package main

import "fmt"

type CreditCard struct{}

func NewCreditCard() CreditCard {
	return CreditCard{}
}

func (c CreditCard) Pay() error {
	// Here you'll add the logic to pay with credit card

	fmt.Println("Processing purchase with CreditCard...")

	return nil
}
