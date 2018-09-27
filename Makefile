.PHONY: test  help
.DEFAULT_GOAL := help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

all: build

init: ## Just run once - initialise dependency management
	dep init 

build: ## build binary 
	dep ensure
	go build -o dist/parametertransform main/main.go

install: build
	cp dist/parametertransform /usr/local/bin


test:
	gotest -v .
