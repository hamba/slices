include github.com/hamba/make/golang

generate:
	@go generate
.PHONY: generate

bench:
	@go test -bench=. .
.PHONY: bench
