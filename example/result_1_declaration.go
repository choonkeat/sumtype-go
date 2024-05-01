package main

// To define this sum type:
//
//	type Result x a
//		= Err x
//		| Ok a
//
// we declare `Result` as an interface
type Result[x, a any] interface {
	Match(s ResultVariants[x, a])
}

// and the variants as a struct
type ResultVariants[x, a any] struct {
	Err func(err x)
	Ok  func(data a)
}
