> A service running gRPC server for elearning stuff

## Envvars

```sh
PORT=3001
ELEARNING_DATABASE_URL=postgres://root:root@127.0.0.1:5432/elearning?sslmode=disable
ELEARNING_DATABASE_PORT=5432
ELEARNING_DATABASE_USER=root
ELEARNING_DATABASE_PASSWORD=root
ELEARNING_DATABASE_NAME=elearning
```

## Setup

```
cp .env.sample .env  # optional if you wish to use non default values
make createdb        # it creates a postgres container
make seed            # pre populate the database
make run             # download all dependencies and run the server
```

## Testing

```
make tests
```

## Development

```
make psql  # open the postgres client for database management
```

