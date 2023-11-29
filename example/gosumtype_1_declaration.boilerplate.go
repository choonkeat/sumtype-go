// Generated code by github.com/choonkeat/sumtype-go
package main

import (
	"time"
)

// Anonymous
type anonymousUserScenarios struct {
}

func (s anonymousUserScenarios) Switch(scenarios UserScenarios) {
	scenarios.Anonymous()
}

func Anonymous() User {
	return anonymousUserScenarios{}
}

// Member
type memberUserScenarios struct {
	email string
	since time.Time
}

func (s memberUserScenarios) Switch(scenarios UserScenarios) {
	scenarios.Member(s.email, s.since)
}

func Member(emailArg string, sinceArg time.Time) User {
	return memberUserScenarios{emailArg, sinceArg}
}

// Admin
type adminUserScenarios struct {
	email string
}

func (s adminUserScenarios) Switch(scenarios UserScenarios) {
	scenarios.Admin(s.email)
}

func Admin(emailArg string) User {
	return adminUserScenarios{emailArg}
}
