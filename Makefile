help: ## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

build: ## builds tracker
	go build ./

fmt: ## format all Go code
	go fmt ./...

test: ## run unit tests
	go test ./...
