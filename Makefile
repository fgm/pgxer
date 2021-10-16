.PHONY: generate
generate:
	go generate ./...

.PHONY: lint
lint:
	staticcheck ./...
