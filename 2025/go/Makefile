DIR := ${CURDIR}

.PHONY: help
help:
	@printf "\033[33mUsage:\033[0m\n  make TARGET\n\n\033[32m#\n# Commands\n#---------------------------------------------------------------------------\033[0m\n\n"
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//' | awk 'BEGIN {FS = ":"}; {printf "\033[33m%s:\033[0m%s\n", $$1, $$2}'

.PHONY: run
run: ## run solvers for all days
	go run main.go

.PHONY: test
test: ## run all tests (cached)
	go test ./...

.PHONY: benchmark
benchmark: ## run all benchmarks
	go test -bench=. ./...
