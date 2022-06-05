package main

import "fmt"

// Purchase is our goal/context which maintains a reference to all the
// possible strategies that the client can use
type Purchase struct {
	// we create a map to save all the strategies implementations
	// so in runtime, when a client execute the Process method
	// he can choose the strategy to use
	paymentMethodStrategies map[string]PaymentMethodStrategy
}

func NewPurchase() Purchase {
	return Purchase{paymentMethodStrategies: make(map[string]PaymentMethodStrategy)}
}

// RegisterStrategy will let us register all the possible strategies that
// the client can use
func (p *Purchase) RegisterStrategy(name string, strategy PaymentMethodStrategy) {
	p.paymentMethodStrategies[name] = strategy
}

// Process processes the purchase by doing the necessary validations
// and calling the Pay method of the selected strategy by the client
func (p Purchase) Process(paymentMethod string) error {
	// you can add logic to query and validate the order

	// in runtime, we chose the payment method strategy
	paymentMethodStrategy := p.paymentMethodStrategies[paymentMethod]

	// after we got the strategy implementation we just use it by calling the Pay method
	if err := paymentMethodStrategy.Pay(); err != nil {
		return fmt.Errorf("purchase.Process.paymentMethodStrategy.Pay(): %w", err)
	}

	// here you may want to create an invoice, send a notification, etc

	fmt.Printf("Purchase with %s completed\n", paymentMethod)

	return nil
}
