.PHONY: test

test:
	@go test -v .

fmt:
	@go fmt ./...