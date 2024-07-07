// Generated code by github.com/choonkeat/sumtype-go
package main

// Err
type errResultVariants[x, a any] struct {
	err x
}

func (s errResultVariants[x, a]) Match(variants ResultVariants[x, a]) {
	variants.Err(s.err)
}

// Err is a constructor function for Result; see ResultVariants for all constructor functions of Result
func Err[x, a any](errArg x) Result[x, a] {
	return errResultVariants[x, a]{errArg}
}

// Ok
type okResultVariants[x, a any] struct {
	data a
}

func (s okResultVariants[x, a]) Match(variants ResultVariants[x, a]) {
	variants.Ok(s.data)
}

// Ok is a constructor function for Result; see ResultVariants for all constructor functions of Result
func Ok[x, a any](dataArg a) Result[x, a] {
	return okResultVariants[x, a]{dataArg}
}

// ResultVariantsMap is parameter type of ResultMap function,
// like ResultVariants is parameter type of result.Match method,
// but with methods that returns a value of generic type
type ResultVariantsMap[x, a, A any] struct {
	Err func(errArg x) A  // when Result value pattern matches to Err, return different value
	Ok  func(dataArg a) A // when Result value pattern matches to Ok, return different value
}

// ResultMap is like result.Match method except it returns a value of generic type
// thus can transform a Result value into anything else
func ResultMap[x, a, A any](value Result[x, a], variants ResultVariantsMap[x, a, A]) A {
	var result A
	value.Match(ResultVariants[x, a]{
		Err: func(errArg x) {
			result = variants.Err(errArg)
		},
		Ok: func(dataArg a) {
			result = variants.Ok(dataArg)
		},
	})
	return result
}
