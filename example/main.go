package main

import (
	"fmt"
	"log"
	"time"
)

// Example usage
func main() {
	users := []User{
		Anonymous(Paypal("nobody@example.com")), // this returns a `User` value
		Member("Alice", time.Now()),             // this also returns a `User` value
		Admin("Bob"),                            // this also returns a `User` value
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

	trees := []Tree[int]{
		Branch[int](Leaf[int](1), Leaf[int](2)), // this returns a `Tree` value
		Leaf[int](3),                            // this also returns a `Tree` value
	}

	for i, tree := range trees {
		log.Println(i, TreeString(tree))
	}
}

func UserString(user User) string {
	var result string
	user.Match(UserVariants{
		Anonymous: func(paymentMethod PaymentMethod) {
			result = "Anonymous coward" + fmt.Sprintf("%#v", paymentMethod)
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
	result.Match(ResultVariants[string, int]{
		Err: func(err string) {
			log.Println(i, "Error:", err)
		},
		Ok: func(data int) {
			log.Println(i, "Data:", data)
		},
	})
}

func TreeString(t Tree[int]) string {
	var result string
	t.Match(TreeVariants[int]{
		Branch: func(left, right Tree[int]) {
			result = "Branch(" + TreeString(left) + ", " + TreeString(right) + ")"
		},
		Leaf: func(s int) {
			result = fmt.Sprintf("Leaf(%d)", s)
		},
	})
	return result
}
