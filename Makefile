run: example-generate example-run test

example-generate:
	go run *.go -input example/gosumtype_1_declaration.go
	go run *.go -input example/result_1_declaration.go
	go run *.go -input example/tree_1_declaration.go

example-run:
	go run `ls example/*.go | grep -v "_test.go" | xargs -I {} echo {}`

test:
	go test ./...
