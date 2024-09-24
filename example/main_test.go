package main

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
)

func TestJSONMarshalUnmarshal(t *testing.T) {
	assertJSONMarshalUnmarshal(t, "users", users)
	assertJSONMarshalUnmarshal(t, "results", results)
	assertJSONMarshalUnmarshal(t, "trees", trees)
}

func assertJSONMarshalUnmarshal[T any](t *testing.T, name string, given T) {
	data, err := json.MarshalIndent(given, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile("testdata/"+name+".json", data, 0o600); err != nil {
		t.Fatal(err)
	}
	var actual T
	err = json.Unmarshal(data, &actual)
	if err != nil {
		t.Fatal(err, string(data))
	}
	if !reflect.DeepEqual(given, actual) {
		t.Fatalf("expect %#v but got %#v", given, actual)
	}
}
