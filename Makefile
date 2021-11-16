install:
	go get

linter:
	golangci-lint run --timeout 3m0s

test-coverage:
	MAKELESS_CONFIG=./makeless.local.json CGO_ENABLED=1 go test -shuffle=on -v -race -coverprofile=coverage.txt -covermode=atomic ./...

proto:
	protoc \
		-I . \
		-I ${GOPATH}/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.6.2 \
		--go_opt=paths=source_relative \
		--validate_opt=paths=source_relative \
		--go_out=. \
        --validate_out="lang=go:." \
		--doc_out=./doc \
		--doc_opt=html,index.html \
        *.proto