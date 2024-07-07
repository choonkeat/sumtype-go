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

// CreditCard is a constructor function for PaymentMethod; see PaymentMethodVariants for all constructor functions of PaymentMethod
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

// Paypal is a constructor function for PaymentMethod; see PaymentMethodVariants for all constructor functions of PaymentMethod
func Paypal(emailArg string) PaymentMethod {
	return paypalPaymentMethodVariants{emailArg}
}

// PaymentMethodVariantsMap is parameter type of PaymentMethodMap function,
// like PaymentMethodVariants is parameter type of paymentMethod.Match method,
// but with methods that returns a value of generic type
type PaymentMethodVariantsMap[A any] struct {
	CreditCard func(numberArg string, expiryArg time.Time) A // when PaymentMethod value pattern matches to CreditCard, return different value
	Paypal     func(emailArg string) A                       // when PaymentMethod value pattern matches to Paypal, return different value
}

// PaymentMethodMap is like paymentMethod.Match method except it returns a value of generic type
// thus can transform a PaymentMethod value into anything else
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

// Anonymous is a constructor function for User; see UserVariants for all constructor functions of User
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

// Member is a constructor function for User; see UserVariants for all constructor functions of User
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

// Admin is a constructor function for User; see UserVariants for all constructor functions of User
func Admin(emailArg string) User {
	return adminUserVariants{emailArg}
}

// UserVariantsMap is parameter type of UserMap function,
// like UserVariants is parameter type of user.Match method,
// but with methods that returns a value of generic type
type UserVariantsMap[A any] struct {
	Anonymous func(arg0Arg PaymentMethod) A               // when User value pattern matches to Anonymous, return different value
	Member    func(emailArg string, sinceArg time.Time) A // when User value pattern matches to Member, return different value
	Admin     func(emailArg string) A                     // when User value pattern matches to Admin, return different value
}

// UserMap is like user.Match method except it returns a value of generic type
// thus can transform a User value into anything else
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
