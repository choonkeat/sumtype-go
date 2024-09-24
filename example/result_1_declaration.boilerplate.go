// Generated code by github.com/choonkeat/sumtype-go
package main

import (
	"encoding/json"
	"fmt"
)

// Err
type errResultVariants[x, a any] struct {
	Err x
}

func (s errResultVariants[x, a]) Match(variants ResultVariants[x, a]) {
	variants.Err(s.Err)
}

// Err is a constructor function for Result; see ResultVariants for all constructor functions of Result
func Err[x, a any](errArg x) Result[x, a] {
	return Result[x, a]{errResultVariants[x, a]{errArg}}
}

// Ok
type okResultVariants[x, a any] struct {
	Data a
}

func (s okResultVariants[x, a]) Match(variants ResultVariants[x, a]) {
	variants.Ok(s.Data)
}

// Ok is a constructor function for Result; see ResultVariants for all constructor functions of Result
func Ok[x, a any](dataArg a) Result[x, a] {
	return Result[x, a]{okResultVariants[x, a]{dataArg}}
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

// Result = Err | Ok
type Result[x, a any] struct {
	result result[x, a]
}

// result is the interface for ResultVariants
type result[x, a any] interface {
	Match(variants ResultVariants[x, a])
}

func (s Result[x, a]) Match(variants ResultVariants[x, a]) {
	s.result.Match(variants)
}
func (s Result[x, a]) MarshalJSON() (data []byte, err error) {
	s.result.Match(ResultVariants[x, a]{
		Err: func(errArg x) {
			data, err = json.Marshal([]any{
				"Err",
				errResultVariants[x, a]{
					Err: errArg,
				}})
		},
		Ok: func(dataArg a) {
			data, err = json.Marshal([]any{
				"Ok",
				okResultVariants[x, a]{
					Data: dataArg,
				}})
		},
	})
	return data, err
}
func (s *Result[x, a]) UnmarshalJSON(data []byte) error {
	// The expected format is ["TypeName", { ... data... }]
	var raw []json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return fmt.Errorf("expected an array with type and data, got error: %w", err)
	}
	if len(raw) != 2 {
		return fmt.Errorf("expected array of two elements [type, data], got %d elements", len(raw))
	}
	// Unmarshal the first element to get the type
	var typeName string
	if err := json.Unmarshal(raw[0], &typeName); err != nil {
		return fmt.Errorf("failed to unmarshal type name: %w", err)
	}
	switch typeName {
	case "Err":
		var temp errResultVariants[x, a]
		if err := json.Unmarshal(raw[1], &temp); err != nil {
			return fmt.Errorf("failed to unmarshal data: %w", err)
		}
		s.result = errResultVariants[x, a]{
			Err: temp.Err,
		}
	case "Ok":
		var temp okResultVariants[x, a]
		if err := json.Unmarshal(raw[1], &temp); err != nil {
			return fmt.Errorf("failed to unmarshal data: %w", err)
		}
		s.result = okResultVariants[x, a]{
			Data: temp.Data,
		}
	default:
		return fmt.Errorf("unknown type %q", typeName)
	}
	return nil
}
