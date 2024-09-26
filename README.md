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
type UserVariants struct { // be sure to suffix the name with `Variants`
	Anonymous func()
	Member    func(email string, since time.Time)
	Admin     func(email string)
}
```

Execute this command to generate `User` type from your `UserVariants` struct

```sh
go install github.com/choonkeat/sumtype-go@v0.4.1 # to install
sumtype-go -input declaration.go
```

To generate `declaration.boilerplate.go` and start using `User`!

```go
users := []User{
	Anonymous(),                 // this returns a `User` value
	Member("Alice", time.Now()), // this also returns a `User` value
	Admin("Bob"),                // this also returns a `User` value
}
```

and we can pattern match `User` values and return a different value depending on pattern matched variant

```go
userString := UserMap(user, UserVariantsMap[string]{
	Anonymous: func() string {
		return "Anonymous coward"
	},
	Member: func(email string, since time.Time) string {
		return email + " (member since " + since.String() + ")"
	},
	Admin: func(email string) string {
		return email + " (admin)"
	},
})
```

or _do_ different things depending on pattern matched variant

```go
func SendReply(u User, comment Comment) {
	u.Match(UserVariants{
		Anonymous: func() {
			// noop
		},
		Member: func(email string, since time.Time) {
			sendEmail(email, comment)
		},
		Admin: func(email string) {
			sendEmail(email, comment)
		},
	})
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
type ResultVariants[x, a any] struct {
	Err func(err x)
	Ok  func(data a)
}
```

Same thing, after executing `sumtype-go` to generate the `.boilerplate.go` file, you can use
the generated `Result` like this

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
go install github.com/choonkeat/sumtype-go@v0.4.1
```

## Usage

After installation, you can start using `sumtype-go` by invoking it from the command line.

```
$ sumtype-go -h
Usage of sumtype-go:
  -input string
    	Input file name
  -pattern-match string
    	Name of the pattern match method (default "Match")
  -suffix string
    	Suffix of the struct defining variants (default "Variants")
```

Here's a basic example of how to use it:

```
sumtype-go -input example/gosumtype_1_declaration.go
```
