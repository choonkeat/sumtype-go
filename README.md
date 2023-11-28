# sumtype-go

## Introduction

`sumtype-go` is a CLI tool designed to facilitate the creation and management of sum types in Go. This tool simplifies the process of generating boilerplate code for sum types, making it easier to work with variant types in Go.

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
	Switch(s UserScenarios)
}

type UserScenarios struct {
	Anonymous func()
	Member    func(email string, since time.Time)
	Admin     func(email string)
}
```

Execute this command

```
go install github.com/choonkeat/sumtype-go@latest
sumtype-go -input declaration.go
```

To generate `declaration.boilerplate.go` and start using like

```go
func main() {
	user1 := Anonymous()
	user2 := Member("Alice", time.Now())
	user3 := Admin("Bob")

	log.Println(
		"\nUser1:", UserString(user1),
		"\nUser2:", UserString(user2),
		"\nUser3:", UserString(user3),
	)
}

func UserString(u User) string {
	var result string
	u.Switch(UserScenarios{
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

Refer to `example/`

## Installation
To install `sumtype-go`, ensure you have Go installed on your system, and then run the following command:

```sh
go install github.com/choonkeat/sumtype-go@v0.1.0
```

## Usage

After installation, you can start using `sumtype-go`` by invoking it from the command line. Here's a basic example of how to use it:

```
sumtype-go -input [declaration.go]
```