
.PHONY: help
help: ## Print this summary of make targets
	@echo "make(1) UTTT targets:"
	@grep -E '^[a-zA-Z_.-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-30s\033[0m  %s\n", $$1, $$2}'

.PHONY: help
build: bin/uttt ## Build UTTT

bin/uttt: cmd/*.go pkg/*.go
	@go build -o bin/uttt cmd/main.go
