# Mini e-Wallet
BE PRETEST DCI (Mini e-wallet)

## Setup local development

### Install tools

- [Golang](https://golang.org/)
- [Migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

```bash
MacOS
$ brew install golang-migrate
Windows
$ scoop install migrate
Linux
$ curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
$ echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list
$ apt-get update
$ apt-get install -y migrate

```

### Setup infrastructure

- Start postgres
- Create miniewallet database
- Update .env file
- Run db migration:

    ```bash
    migrate -path db/migrations -database "postgresql://user:password@localhost:5432/database?sslmode=disable" -verbose up
    ```

### How to run

- Run server:

    ```bash
    go run main.go
    ```
- Postman:

    ```bash
    https://github.com/rizalaslams/miniewallet/blob/master/mini%20e-wallet.postman_collection.json
    ```