
.PHONY: help
help: ## Print this summary of make targets
	@echo "GNU make(1) UTTT targets:"
	@grep -E '^[a-zA-Z_.-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-30s\033[0m  %s\n", $$1, $$2}'

.PHONY: help
build: uttt ## Build UTTT

uttt: cmd/*.go pkg/**/*.go
	@go build -o uttt cmd/main.go
