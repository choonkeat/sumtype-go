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
// we declare `User` as an interface
type User interface {
	Switch(s UserScenarios)
}

// and the variants as a struct
type UserScenarios struct {
	Anonymous func()
	Member    func(email string, since time.Time)
	Admin     func(email string)
}
