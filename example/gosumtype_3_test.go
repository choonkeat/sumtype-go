package main

import (
	"log"
	"strconv"
	"testing"
	"time"
)

func TestUser(t *testing.T) {
	testCases := []struct {
		givenUser User
	}{
		{
			givenUser: Anonymous(),
		},
		{
			givenUser: Member("bob@example.com", time.Now()),
		},
		{
			givenUser: Admin("boss@example.com"),
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// using names are very helpful, but loses the exhaustive check at compile time
			// since Go happily set the undefined scenarios as function zero value: nil
			//
			// but we can use https://golangci-lint.run/usage/linters/#exhaustruct
			// to check at CI instead of suffering from zero value at runtime
			tc.givenUser.Switch(UserScenarios{
				Anonymous: func() {
					log.Println("i am anonymous")
				},
				Member: func(email string, since time.Time) {
					log.Println("member", email, since)
				},
				Admin: func(email string) {
					log.Println("admin", email)
				},
			})
		})
	}
}
