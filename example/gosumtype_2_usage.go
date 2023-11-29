package main

import (
	"log"
	"time"
)

// Example usage
func main() {
	users := []User{
		Anonymous(),                 // this returns a `User` value
		Member("Alice", time.Now()), // this also returns a `User` value
		Admin("Bob"),                // this also returns a `User` value
	}

	for i, user := range users {
		log.Println(i, UserString(user))
	}

	results := []Result[string, int]{
		Err[string, int]("Oops err"), // this returns a `Result` value
		Ok[string, int](42),          // this also returns a `Result` value
	}

	for i, result := range results {
		HandleResult(i, result)
	}
}

func UserString(u User) string {
	var result string
	u.Switch(UserScenarios{
		Anonymous: func() {
			result = "Anonymous coward"
		},
		Member: func(email string, since time.Time) {
			result = email + " (member since " + since.String() + ")"
		},
		Admin: func(email string) {
			result = email + " (admin)"
		},
	})
	return result
}

func HandleResult(i int, result Result[string, int]) {
	result.Switch(ResultScenarios[string, int]{
		Err: func(err string) {
			log.Println(i, "Error:", err)
		},
		Ok: func(data int) {
			log.Println(i, "Data:", data)
		},
	})
}
