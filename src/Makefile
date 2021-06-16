PROJECT_NAME := "bcg-checkout"
PKG := "github.com/aceslick911/$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)
 
.PHONY: all
 
all: dep test test-coverage build clean
#lint

dep: ## Get the dependencies
	@go mod download

lint: ## Lint Golang files
	@golint -set_exit_status ${PKG_LIST}

vet: ## Run go vet
	@go vet ${PKG_LIST}
 
clean: ## Remove previous build
	@rm -f $(PROJECT_NAME)/public

get-docs:
	go get -u github.com/swaggo/swag/cmd/swag

docs: get-docs
	swag init --dir cmd/api --parseDependency --output docs

build:
## @go build -i -o build/main $(PKG)
	go build -o bin/restapi cmd/api/main.go

ci: docs build

run:
	go run cmd/api/main.go

test:
	# go test -short ${PKG_LIST}
	go test -v ./test/...

test-coverage: test ## Run tests with coverage
	@go test -short -coverprofile cover.out -covermode=atomic ${PKG_LIST} 
	@cat cover.out >> coverage.txt
