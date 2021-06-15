# PROJECT_NAME := "bcg-checkout"
# PKG := "github.com/aceslick911/$(PROJECT_NAME)"
# PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
# GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)
 
# .PHONY: all dep lint vet test test-coverage build clean
 
all: docs-docker run-docker

dep: ## Get the dependencies
	@go mod download

lint: ## Lint Golang files
	@golint -set_exit_status ${PKG_LIST}

vet: ## Run go vet
	@go vet ${PKG_LIST}

test: ## Run unittests
	@go test -short ${PKG_LIST}

test-coverage: ## Run tests with coverage
	@go test -short -coverprofile cover.out -covermode=atomic ${PKG_LIST} 
	@cat cover.out >> coverage.txt

# build: dep ## Build the binary file
# 	@go build -i -o build/main $(PKG)
 
# clean: ## Remove previous build
# 	@rm -f $(PROJECT_NAME)/build
 
# help: ## Display this help screen
# 	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

get-docs:
	go get -u github.com/swaggo/swag/cmd/swag

docs: get-docs
	swag init --dir cmd/api --parseDependency --output docs

build:
	go build -o bin/restapi cmd/api/main.go

run:
	go run cmd/api/main.go

test:
	go test -v ./test/...


docker:
# DOCKER_BUILDKIT=0
	 docker build . 

docs-docker: docker
	docker build --output type=local,from=/src/docs/,dest=out .

build-docker: docker
	docker build . -t bcg-api --target API_SERVER

run-docker: docker
	docker run -p 3000:3000 bcg-api
