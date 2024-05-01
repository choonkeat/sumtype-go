package main

import (
	"fmt"
	"log"
	"time"
)

// Example usage
func main() {
	users := []User{
		Anonymous(CreditCard("xxx1234", time.Now())),                   // this returns a `User` value
		Member(CreditCard("xxx1234", time.Now()), "Alice", time.Now()), // this also returns a `User` value
		Admin(Paypal("nobody@example.com"), "Bob"),                     // this also returns a `User` value
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

func UserString(u User) string {
	var result string
	u.Switch(UserScenarios{
		Anonymous: func(paymentMethod PaymentMethod) {
			result = "Anonymous coward" + fmt.Sprintf("%#v", paymentMethod)
		},
		Member: func(paymentMethod PaymentMethod, email string, since time.Time) {
			result = email + " (member since " + since.String() + ")" + fmt.Sprintf("%#v", paymentMethod)
		},
		Admin: func(paymentMethod PaymentMethod, email string) {
			result = email + " (admin)" + fmt.Sprintf("%#v", paymentMethod)
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

func TreeString(t Tree[int]) string {
	var result string
	t.Switch(TreeScenarios[int]{
		Branch: func(left, right Tree[int]) {
			result = "Branch(" + TreeString(left) + ", " + TreeString(right) + ")"
		},
		Leaf: func(s int) {
			result = fmt.Sprintf("Leaf(%d)", s)
		},
	})
	return result
}
