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
	// map
	userCodes := []int{}
	for _, user := range users {
		userCodes = append(userCodes, UserMap(user, UserVariantsMap[int]{
			Anonymous: func(paymentMethod PaymentMethod) int {
				return 1
			},
			Member: func(email string, since time.Time) int {
				return 20
			},
			Admin: func(email string) int {
				return 300
			},
		}))
	}
	fmt.Println("User -> Int userCodes =", userCodes)

	results := []Result[string, int]{
		Err[string, int]("Oops err"), // this returns a `Result` value
		Ok[string, int](42),          // this also returns a `Result` value
	}

	for i, result := range results {
		HandleResult(i, result)
	}
	// map
	resultCodes := []int{}
	for _, result := range results {
		resultCodes = append(resultCodes, ResultMap(result, ResultVariantsMap[string, int, int]{
			Err: func(err string) int {
				return -1
			},
			Ok: func(data int) int {
				return data
			},
		}))
	}
	fmt.Println("Result -> Int resultCodes =", resultCodes)

	trees := []Tree[int]{
		Branch[int](Leaf[int](1), Leaf[int](2)), // this returns a `Tree` value
		Leaf[int](3),                            // this also returns a `Tree` value
	}

	for i, tree := range trees {
		log.Println(i, TreeString(tree))
	}

	// map
	treeValues := []int{}
	for _, tree := range trees {
		treeValues = append(treeValues, TreeMap(tree, TreeVariantsMap[int, int]{
			Branch: func(leftArg Tree[int], rightArg Tree[int]) int {
				return TreeMap(leftArg, TreeVariantsMap[int, int]{
					Branch: func(leftArg Tree[int], rightArg Tree[int]) int {
						return 0
					},
					Leaf: func(sArg int) int {
						return sArg
					},
				}) + TreeMap(rightArg, TreeVariantsMap[int, int]{
					Branch: func(leftArg Tree[int], rightArg Tree[int]) int {
						return 0
					},
					Leaf: func(sArg int) int {
						return sArg
					},
				})
			},
			Leaf: func(sArg int) int {
				return sArg
			},
		}))
	}
	fmt.Println("Tree -> Int treeValues =", treeValues)
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
