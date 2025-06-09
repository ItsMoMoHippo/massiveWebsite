# simple makefile for project

# defualt target
run:
	templ generate
	go run .

# generate templ files seperatly 
generate:
	templ generate

# build binary
build:
	templ generate
	go build .

# clean generated files
clean:
	go clean

.PHONY: run generate clean build
