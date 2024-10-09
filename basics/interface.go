package basics

import (
	"fmt"
)

// use interface - naming convention add er in end
type paymenter interface {
	Pay(amount float32)
}

type payment struct {
	// gateway stripe
	// gateway razorPay
	gateway paymenter
}

func (p payment) makePayment(amount float32) {
	// razorPaymentGw := razorPay{}
	// razorPaymentGw.Pay(amount)
	// stripePaymentGw := stripe{}
	// stripePaymentGw.Pay(amount)
	p.gateway.Pay(amount)
}

// type razorPay struct{}

// func (p razorPay) Pay(amount float32) {
// 	fmt.Println("making payment using razorPay")
// }

type stripe struct{}

func (s stripe) Pay(amount float32) {
	fmt.Println("making payment using stripe")
}

func Interfaces() {
	stripePaymentGw := stripe{}
	// razorPaymentGw := razorPay{}
	newPaymentGw := payment{
		gateway: stripePaymentGw,
		// gateway: razorPaymentGw,
	}
	newPaymentGw.makePayment(100)
	// razorPaymentGw := razorPay{}
	// stripePaymentGw := payment{}
	// stripePaymentGw.makePayment(100)
	// razorPaymentGw.Pay(10000)
}
