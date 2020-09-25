.POSIX:

.PHONY: help
help: ## Show this help
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: fmt
fmt: ## Run all formatings
	go mod vendor
	go mod tidy
	go fmt ./...

.PHONY: test
test: ## Run all test
	go test -v ./...

.PHONY: coverage
coverage: ## Show test coverage
	@go test -coverprofile=coverage.out ./... > /dev/null
	go tool cover -func=coverage.out
	@rm coverage.out

.PHONY: docs
docs: ## Generate documentation
	go run docs/gen.go

.PHONY: godoc
godoc: ## Start local godoc server
	@echo "See Documentation:"
	@echo "    http://localhost:6060/pkg/github.com/erdaltsksn/gh-label"
	@echo ""
	@godoc -http=:6060

.PHONY: clean
clean: ## Clean all generated files
	rm -rf ./vendor/
	rm -rf go.sum
