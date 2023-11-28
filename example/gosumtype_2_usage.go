package main

import (
	"log"
	"time"
)

// Example usage
func main() {
	user1 := Anonymous()
	user2 := Member("Alice", time.Now())
	user3 := Admin("Bob")

	log.Println(
		"\nUser1:", UserString(user1),
		"\nUser2:", UserString(user2),
		"\nUser3:", UserString(user3),
	)
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
