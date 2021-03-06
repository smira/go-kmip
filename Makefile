all: test lint

.PHONY: test
test:
	go test -v -race -coverprofile=coverage.txt -covermode=atomic -count 1 .

.PHONY: lint
lint:
	golangci-lint run
