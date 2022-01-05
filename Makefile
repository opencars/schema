PKG_PATH=github.com/opencars/schema

.PHONY: proto
proto: generate

generate:
	@docker run --rm -v ${PWD}:/proto opencars/protoc ${PKG_PATH} ./

format:
	@echo 'Formatting the code...'
	@gofmt -w .
	@goimports -local "${PKG_PATH}" -w .

lint-golangci:
	@echo 'Linting with golangci...'
	@golangci-lint run --timeout=1h ./...

lint: lint-golangci
