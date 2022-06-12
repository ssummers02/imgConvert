LINTER_URL := https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh

# Default command for prepare code for a commit
default: format lint-fix

### Linter ###

lint-install:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.44.0

# [CI\CD] lint code
lint:
	golangci-lint run
# lint and auto-fix possible problems
lint-fix:
	golangci-lint run --fix


# [CI\CD] Auto-format code
format:
	gofmt -s -w . && \
	go vet ./... && \
	go mod tidy
