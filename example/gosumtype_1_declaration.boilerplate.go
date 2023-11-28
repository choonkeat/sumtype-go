// Generated code by github.com/choonkeat/sumtype-go
package main

import (
	"time"
)

// Anonymous
type anonymous struct {
}

func (s anonymous) Switch(scenarios UserScenarios) {
	scenarios.Anonymous()
}

func Anonymous() User {
	return anonymous{}
}

// Member
type member struct {
	email string
	since time.Time
}

func (s member) Switch(scenarios UserScenarios) {
	scenarios.Member(s.email, s.since)
}

func Member(email string, since time.Time) User {
	return member{email, since}
}

// Admin
type admin struct {
	email string
}

func (s admin) Switch(scenarios UserScenarios) {
	scenarios.Admin(s.email)
}

func Admin(email string) User {
	return admin{email}
}
