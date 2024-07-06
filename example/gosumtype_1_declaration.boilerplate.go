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

type PaymentMethodVariantsMap[A any] struct {
	CreditCard func(numberArg string, expiryArg time.Time) A
	Paypal     func(emailArg string) A
}

func PaymentMethodMap[A any](value PaymentMethod, variants PaymentMethodVariantsMap[A]) A {
	var result A
	value.Match(PaymentMethodVariants{
		CreditCard: func(numberArg string, expiryArg time.Time) {
			result = variants.CreditCard(numberArg, expiryArg)
		},
		Paypal: func(emailArg string) {
			result = variants.Paypal(emailArg)
		},
	})
	return result
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

type UserVariantsMap[A any] struct {
	Anonymous func(arg0Arg PaymentMethod) A
	Member    func(emailArg string, sinceArg time.Time) A
	Admin     func(emailArg string) A
}

func UserMap[A any](value User, variants UserVariantsMap[A]) A {
	var result A
	value.Match(UserVariants{
		Anonymous: func(arg0Arg PaymentMethod) {
			result = variants.Anonymous(arg0Arg)
		},
		Member: func(emailArg string, sinceArg time.Time) {
			result = variants.Member(emailArg, sinceArg)
		},
		Admin: func(emailArg string) {
			result = variants.Admin(emailArg)
		},
	})
	return result
}
