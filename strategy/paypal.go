package main

import "fmt"

// PayPal implementation of the PaymentMethodStrategy interface
// with the necessary logic to pay with PayPal
type PayPal struct{}

func NewPaypal() PayPal {
	return PayPal{}
}

func (p PayPal) Pay() error {
	// Here you'll add the logic to pay with PayPal

	fmt.Println("Processing purchase with PayPal...")

	return nil
}
