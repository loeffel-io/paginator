install:
	go get

linter:
	golangci-lint run --timeout 3m0s

test-coverage:
	MAKELESS_CONFIG=./makeless.local.json CGO_ENABLED=1 go test -shuffle=on -v -race -coverprofile=coverage.txt -covermode=atomic ./...
