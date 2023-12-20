[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/hiago-balbino/random-luck)](https://goreportcard.com/report/github.com/hiago-balbino/random-luck)
[![Linter](https://github.com/hiago-balbino/random-luck/actions/workflows/linter.yml/badge.svg?branch=main&event=push&logo=github&style=flat-square)](https://github.com/hiago-balbino/random-luck/actions/workflows/linter.yml)
[![Security](https://github.com/hiago-balbino/random-luck/actions/workflows/security.yml/badge.svg?branch=main&event=push&logo=github&style=flat-square)](https://github.com/hiago-balbino/random-luck/actions/workflows/security.yml)
[![Tests](https://github.com/hiago-balbino/random-luck/actions/workflows/tests.yml/badge.svg?branch=main&event=push&logo=github&style=flat-square)](https://github.com/hiago-balbino/random-luck/actions/workflows/tests.yml)

# 🍀 RANDOM LUCK

This project was created for learning purposes and is a lucky random number generator for Mega Sena.

Some tools used do not represent the best choice, they were only used for learning purposes. For example the application can be run as a Command Line or it can be run as a Web where the `template package` was used.

## 🧰 Dependencies

- [Go](https://golang.google.cn/dl) 1.20+
- [Docker](https://www.docker.com/products/docker-desktop)
- [Docker-compose](https://docs.docker.com/compose/install)
- [Nginx](https://nginx.org)
- [GNU make](https://www.gnu.org/software/make)
- [Direnv](https://direnv.net)
  - This is not mandatory but is a easily way to control your environment variable for each project without configuring the variables globally
- [Golangci-lint](https://golangci-lint.run)
- [Godoc](https://pkg.go.dev/golang.org/x/tools/cmd/godoc)
- [Govulncheck](https://pkg.go.dev/golang.org/x/vuln/cmd/govulncheck)
- [Viper](https://github.com/spf13/viper)
  - Configuration solution
- [Gin web framework](https://github.com/gin-gonic/gin)
- [Testify](https://github.com/stretchr/testify)
  - Tools for testifying
- [Zap logger](https://go.uber.org/zap)
- [HTTP expect](https://github.com/gavv/httpexpect)
  - Used for REST API testing
- [Statsviz](https://github.com/arl/statsviz)
- [Sonarqube](https://www.sonarqube.org)

## 🛠️ Useful commands

You can run the command below to see all the useful commands available and your goals.

```
make help
```

```
help: show this help
setup: run the command mod download and tidy from Go
fmt: run go formatter recursively on all files
lint: run all linters configured
vulncheck: run all vulnerability checks
tests: run all unit tests
pipeline: run all necessary steps to go through the pipeline
cover: run the command tool cover to open coverage file as HTML
sonarqube-up: start sonarqube container
sonarqube-down: stop sonarqube container
sonarqube-analysis: run sonar scanner
build-ctl: create an executable of the command tool
build-web: create an executable of the web application
run-ctl: build project and run the command tool using the built binary
run-web: build project and run the Web using the built binary
compose-ps: list all containers running
compose-up: start the app with docker
compose-down: stop the app with docker
clean: run the go clean command and removes the application binary
doc: run the project documentation using HTTP
```

## ⚙️ Running the Application

### Command Line

You can install the application using `go install`.

```
go install github.com/hiago-balbino/random-luck/v2/cmd/luckctl@latest
```

Then you can run the application using `luckctl`. Note that if you want to run anywhere you need to export the `$GOPATH/bin/luckctl` to your `$PATH` environment variables.

```
luckctl --games=2 --numbers=6
```

If you prefer can just run `make run-ctl`.

```
make run-ctl games=2 numbers=6
```

### Web

To run the project locally you need to export some environment variables and this can be done using `direnv`. You can export the variables below.

```
NGINX_PORT='80'
LOG_LEVEL='ERROR'
```

After exporting the environment variables, you can run the `make compose-up` command and the application will run on Docker.

- The application will be accessible at `http://localhost/index` and the parameters need to be filled in and you will do this by following the HTML form:
  - amount_of_games (must be greater than zero)
  - amount_of_numbers_per_game (must be greater than or equal to six and less than or equal to nine)

If you prefer to run outside of docker, just run `make run-web` (accessible at `http://localhost:8888/index`).

## 🪲 Debugging

- If you want to debug the application, in your IDE you need to run the main file inside the `cmd/luckweb` or `cmd/luckctl` directory.

## 📜 Running Internal Documentation

You can do this by running the `make doc` command and going to the address `http://localhost:6060`.

## 📊 Runtime metrics

In the Web application you can view real-time Go program runtime metrics data in your browser at `http://localhost/metrics` or `http://localhost:8888/metrics`(when running locally).

- Heap (global)
- Heap (details)
- Live Objects in Heap
- Live Bytes in Heap
- MSpan/MCache
- Goroutines
- Size Classes
- Stop-the-world Pause Latencies
- Time Goroutines Spend in 'Runnable'
- Starting Size of Goroutines Stacks
- Goroutine Scheduling Events
- CGO Calls

## 🎯 How to run sonarqube locally

The project is set up to run `sonarqube` locally and has three commands in the Makefile. The `sonarqube` will be downloaded by Docker, but you need to [install sonar-scanner following your operating system](https://docs.sonarqube.org/latest/analyzing-source-code/scanners/sonarscanner).

To run `sonarqube` locally, you need to export the following environment variables. You can do this using `direnv`.

```
SONAR_PORT='9000'
SONAR_HOST='http://localhost:9000'
SONAR_LOGIN='admin'
SONAR_PASSWORD='admin'
SONAR_BINARY='Here you need to fill it according to your operational system. Example: sonar-scanner for Linux/MacOS or sonar-scanner.bat for Windows'
```

After installing and configuring `sonar-scanner` in `$PATH`(if needed) you will be able to run the commands below. By running the `make sonarqube-up` and `make sonarqube-analysis` commands you can open the `http://localhost:9000` address in your browser and login and password as `admin`(perhaps `sonarqube` may prompt you to change your password).

```
sonarqube-up: start sonarqube container
sonarqube-analysis: run sonar scanner
sonarqube-down: stop sonarqube container
```

# 🤞 Good Luck...
