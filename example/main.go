package main

import (
	"fmt"
	"log"
	"math"
	"time"
)

var users = []User{
	Anonymous(Paypal("nobody@example.com")), // this returns a `User` value
	Member("Alice", time.Now()),             // this also returns a `User` value
	Admin("Bob"),                            // this also returns a `User` value
}

var results = []Result[string, int]{
	Err[string, int]("Oops err"), // this returns a `Result` value
	Ok[string, int](42),          // this also returns a `Result` value
}

var trees = []Tree[int]{
	Branch[int](
		Leaf[int](1),
		Branch[int](
			Leaf[int](42),
			Leaf[int](2),
		)), // this returns a `Tree` value
	Leaf[int](3), // this also returns a `Tree` value
}

// Example usage
func main() {
	for i, user := range users {
		printUserString(i, user)
	}

	// map
	userToCode := UserVariantsMap[int]{
		Anonymous: func(paymentMethod PaymentMethod) int {
			return 1
		},
		Member: func(email string, since time.Time) int {
			return 20
		},
		Admin: func(email string) int {
			return 300
		},
	}
	for i, user := range users {
		log.Printf("%d UserMap: %#v -> %d\n", i, user, UserMap(user, userToCode))
	}

	for i, result := range results {
		printHandleResult(i, result)
	}

	// map
	resultToCode := ResultVariantsMap[string, int, int]{
		Err: func(err string) int {
			return -1
		},
		Ok: func(data int) int {
			return data
		},
	}
	for i, result := range results {
		log.Printf("%d ResultMap: %#v -> %d\n", i, result, ResultMap(result, resultToCode))
	}

	for _, tree := range trees {
		printTreeString(0, tree)
	}

	// map
	for i, tree := range trees {
		log.Printf("%d TreeMap: %#v -> %f\n", i, tree, DepthOf(tree))
	}
}

func printUserString(i int, user User) {
	user.Match(UserVariants{
		Anonymous: func(paymentMethod PaymentMethod) {
			log.Println(i, "Anonymous coward"+fmt.Sprintf("%#v", paymentMethod))
		},
		Member: func(email string, since time.Time) {
			log.Println(i, email+" (member since "+since.String()+")")
		},
		Admin: func(email string) {
			log.Println(i, email+" (admin)")
		},
	})
}

func printHandleResult(i int, result Result[string, int]) {
	result.Match(ResultVariants[string, int]{
		Err: func(err string) {
			log.Println(i, "Error:", err)
		},
		Ok: func(data int) {
			log.Println(i, "Data:", data)
		},
	})
}

func printTreeString(i int, t Tree[int]) {
	format := "%" + fmt.Sprintf("%d", i*2) + "s %s\n"
	t.Match(TreeVariants[int]{
		Branch: func(left, right Tree[int]) {
			log.Printf(format, "-", "Branch")
			printTreeString(i+1, left)
			printTreeString(i+1, right)
		},
		Leaf: func(s int) {
			log.Printf(format, "-", fmt.Sprintf("Leaf(%d)", s))
		},
	})
}

func DepthOf(t Tree[int]) float64 {
	return TreeMap(t, TreeVariantsMap[int, float64]{
		Branch: func(left, right Tree[int]) float64 {
			return 1 + math.Max(DepthOf(left), DepthOf(right))
		},
		Leaf: func(s int) float64 {
			return 1
		},
	})
}
