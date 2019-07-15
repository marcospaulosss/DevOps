CONFIGSEVERINO ?= /Users/marcospaulodesousasantos/GoProject/backend/apps/severino
CONFIGELEARNING ?= /Users/marcospaulodesousasantos/GoProject/backend/apps/elearning
CONFIGACCOUNTS ?= /Users/marcospaulodesousasantos/GoProject/backend/apps/accounts

DIRCOVERAGE ?= /Users/marcospaulodesousasantos/GoProject/backend/coverage

test:
	@cd ${CONFIGSEVERINO} && APP_ENV=testing go test -coverprofile=/Users/marcospaulodesousasantos/GoProject/backend/coverage/coverage1.out ./src/... \
	&& cd ${CONFIGELEARNING} && APP_ENV=testing CONFIG_DIR=${CONFIGELEARNING} go test -p 1 -coverprofile=/Users/marcospaulodesousasantos/GoProject/backend/coverage/coverage2.out -failfast ./src/... \
	&& cd ${CONFIGACCOUNTS} && APP_ENV=testing CONFIG_DIR=${CONFIGACCOUNTS} go test -p 1 -coverprofile=/Users/marcospaulodesousasantos/GoProject/backend/coverage/coverage3.out -failfast ./src/... \
	&& cd ${DIRCOVERAGE} && sed '/mode/d' coverage2.out > coverage4.out \
	&& sed '/mode/d' coverage3.out > coverage5.out \
	&& rm coverage.out && cat coverage1.out coverage4.out coverage5.out >> coverage.out \
	&& rm -rf coverage1.out coverage2.out coverage3.out coverage4.out coverage5.out
