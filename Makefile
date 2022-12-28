.PHONY: all help setup vet lint vulncheck fmt tests cover sonarqube-up sonarqube-down sonarqube-analysis build clean

APP_NAME=random_luck

## help: show this help
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

## tests: run all unit tests
tests:
	go test -race -coverprofile coverage.out ./... -short=true -count=1

## cover: run the command tool cover to open coverage file as HTML
cover: tests
	go tool cover -html coverage.out

## sonarqube-up: start sonarqube container
sonarqube-up:
	docker run -d --name sonarqube -p ${SONAR_PORT}:${SONAR_PORT} sonarqube

## sonarqube-down: stop sonarqube container
sonarqube-down:
	docker rm sonarqube -f

## sonarqube-analysis: run sonar scanner
sonarqube-analysis: tests
	${SONAR_BINARY} -Dsonar.host.url=${SONAR_HOST} -Dsonar.login=${SONAR_LOGIN} -Dsonar.password=${SONAR_PASSWORD}

## build: create an executable of the application
build:
	go build -o ${APP_NAME} .

## clean: run the go clean command and removes the application binary
clean:
	go clean
	rm ${APP_NAME}

all: help