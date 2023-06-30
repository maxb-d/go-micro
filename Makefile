build:
	@go build -o bin/gomicro

run: build
	@./bin/gomicro

test:
	@go test -v ./...

test_coverage:
	@go test ./... -coverprofile=coverage.out