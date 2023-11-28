run:
	go run *.go -input example/gosumtype_1_declaration.go
	go test ./...
	go run example/*1*.go example/*2*.go
