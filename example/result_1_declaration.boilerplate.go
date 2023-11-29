// Generated code by github.com/choonkeat/sumtype-go
package main

// Err
type errResultScenarios[x, a interface{}] struct {
	err x
}

func (s errResultScenarios[x, a]) Switch(scenarios ResultScenarios[x, a]) {
	scenarios.Err(s.err)
}

func Err[x, a interface{}](errArg x) Result[x, a] {
	return errResultScenarios[x, a]{errArg}
}

// Ok
type okResultScenarios[x, a interface{}] struct {
	data a
}

func (s okResultScenarios[x, a]) Switch(scenarios ResultScenarios[x, a]) {
	scenarios.Ok(s.data)
}

func Ok[x, a interface{}](dataArg a) Result[x, a] {
	return okResultScenarios[x, a]{dataArg}
}
