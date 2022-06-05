package main

import "fmt"

type Purchase struct{}

func NewPurchase() Purchase {
	return Purchase{}
}

func (p Purchase) Process(paymentMethod string) error {
	// you can add logic to query and validate the order

	switch paymentMethod {
	case "PayPal":
		if err := p.payWithPayPal(); err != nil {
			return err
		}
	case "CreditCard":
		if err := p.payWithCreditCard(); err != nil {
			return err
		}
	case "Bank":
		if err := p.payWithBank(); err != nil {
			return err
		}
	}

	// here you may want to create an invoice, send a notification, etc

	fmt.Printf("Purchase with %s completed\n", paymentMethod)

	return nil
}

func (p Purchase) payWithPayPal() error {
	// Here you'll add the logic to pay with PayPal

	fmt.Println("Processing purchase with PayPal...")

	return nil
}

func (p Purchase) payWithCreditCard() error {
	// Here you'll add the logic to pay with CreditCard

	fmt.Println("Processing purchase with CreditCard...")

	return nil
}

func (p Purchase) payWithBank() error {
	// Here you'll add the logic to pay with Bank

	fmt.Println("Processing purchase with Bank...")

	return nil
}
