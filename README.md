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

### Postman Collection
You can get postman collection [here](https://github.com/rizalaslams/miniewallet/blob/master/mini%20e-wallet.postman_collection.json).

## Other
------------
`Mini e-Wallet` based on following plugins or services:

+ [Gin](https://github.com/gin-gonic/gin)
+ [GORM](github.com/jinzhu/gorm)
+ [jwt-go](github.com/dgrijalva/jwt-go)
+ [GoDotEnv ](github.com/joho/godotenv)
+ and other
