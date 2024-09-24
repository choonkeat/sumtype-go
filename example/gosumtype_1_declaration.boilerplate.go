// Generated code by github.com/choonkeat/sumtype-go
package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// CreditCard
type creditCardPaymentMethodVariants struct {
	Number string
	Expiry time.Time
}

func (s creditCardPaymentMethodVariants) Match(variants PaymentMethodVariants) {
	variants.CreditCard(s.Number, s.Expiry)
}

// CreditCard is a constructor function for PaymentMethod; see PaymentMethodVariants for all constructor functions of PaymentMethod
func CreditCard(numberArg string, expiryArg time.Time) PaymentMethod {
	return PaymentMethod{creditCardPaymentMethodVariants{numberArg, expiryArg}}
}

// Paypal
type paypalPaymentMethodVariants struct {
	Email string
}

func (s paypalPaymentMethodVariants) Match(variants PaymentMethodVariants) {
	variants.Paypal(s.Email)
}

// Paypal is a constructor function for PaymentMethod; see PaymentMethodVariants for all constructor functions of PaymentMethod
func Paypal(emailArg string) PaymentMethod {
	return PaymentMethod{paypalPaymentMethodVariants{emailArg}}
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

// PaymentMethod = CreditCard | Paypal
type PaymentMethod struct {
	paymentMethod paymentMethod
}

// paymentMethod is the interface for PaymentMethodVariants
type paymentMethod interface {
	Match(variants PaymentMethodVariants)
}

func (s PaymentMethod) Match(variants PaymentMethodVariants) {
	s.paymentMethod.Match(variants)
}
func (s PaymentMethod) MarshalJSON() (data []byte, err error) {
	s.paymentMethod.Match(PaymentMethodVariants{
		CreditCard: func(numberArg string, expiryArg time.Time) {
			data, err = json.Marshal([]any{
				"CreditCard",
				creditCardPaymentMethodVariants{
					Number: numberArg,
					Expiry: expiryArg,
				}})
		},
		Paypal: func(emailArg string) {
			data, err = json.Marshal([]any{
				"Paypal",
				paypalPaymentMethodVariants{
					Email: emailArg,
				}})
		},
	})
	return data, err
}
func (s *PaymentMethod) UnmarshalJSON(data []byte) error {
	// The expected format is ["TypeName", { ... data... }]
	var raw []json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return fmt.Errorf("expected an array with type and data, got error: %w", err)
	}
	if len(raw) != 2 {
		return fmt.Errorf("expected array of two elements [type, data], got %d elements", len(raw))
	}
	// Unmarshal the first element to get the type
	var typeName string
	if err := json.Unmarshal(raw[0], &typeName); err != nil {
		return fmt.Errorf("failed to unmarshal type name: %w", err)
	}
	switch typeName {
	case "CreditCard":
		var temp creditCardPaymentMethodVariants
		if err := json.Unmarshal(raw[1], &temp); err != nil {
			return fmt.Errorf("failed to unmarshal data: %w", err)
		}
		s.paymentMethod = creditCardPaymentMethodVariants{
			Number: temp.Number,
			Expiry: temp.Expiry,
		}
	case "Paypal":
		var temp paypalPaymentMethodVariants
		if err := json.Unmarshal(raw[1], &temp); err != nil {
			return fmt.Errorf("failed to unmarshal data: %w", err)
		}
		s.paymentMethod = paypalPaymentMethodVariants{
			Email: temp.Email,
		}
	default:
		return fmt.Errorf("unknown type %q", typeName)
	}
	return nil
}

// Anonymous
type anonymousUserVariants struct {
	Arg0 PaymentMethod
}

func (s anonymousUserVariants) Match(variants UserVariants) {
	variants.Anonymous(s.Arg0)
}

// Anonymous is a constructor function for User; see UserVariants for all constructor functions of User
func Anonymous(arg0Arg PaymentMethod) User {
	return User{anonymousUserVariants{arg0Arg}}
}

// Member
type memberUserVariants struct {
	Email string
	Since time.Time
}

func (s memberUserVariants) Match(variants UserVariants) {
	variants.Member(s.Email, s.Since)
}

// Member is a constructor function for User; see UserVariants for all constructor functions of User
func Member(emailArg string, sinceArg time.Time) User {
	return User{memberUserVariants{emailArg, sinceArg}}
}

// Admin
type adminUserVariants struct {
	Email string
}

func (s adminUserVariants) Match(variants UserVariants) {
	variants.Admin(s.Email)
}

// Admin is a constructor function for User; see UserVariants for all constructor functions of User
func Admin(emailArg string) User {
	return User{adminUserVariants{emailArg}}
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

// User = Anonymous | Member | Admin
type User struct {
	user user
}

// user is the interface for UserVariants
type user interface {
	Match(variants UserVariants)
}

func (s User) Match(variants UserVariants) {
	s.user.Match(variants)
}
func (s User) MarshalJSON() (data []byte, err error) {
	s.user.Match(UserVariants{
		Anonymous: func(arg0Arg PaymentMethod) {
			data, err = json.Marshal([]any{
				"Anonymous",
				anonymousUserVariants{
					Arg0: arg0Arg,
				}})
		},
		Member: func(emailArg string, sinceArg time.Time) {
			data, err = json.Marshal([]any{
				"Member",
				memberUserVariants{
					Email: emailArg,
					Since: sinceArg,
				}})
		},
		Admin: func(emailArg string) {
			data, err = json.Marshal([]any{
				"Admin",
				adminUserVariants{
					Email: emailArg,
				}})
		},
	})
	return data, err
}
func (s *User) UnmarshalJSON(data []byte) error {
	// The expected format is ["TypeName", { ... data... }]
	var raw []json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return fmt.Errorf("expected an array with type and data, got error: %w", err)
	}
	if len(raw) != 2 {
		return fmt.Errorf("expected array of two elements [type, data], got %d elements", len(raw))
	}
	// Unmarshal the first element to get the type
	var typeName string
	if err := json.Unmarshal(raw[0], &typeName); err != nil {
		return fmt.Errorf("failed to unmarshal type name: %w", err)
	}
	switch typeName {
	case "Anonymous":
		var temp anonymousUserVariants
		if err := json.Unmarshal(raw[1], &temp); err != nil {
			return fmt.Errorf("failed to unmarshal data: %w", err)
		}
		s.user = anonymousUserVariants{
			Arg0: temp.Arg0,
		}
	case "Member":
		var temp memberUserVariants
		if err := json.Unmarshal(raw[1], &temp); err != nil {
			return fmt.Errorf("failed to unmarshal data: %w", err)
		}
		s.user = memberUserVariants{
			Email: temp.Email,
			Since: temp.Since,
		}
	case "Admin":
		var temp adminUserVariants
		if err := json.Unmarshal(raw[1], &temp); err != nil {
			return fmt.Errorf("failed to unmarshal data: %w", err)
		}
		s.user = adminUserVariants{
			Email: temp.Email,
		}
	default:
		return fmt.Errorf("unknown type %q", typeName)
	}
	return nil
}
