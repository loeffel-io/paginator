install:
	(cd go && make install)
	(cd js && make install)

proto:
	protoc \
		-I . \
		-I ${GOPATH}/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.6.2 \
		--go_opt=paths=source_relative \
		--validate_opt=paths=source_relative \
		--go_out=go \
        --validate_out="lang=go:go" \
        --js_out=import_style=commonjs:js \
        --grpc-web_out=import_style=typescript,mode=grpcwebtext:js \
		--doc_out=./doc \
		--doc_opt=markdown,index.md \
        *.proto

	make proto-cleanup

proto-cleanup:
	sed -i '' '/validate/d' js/paginator_pb.d.ts
	sed -i '' '/validate/d' js/paginator_pb.js