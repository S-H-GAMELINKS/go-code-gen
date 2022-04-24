NAME := go-code-gen
VERSION := 0.1.0

code-gen:
	go run cmd/code-gen/${type}/main.go ${name}
