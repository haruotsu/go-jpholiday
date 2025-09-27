.PHONY: help setup test test-coverage test-race bench lint fmt vet security check build install run run-dry clean deps update-deps docs profile ci tag
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

setup:
	@echo "Setting up project structure..."
	@mkdir -p holiday fetcher data cmd/update-holidays .github/workflows
	@go mod init github.com/haruotsu/go-jpholiday 2>/dev/null || true
	@go mod tidy
	@echo "Setup complete!"

test:
	go test -v ./...

test-coverage:
	go test -v -race -coverprofile=coverage.out -covermode=atomic ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

test-race:
	go test -race ./...

lint:
	@which golangci-lint > /dev/null 2>&1 || (echo "Installing golangci-lint..." && go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest)
	golangci-lint run ./...

fmt:
	go fmt ./...
	gofmt -s -w .

vet:
	go vet ./...

security:
	@which gosec > /dev/null 2>&1 || (echo "Installing gosec..." && go install github.com/securego/gosec/v2/cmd/gosec@latest)
	gosec ./...

check: fmt vet lint test

build:
	go build -o bin/update-holidays cmd/update-holidays/main.go

install:
	go install ./cmd/update-holidays

run:
	@if [ -z "$(GOOGLE_API_KEY)" ]; then \
		echo "Error: GOOGLE_API_KEY is not set"; \
		exit 1; \
	fi
	go run cmd/update-holidays/main.go

run-dry:
	@if [ -z "$(GOOGLE_API_KEY)" ]; then \
		echo "Error: GOOGLE_API_KEY is not set"; \
		exit 1; \
	fi
	go run cmd/update-holidays/main.go -dry-run

clean:
	rm -rf bin/ coverage.* *.test
	go clean -testcache

deps:
	go mod download
	go mod tidy

docs:
	@which godoc > /dev/null 2>&1 || (echo "Installing godoc..." && go install golang.org/x/tools/cmd/godoc@latest)
	@echo "Opening documentation at http://localhost:6060/pkg/github.com/haruotsu/go-jpholiday/"
	@godoc -http=:6060

profile:
	go test -cpuprofile=cpu.prof -memprofile=mem.prof -bench=.
	go tool pprof cpu.prof

ci: deps check test-coverage

tag:
	@which tagpr > /dev/null 2>&1 || (echo "Installing tagpr..." && go install github.com/Songmu/tagpr@latest)
	tagpr

.DEFAULT_GOAL := help
