// Generated code by github.com/choonkeat/sumtype-go
package main

// Err
type errResultVariants[x, a any] struct {
	err x
}

func (s errResultVariants[x, a]) Match(Variants ResultVariants[x, a]) {
	Variants.Err(s.err)
}

func Err[x, a any](errArg x) Result[x, a] {
	return errResultVariants[x, a]{errArg}
}

// Ok
type okResultVariants[x, a any] struct {
	data a
}

func (s okResultVariants[x, a]) Match(Variants ResultVariants[x, a]) {
	Variants.Ok(s.data)
}

func Ok[x, a any](dataArg a) Result[x, a] {
	return okResultVariants[x, a]{dataArg}
}
