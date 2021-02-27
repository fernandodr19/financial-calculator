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

.PHONY: test
test:
	@echo "==> Running Tests"
	go test -race -v ./...
