RELEASE_DIR := dist
LDFLAGS := -ldflags "-s -w"

.PHONY: clean
clean: ## Remove release binaries
	rm -rf $(RELEASE_DIR)

.PHONY: build-dirs
build-dirs: clean
	mkdir -p $(RELEASE_DIR)

.PHONY: build
build: build-dirs
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build $(LDFLAGS) -tags=jsoniter -o ./dist/cache-test