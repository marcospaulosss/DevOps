> A service for authentication, authorization and send notifications

## Envvars

```sh
PORT=3002
ACCOUNTS_DATABASE_URL=postgres://root:root@127.0.0.1:5432/accounts?sslmode=disable
ACCOUNTS_DATABASE_PORT=5432
ACCOUNTS_DATABASE_USER=root
ACCOUNTS_DATABASE_PASSWORD=root
ACCOUNTS_DATABASE_NAME=accounts
AWS_SMTP_USER=AKIAUMNY5JYFCIDZPN37
AWS_SMTP_SECRET=BOUAp2Bky8QD2G6Fl2ssesitgxJQUwN6g2QW+2tjI7hS
TWILIO_ID=ACf00c39b91847c3057ddeb91c21182f92
TWILIO_SECRET=77d04d6852cb261f63ca977477c57c9a
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

