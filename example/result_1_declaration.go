package main

// To define this sum type:
//
//	type Result x a
//		= Err x
//		| Ok a
//
// we declare `Result` as an interface
type Result[x, a interface{}] interface {
	Switch(s ResultScenarios[x, a])
}

// and the variants as a struct
type ResultScenarios[x, a interface{}] struct {
	Err func(err x)
	Ok  func(data a)
}
