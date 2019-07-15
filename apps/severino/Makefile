.PHONY: docs mocks
test:
	APP_ENV=testing CONFIG_DIR=${PWD} go test -v -cover -count 1 -failfast ./src/...

tests: test

run:
	APP_ENV=development GRPC_GO_LOG_SEVERITY_LEVEL=debug go run main.go

mocks:
	rm -rf mocks 2>/dev/null && mockery -all -dir src/interfaces
	mockery -all -dir ../../proto

docs:
	@docker run --rm -v ${PWD}/docs:/docs gustavohenrique/aglio --theme-full-width -i /docs/index.apib -o /docs/index.html

linux:
	@GOOS=linux go build -ldflags="-s -w" main.go

clearcache:
	@rm -rf ${GOPATH}/pkg/mod
