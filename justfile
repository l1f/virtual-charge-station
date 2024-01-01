set dotenv-load
default: build

_check-build-deps:
	go version
	templ version

_check-dev-deps: _check-build-deps
	air -v

_check-migration-deps:
	migrate -version

build: _check-build-deps
	go mod download
	templ generate
	mkdir -p build
	go build -ldflags="-s -w" -o ./build/vss cmd/main.go

dev: _check-dev-deps
	air

