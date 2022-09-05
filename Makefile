.PHONY: lint
lint:
	golangci-lint run --config ./.golangci.yml ./...

.PHONY: test
test:
	go test ./... -count=1
