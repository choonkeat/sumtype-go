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

func (errInstance errResultVariants[x, a]) Match(errVariants ResultVariants[x, a]) {
	errVariants.Err(errInstance.Err)
}

// Err is a constructor function for Result; see ResultVariants for all constructor functions of Result
func Err[x, a any](errArg x) Result[x, a] {
	return Result[x, a]{errResultVariants[x, a]{errArg}}
}

// Ok
type okResultVariants[x, a any] struct {
	Data a
}

func (okInstance okResultVariants[x, a]) Match(okVariants ResultVariants[x, a]) {
	okVariants.Ok(okInstance.Data)
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
func ResultMap[x, a, A any](resultValue Result[x, a], resultVariants ResultVariantsMap[x, a, A]) A {
	var resultTemp A
	resultValue.Match(ResultVariants[x, a]{
		Err: func(errArg x) {
			resultTemp = resultVariants.Err(errArg)
		},
		Ok: func(dataArg a) {
			resultTemp = resultVariants.Ok(dataArg)
		},
	})
	return resultTemp
}

// Result = Err | Ok
type Result[x, a any] struct {
	result result[x, a]
}

// result is the interface for ResultVariants
type result[x, a any] interface {
	Match(variants ResultVariants[x, a])
}

func (resultInstance Result[x, a]) Match(resultVariants ResultVariants[x, a]) {
	resultInstance.result.Match(resultVariants)
}
func (resultInstance Result[x, a]) MarshalJSON() (resultData []byte, resultErr error) {
	resultInstance.result.Match(ResultVariants[x, a]{
		Err: func(errArg x) {
			resultData, resultErr = json.Marshal([]any{
				"Err",
				errResultVariants[x, a]{
					Err: errArg,
				}})
		},
		Ok: func(dataArg a) {
			resultData, resultErr = json.Marshal([]any{
				"Ok",
				okResultVariants[x, a]{
					Data: dataArg,
				}})
		},
	})
	return resultData, resultErr
}
func (resultInstance *Result[x, a]) UnmarshalJSON(resultData []byte) error {
	// The expected format is ["TypeName", { ... data... }]
	var resultRaw []json.RawMessage
	if err := json.Unmarshal(resultData, &resultRaw); err != nil {
		return fmt.Errorf("expected an array with type and data, got error: %w", err)
	}
	if len(resultRaw) != 2 {
		return fmt.Errorf("expected array of two elements [type, data], got %d elements", len(resultRaw))
	}
	// Unmarshal the first element to get the type
	var resultVariantName string
	if err := json.Unmarshal(resultRaw[0], &resultVariantName); err != nil {
		return fmt.Errorf("failed to unmarshal type name: %w", err)
	}
	switch resultVariantName {
	case "Err":
		var resultTemp errResultVariants[x, a]
		if err := json.Unmarshal(resultRaw[1], &resultTemp); err != nil {
			return fmt.Errorf("failed to unmarshal data: %w", err)
		}
		resultInstance.result = errResultVariants[x, a]{
			Err: resultTemp.Err,
		}
	case "Ok":
		var resultTemp okResultVariants[x, a]
		if err := json.Unmarshal(resultRaw[1], &resultTemp); err != nil {
			return fmt.Errorf("failed to unmarshal data: %w", err)
		}
		resultInstance.result = okResultVariants[x, a]{
			Data: resultTemp.Data,
		}
	default:
		return fmt.Errorf("unknown type %q", resultVariantName)
	}
	return nil
}
