package main

import "fmt"

type Bank struct{}

func NewBank() Bank {
	return Bank{}
}

func (b Bank) Pay() error {
	// Here you'll add the logic to pay with a bank

	fmt.Println("Processing purchase with Bank...")

	return nil
}
