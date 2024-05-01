// Generated code by github.com/choonkeat/sumtype-go
package main

import (
	"time"
)

// CreditCard
type creditCardPaymentMethodVariants struct {
	number string
	expiry time.Time
}

func (s creditCardPaymentMethodVariants) Match(variants PaymentMethodVariants) {
	variants.CreditCard(s.number, s.expiry)
}

func CreditCard(numberArg string, expiryArg time.Time) PaymentMethod {
	return creditCardPaymentMethodVariants{numberArg, expiryArg}
}

// Paypal
type paypalPaymentMethodVariants struct {
	email string
}

func (s paypalPaymentMethodVariants) Match(variants PaymentMethodVariants) {
	variants.Paypal(s.email)
}

func Paypal(emailArg string) PaymentMethod {
	return paypalPaymentMethodVariants{emailArg}
}

// Anonymous
type anonymousUserVariants struct {
	arg0 PaymentMethod
}

func (s anonymousUserVariants) Match(variants UserVariants) {
	variants.Anonymous(s.arg0)
}

func Anonymous(arg0Arg PaymentMethod) User {
	return anonymousUserVariants{arg0Arg}
}

// Member
type memberUserVariants struct {
	email string
	since time.Time
}

func (s memberUserVariants) Match(variants UserVariants) {
	variants.Member(s.email, s.since)
}

func Member(emailArg string, sinceArg time.Time) User {
	return memberUserVariants{emailArg, sinceArg}
}

// Admin
type adminUserVariants struct {
	email string
}

func (s adminUserVariants) Match(variants UserVariants) {
	variants.Admin(s.email)
}

func Admin(emailArg string) User {
	return adminUserVariants{emailArg}
}
