package main

import (
	"time"
)

// To define this sum type:
//
//	type User
//	    = Anonymous
//	    | Member String Time
//	    | Admin String
//
// Ideally, we just code something like this and the
// rest of the boiler plate can be generated
type User interface {
	Switch(s UserScenarios)
}

type UserScenarios struct {
	Anonymous func()
	Member    func(email string, since time.Time)
	Admin     func(email string)
}
