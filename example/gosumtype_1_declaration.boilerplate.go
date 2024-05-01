// Generated code by github.com/choonkeat/sumtype-go
package main

import (
	"time"
)

// Anonymous
type anonymousUserScenarios struct {
	payment PaymentMethod
}

func (s anonymousUserScenarios) Switch(scenarios UserScenarios) {
	scenarios.Anonymous(s.payment)
}

func Anonymous(paymentArg PaymentMethod) User {
	return anonymousUserScenarios{paymentArg}
}

// Member
type memberUserScenarios struct {
	payment PaymentMethod
	email   string
	since   time.Time
}

func (s memberUserScenarios) Switch(scenarios UserScenarios) {
	scenarios.Member(s.payment, s.email, s.since)
}

func Member(paymentArg PaymentMethod, emailArg string, sinceArg time.Time) User {
	return memberUserScenarios{paymentArg, emailArg, sinceArg}
}

// Admin
type adminUserScenarios struct {
	payment PaymentMethod
	email   string
}

func (s adminUserScenarios) Switch(scenarios UserScenarios) {
	scenarios.Admin(s.payment, s.email)
}

func Admin(paymentArg PaymentMethod, emailArg string) User {
	return adminUserScenarios{paymentArg, emailArg}
}

// CreditCard
type creditCardPaymentMethodScenarios struct {
	number string
	expiry time.Time
}

func (s creditCardPaymentMethodScenarios) Switch(scenarios PaymentMethodScenarios) {
	scenarios.CreditCard(s.number, s.expiry)
}

func CreditCard(numberArg string, expiryArg time.Time) PaymentMethod {
	return creditCardPaymentMethodScenarios{numberArg, expiryArg}
}

// Paypal
type paypalPaymentMethodScenarios struct {
	email string
}

func (s paypalPaymentMethodScenarios) Switch(scenarios PaymentMethodScenarios) {
	scenarios.Paypal(s.email)
}

func Paypal(emailArg string) PaymentMethod {
	return paypalPaymentMethodScenarios{emailArg}
}
