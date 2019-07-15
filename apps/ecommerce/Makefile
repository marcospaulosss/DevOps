test:
	APP_ENV=testing CONFIG_DIR=${PWD} go test -v -cover -failfast ./src/...

tests: test

run:
	GRPC_GO_LOG_VERBOSITY_LEVEL=99 GRPC_GO_LOG_SEVERITY_LEVEL=info APP_ENV=development go run main.go

createdb: dropdb
	@docker run -d --name postgres \
		-e POSTGRES_USER=root \
		-e POSTGRES_PASSWORD=root \
		-e POSTGRES_DB=elearning \
		-e ADDITIONAL_DB=accounts,ecommerce \
		-p 5432:5432 estrategiaconcursos/postgres:alpine

dropdb:
	@docker rm -f postgres 2>/dev/null || exit 0

seed:
	@docker run --rm --link postgres:server -v ${PWD}/../../schemas:/schemas postgres sh -c "PGPASSWORD=root psql -U root -h server -d ecommerce < /schemas/ecommerce.sql"

psql:
	@docker exec -it postgres sh -c "PGPASSWORD=root psql -U root -d ecommerce"

compile:
	@GOOS=linux go build -ldflags="-s -w" main.go

build: compile

