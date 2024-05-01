package main

import (
	"time"
)

// To define these sum types:
//
//	type User
//	    = Anonymous PaymentMethod
//	    | Member String Time
//	    | Admin String
//
//	type PaymentMethod
//	    = CreditCard String Time
//	    | Paypal String
//
// we declare `User` as an interface
type User interface {
	Match(s UserVariants)
}

// and the variants as a struct
type UserVariants struct {
	Anonymous func(PaymentMethod)                 // preferably named, but can be anonymous
	Member    func(email string, since time.Time) // named for clarity
	Admin     func(email string)
}

type PaymentMethod interface {
	Match(s PaymentMethodVariants)
}

type PaymentMethodVariants struct {
	CreditCard func(number string, expiry time.Time)
	Paypal     func(email string)
}
