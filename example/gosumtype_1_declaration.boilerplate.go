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

func (creditCardInstance creditCardPaymentMethodVariants) Match(creditCardVariants PaymentMethodVariants) {
	creditCardVariants.CreditCard(creditCardInstance.Number, creditCardInstance.Expiry)
}

// CreditCard is a constructor function for PaymentMethod; see PaymentMethodVariants for all constructor functions of PaymentMethod
func CreditCard(numberArg string, expiryArg time.Time) PaymentMethod {
	return PaymentMethod{creditCardPaymentMethodVariants{numberArg, expiryArg}}
}

// Paypal
type paypalPaymentMethodVariants struct {
	Email string
}

func (paypalInstance paypalPaymentMethodVariants) Match(paypalVariants PaymentMethodVariants) {
	paypalVariants.Paypal(paypalInstance.Email)
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
func PaymentMethodMap[A any](paymentMethodValue PaymentMethod, paymentMethodVariants PaymentMethodVariantsMap[A]) A {
	var paymentMethodTemp A
	paymentMethodValue.Match(PaymentMethodVariants{
		CreditCard: func(numberArg string, expiryArg time.Time) {
			paymentMethodTemp = paymentMethodVariants.CreditCard(numberArg, expiryArg)
		},
		Paypal: func(emailArg string) {
			paymentMethodTemp = paymentMethodVariants.Paypal(emailArg)
		},
	})
	return paymentMethodTemp
}

// PaymentMethod = CreditCard | Paypal
type PaymentMethod struct {
	paymentMethod paymentMethod
}

// paymentMethod is the interface for PaymentMethodVariants
type paymentMethod interface {
	Match(variants PaymentMethodVariants)
}

func (paymentMethodInstance PaymentMethod) Match(paymentMethodVariants PaymentMethodVariants) {
	paymentMethodInstance.paymentMethod.Match(paymentMethodVariants)
}
func (paymentMethodInstance PaymentMethod) MarshalJSON() (paymentMethodData []byte, paymentMethodErr error) {
	paymentMethodInstance.paymentMethod.Match(PaymentMethodVariants{
		CreditCard: func(numberArg string, expiryArg time.Time) {
			paymentMethodData, paymentMethodErr = json.Marshal([]any{
				"CreditCard",
				creditCardPaymentMethodVariants{
					Number: numberArg,
					Expiry: expiryArg,
				}})
		},
		Paypal: func(emailArg string) {
			paymentMethodData, paymentMethodErr = json.Marshal([]any{
				"Paypal",
				paypalPaymentMethodVariants{
					Email: emailArg,
				}})
		},
	})
	return paymentMethodData, paymentMethodErr
}
func (paymentMethodInstance *PaymentMethod) UnmarshalJSON(paymentMethodData []byte) error {
	// The expected format is ["TypeName", { ... data... }]
	var paymentMethodRaw []json.RawMessage
	if err := json.Unmarshal(paymentMethodData, &paymentMethodRaw); err != nil {
		return fmt.Errorf("expected an array with type and data, got error: %w", err)
	}
	if len(paymentMethodRaw) != 2 {
		return fmt.Errorf("expected array of two elements [type, data], got %d elements", len(paymentMethodRaw))
	}
	// Unmarshal the first element to get the type
	var paymentMethodVariantName string
	if err := json.Unmarshal(paymentMethodRaw[0], &paymentMethodVariantName); err != nil {
		return fmt.Errorf("failed to unmarshal type name: %w", err)
	}
	switch paymentMethodVariantName {
	case "CreditCard":
		var paymentMethodTemp creditCardPaymentMethodVariants
		if err := json.Unmarshal(paymentMethodRaw[1], &paymentMethodTemp); err != nil {
			return fmt.Errorf("failed to unmarshal data: %w", err)
		}
		paymentMethodInstance.paymentMethod = creditCardPaymentMethodVariants{
			Number: paymentMethodTemp.Number,
			Expiry: paymentMethodTemp.Expiry,
		}
	case "Paypal":
		var paymentMethodTemp paypalPaymentMethodVariants
		if err := json.Unmarshal(paymentMethodRaw[1], &paymentMethodTemp); err != nil {
			return fmt.Errorf("failed to unmarshal data: %w", err)
		}
		paymentMethodInstance.paymentMethod = paypalPaymentMethodVariants{
			Email: paymentMethodTemp.Email,
		}
	default:
		return fmt.Errorf("unknown type %q", paymentMethodVariantName)
	}
	return nil
}

// Anonymous
type anonymousUserVariants struct {
	Arg0 PaymentMethod
}

func (anonymousInstance anonymousUserVariants) Match(anonymousVariants UserVariants) {
	anonymousVariants.Anonymous(anonymousInstance.Arg0)
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

func (memberInstance memberUserVariants) Match(memberVariants UserVariants) {
	memberVariants.Member(memberInstance.Email, memberInstance.Since)
}

// Member is a constructor function for User; see UserVariants for all constructor functions of User
func Member(emailArg string, sinceArg time.Time) User {
	return User{memberUserVariants{emailArg, sinceArg}}
}

// Admin
type adminUserVariants struct {
	Email string
}

func (adminInstance adminUserVariants) Match(adminVariants UserVariants) {
	adminVariants.Admin(adminInstance.Email)
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
func UserMap[A any](userValue User, userVariants UserVariantsMap[A]) A {
	var userTemp A
	userValue.Match(UserVariants{
		Anonymous: func(arg0Arg PaymentMethod) {
			userTemp = userVariants.Anonymous(arg0Arg)
		},
		Member: func(emailArg string, sinceArg time.Time) {
			userTemp = userVariants.Member(emailArg, sinceArg)
		},
		Admin: func(emailArg string) {
			userTemp = userVariants.Admin(emailArg)
		},
	})
	return userTemp
}

// User = Anonymous | Member | Admin
type User struct {
	user user
}

// user is the interface for UserVariants
type user interface {
	Match(variants UserVariants)
}

func (userInstance User) Match(userVariants UserVariants) {
	userInstance.user.Match(userVariants)
}
func (userInstance User) MarshalJSON() (userData []byte, userErr error) {
	userInstance.user.Match(UserVariants{
		Anonymous: func(arg0Arg PaymentMethod) {
			userData, userErr = json.Marshal([]any{
				"Anonymous",
				anonymousUserVariants{
					Arg0: arg0Arg,
				}})
		},
		Member: func(emailArg string, sinceArg time.Time) {
			userData, userErr = json.Marshal([]any{
				"Member",
				memberUserVariants{
					Email: emailArg,
					Since: sinceArg,
				}})
		},
		Admin: func(emailArg string) {
			userData, userErr = json.Marshal([]any{
				"Admin",
				adminUserVariants{
					Email: emailArg,
				}})
		},
	})
	return userData, userErr
}
func (userInstance *User) UnmarshalJSON(userData []byte) error {
	// The expected format is ["TypeName", { ... data... }]
	var userRaw []json.RawMessage
	if err := json.Unmarshal(userData, &userRaw); err != nil {
		return fmt.Errorf("expected an array with type and data, got error: %w", err)
	}
	if len(userRaw) != 2 {
		return fmt.Errorf("expected array of two elements [type, data], got %d elements", len(userRaw))
	}
	// Unmarshal the first element to get the type
	var userVariantName string
	if err := json.Unmarshal(userRaw[0], &userVariantName); err != nil {
		return fmt.Errorf("failed to unmarshal type name: %w", err)
	}
	switch userVariantName {
	case "Anonymous":
		var userTemp anonymousUserVariants
		if err := json.Unmarshal(userRaw[1], &userTemp); err != nil {
			return fmt.Errorf("failed to unmarshal data: %w", err)
		}
		userInstance.user = anonymousUserVariants{
			Arg0: userTemp.Arg0,
		}
	case "Member":
		var userTemp memberUserVariants
		if err := json.Unmarshal(userRaw[1], &userTemp); err != nil {
			return fmt.Errorf("failed to unmarshal data: %w", err)
		}
		userInstance.user = memberUserVariants{
			Email: userTemp.Email,
			Since: userTemp.Since,
		}
	case "Admin":
		var userTemp adminUserVariants
		if err := json.Unmarshal(userRaw[1], &userTemp); err != nil {
			return fmt.Errorf("failed to unmarshal data: %w", err)
		}
		userInstance.user = adminUserVariants{
			Email: userTemp.Email,
		}
	default:
		return fmt.Errorf("unknown type %q", userVariantName)
	}
	return nil
}
