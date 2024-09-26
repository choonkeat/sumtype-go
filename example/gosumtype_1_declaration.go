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
// if we declare the `UserVariants` as a struct, we'll get the `User` type.
type UserVariants struct {
	Anonymous func(PaymentMethod)                 // preferably named, but can be anonymous
	Member    func(email string, since time.Time) // named for clarity
	Admin     func()
}

type PaymentMethodVariants struct {
	CreditCard func(number string, expiry time.Time)
	Paypal     func(email string)
}
