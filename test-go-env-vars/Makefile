SHELL := /usr/bin/env bash -o pipefail

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## build: build this project
.PHONY: build
build:
	go build -o app ./...

## build/image: build the docker image
.PHONY: build/image
build/image: 
	docker buildx build -t myapp .

## deploy: deploy to k8s
.PHONY: deploy
deploy:
	./deploy.sh
