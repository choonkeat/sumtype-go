# sumtype-go

## Introduction

`sumtype-go` is a CLI tool designed to facilitate the creation and management of sum types (aka union types) in Go. This tool simplifies the process of generating boilerplate code for discriminated union types, making it easier to work with its variants in Go.

## Quick Tour

To define this sum type:

```elm
type User
    = Anonymous
    | Member String Time
    | Admin String
```

write this Go type definition, e.g. in `declaration.go`

```go
type User interface {
	Match(s UserVariants)
}

type UserVariants struct { // be sure to suffix the name with `Variants`
	Anonymous func()
	Member    func(email string, since time.Time)
	Admin     func(email string)
}
```

Execute this command

```sh
go install github.com/choonkeat/sumtype-go@v0.3.2 # to install
sumtype-go -input declaration.go
```

To generate `declaration.boilerplate.go` and start using it!

```go
users := []User{
	Anonymous(),                 // this returns a `User` value
	Member("Alice", time.Now()), // this also returns a `User` value
	Admin("Bob"),                // this also returns a `User` value
}
```

and we can write functions that work with `User`

```go
func UserString(u User) string {
	var result string
	u.Match(UserVariants{
		Anonymous: func() {
			result = "Anonymous coward"
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
```

Refer to `example/gosumtype_1_*.go`

## Generics

We support generics too. e.g. the classic `Result` type

```elm
type Result x a
    = Err x
    | Ok a
```

can be defined as

```go
type Result[x, a any] interface {
	Match(s ResultVariants[x, a])
}

type ResultVariants[x, a any] struct {
	Err func(err x)
	Ok  func(data a)
}
```

Same thing, after executing `sumtype-go` to generate the `.boilerplate.go` file, you can use it like

```go
results := []Result[string, int]{
	Err[string, int]("Oops err"), // this returns a `Result` value
	Ok[string, int](42),          // this also returns a `Result` value
}

for i, result := range results {
	HandleResult(i, result) // implement your own `func HandleResult(int, Result[string, int])`
}
```

Refer to `example/result_1_*.go`

## Installation

To install `sumtype-go`, ensure you have Go installed on your system, and then run the following command:

```sh
go install github.com/choonkeat/sumtype-go@v0.3.2
```

## Usage

After installation, you can start using `sumtype-go` by invoking it from the command line. Here's a basic example of how to use it:

```
  -input string
    	Input file name
  -pattern-match string
    	Name of the pattern match method (default "Match")
  -suffix string
    	Suffix of the struct defining variants (default "Variants")
```
