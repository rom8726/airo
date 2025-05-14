GOCMD=go
GOBUILD=$(GOCMD) build
GOPROXY=https://proxy.golang.org,direct

.PHONY: test
test: ## Run tests
	@go test -v ./...

.PHONY: build
build: ## Build binary
	@echo "\nBuilding binary..."
	@echo
	go env -w GOPROXY=${GOPROXY}

	CGO_ENABLED=0 $(GOBUILD) -trimpath -o bin/airo ./cmd/airo
