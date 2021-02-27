OS ?= linux
CMD ?= github.com/fernandodr19/financial-calculator/cmd
GIT_COMMIT=$(shell git rev-parse HEAD)
GIT_BUILD_TIME=$(shell date '+%Y-%m-%d__%I:%M:%S%p')

.PHONY: compile
compile: clean
	@echo "==> Go Building"
	@env GOOS=${OS} GOARCH=amd64 go build -v -o  build/calculator \
	-ldflags "-X main.BuildGitCommit=$(GIT_COMMIT) -X main.BuildTime=$(GIT_BUILD_TIME)" ./cmd/

.PHONY: clean
clean:
	@echo "==> Cleaning releases"
	@GOOS=${OS} go clean -i -x ./...
	@rm -f build/*

.PHONY: lint
lint:
	@echo "==> Running linter"
ifeq (, $(shell which $$(go env GOPATH)/bin/golangci-lint))
	@echo "==> installing golangci-lint"
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin
	go install ./...
	go test -i ./...
endif
	@$$(go env GOPATH)/bin/golangci-lint run -c ./.golangci.yml ./...

.PHONY: test
test:
	@echo "==> Running tests"
	go test -race -v ./...

.PHONY: test-coverage
test-coverage:
	@echo "==> Running tests coverage"
	@richgo test -failfast -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html