// Generated code by github.com/choonkeat/sumtype-go
package main

import (
	"time"
)

// Anonymous
type anonymousUserVariants struct {
	payment PaymentMethod
}

func (s anonymousUserVariants) Match(Variants UserVariants) {
	Variants.Anonymous(s.payment)
}

func Anonymous(paymentArg PaymentMethod) User {
	return anonymousUserVariants{paymentArg}
}

// Member
type memberUserVariants struct {
	email string
	since time.Time
}

func (s memberUserVariants) Match(Variants UserVariants) {
	Variants.Member(s.email, s.since)
}

func Member(emailArg string, sinceArg time.Time) User {
	return memberUserVariants{emailArg, sinceArg}
}

// Admin
type adminUserVariants struct {
	email string
}

func (s adminUserVariants) Match(Variants UserVariants) {
	Variants.Admin(s.email)
}

func Admin(emailArg string) User {
	return adminUserVariants{emailArg}
}

// CreditCard
type creditCardPaymentMethodVariants struct {
	number string
	expiry time.Time
}

func (s creditCardPaymentMethodVariants) Match(Variants PaymentMethodVariants) {
	Variants.CreditCard(s.number, s.expiry)
}

func CreditCard(numberArg string, expiryArg time.Time) PaymentMethod {
	return creditCardPaymentMethodVariants{numberArg, expiryArg}
}

// Paypal
type paypalPaymentMethodVariants struct {
	email string
}

func (s paypalPaymentMethodVariants) Match(Variants PaymentMethodVariants) {
	Variants.Paypal(s.email)
}

func Paypal(emailArg string) PaymentMethod {
	return paypalPaymentMethodVariants{emailArg}
}
