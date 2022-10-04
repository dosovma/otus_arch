.PHONY: lint
lint:
	golangci-lint run --config ./.golangci.yml ./...

.PHONY: clean
clean:
	find ./internal/* -type d -name "*mock*" | xargs rm -dfR

.PHONY: generate
generate:
	make clean
	@PATH=${ROOTDIR}/bin:${PATH}:${GOPATH} go generate ./...

.PHONY: test
test:
	make generate
	go test ./... -count=1
