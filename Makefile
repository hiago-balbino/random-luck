.PHONY: all help setup vet lint vulncheck fmt clean

APP_NAME=random_luck_api

## help: show this help.
help:
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

## setup: run the command mod download and tidy from Go
setup:
	GO111MODULE=on go mod download
	go mod tidy
	go mod verify

## vet: run the command vet from Go
vet:
	go vet ./...

## lint: run all linters configured
lint:
	golangci-lint run ./...

## vulncheck: run all vulnerability checks
vulncheck:
	govulncheck ./...

## fmt: run go formatter recursively on all files
fmt:
	gofmt -s -w .

## clean: run the go clean command and removes the application binary
clean:
	go clean
	rm ${APP_NAME}

all: help