package main

import (
	"time"
)

// To define this sum type:
//
//	type User
//	    = Anonymous
//	    | Member String Time
//	    | Admin String
//
// we declare `User` as an interface
type User interface {
	Switch(s UserScenarios)
}

// and the variants as a struct
type UserScenarios struct {
	Anonymous func(payment PaymentMethod)
	Member    func(payment PaymentMethod, email string, since time.Time)
	Admin     func(payment PaymentMethod, email string)
}

type PaymentMethod interface {
	Switch(s PaymentMethodScenarios)
}

type PaymentMethodScenarios struct {
	CreditCard func(number string, expiry time.Time)
	Paypal     func(email string)
}
