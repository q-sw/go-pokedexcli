help:
	@printf "Available targets:\n"
	@grep -E '^[1-9a-zA-Z_-]+:.*?## .*$$|(^#--)' $(MAKEFILE_LIST) \
	| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m %-43s\033[0m %s\n", $$1, $$2}' \
	| sed -e 's/\[32m #-- /[33m/'

default: help

build: ## Build Go project
	go build -o bin/pokedexcli

run: build ## Run Package build
	./bin/pokedexcli