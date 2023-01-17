SOURCE=./...

.DEFAULT_GOAL := test

test: vet test-fmt
	go test $(SOURCE) -v -coverprofile=coverage.out
.PHONY: test

vet:
	go vet $(SOURCE)
.PHONY: vet

test-fmt:
	test -z $(shell go fmt $(SOURCE))
.PHONY: test-fmt
