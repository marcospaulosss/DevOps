> An API Gateway like exposing REST API and forwarding data to others services via gRPC

## Envvars

```sh
PORT=4001
ELEARNING_URL=localhost:3001
ACCOUNT_URL=localhost:3002
TOKEN_EXP=5  # TEST ONLY: set expiration token period in seconds
TOKEN_SECRET=mywordismypassword
```

## Setup

```sh
ELEARNING_URL=localhost:3001 PORT=4001 make run  # it will download all dependencies and starts up the server
```

## Testing

```sh
go get github.com/vektra/mockery/cmd/mockery
make mocks
make tests
```


