package main

type PaymentMethodStrategy interface {
	Pay() error
}
