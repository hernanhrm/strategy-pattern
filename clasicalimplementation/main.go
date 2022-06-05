package main

import (
	"fmt"
	"log"
)

func main() {
	// unlike the Strategy Pattern implementation
	// here we don't have to initialize our strategies
	// because all the logic is inside the Purchase "Class"
	purchase := NewPurchase()

	fmt.Println("Enter the payment method you want to use (PayPal, CreditCard or Bank: ")
	var paymentMethod string
	fmt.Scanln(&paymentMethod)

	if err := purchase.Process(paymentMethod); err != nil {
		log.Fatalln(err)
	}
}
