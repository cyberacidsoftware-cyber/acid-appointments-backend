# Appointments API in Go

## Get started

```bash
# Select an environment
cp .env.dev .env
cp .env.test .env
```

Up the database using docker:

```bash
# Up and migrate the Database
docker-compose --project-name appointments up db --detach
go generate
```

Up the API using docker:

```bash
# Build and up the API
docker build --tag appointments-api .
docker-compose --project-name appointments up --detach

# (Optional) Feed with fake data
go run ./cmd/cli/main.go run-input -input ./inputs/feed.txt
go run ./cmd/cli/main.go run-input -input ./inputs/feed-extended.txt
```

Or up the API locally:

```bash
# only run
go run ./cmd/api/main.go

# watch
air -c .air.toml
```

Look the **Commands section** to see other possible commands.

## Commands

### Database

```bash
# Database
docker-compose --project-name appointments up db --detach
docker-compose --project-name appointments down db
```

### API

```bash
# dev
go run ./cmd/api/main.go
# dev watch
air -c .air.toml

# docker
docker build --tag appointments-api .
docker run --deamon --port 8080:8080 --name appointments-api appointments-api
# or
docker-compose --project-name appointments up --detach
```

### CLI

```bash
# go run ./cmd/cli/main.go login --email admin@staff.com --password 123mudar # Remove?
go run ./cmd/cli/main.go create-admin --name Admin --email admin@staff.com --password 123mudar
go run ./cmd/cli/main.go create-secretary --name "Secretary" --email secretary@staff.com --password 123mudar --phone 55123456789 --birthdate "2010-10-10" --cpf "45678912398"
go run ./cmd/cli/main.go create-customer -name "Customer" -email customer@people.com -phone 55789456321 -birthdate "2000-01-01" -cpf "12345678954"
go run ./cmd/cli/main.go create-specialist -name "Specialist" -email specialist@people.com -phone 55456789123 -birthdate "2001-02-02" -cpf "78945612354" -cnpj "98765432198722"
go run ./cmd/cli/main.go create-specialization -name "Specialization" 
go run ./cmd/cli/main.go create-service -name "Service" -specialization "Specialization"

go run ./cmd/cli/main.go list-appointments --date=2020-10-10

go run ./cmd/cli/main.go reset-db
```

### Air Tool

```bash
go install github.com/air-verse/air@latest
air init
air
```

## Migration with tern

```bash
# Use the following script to run a command with .env loaded
# go run ./cmd/wenv/main.go

go run ./cmd/wenv/main.go tern init ./internal/infra/migrations
go run ./cmd/wenv/main.go tern new --migrations ./internal/infra/migrations/ create_users_table
go run ./cmd/wenv/main.go tern status --migrations ./internal/infra/migrations/ --config ./internal/infra/migrations/tern.conf
go run ./cmd/wenv/main.go tern migrate --migrations ./internal/infra/migrations/ --config ./internal/infra/migrations/tern.conf
```

### Generating GO code to access the database

```bash
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
sqlc generate -f "./internal/infra/sqlc.yaml"
docker run --rm -v ".:/src" -w /src kjconroy/sqlc generate -f "./internal/infra/sqlc.yaml"
```

### Generating DOCS code from comments

```bash
go install github.com/swaggo/swag/cmd/swag@latest
swag init --dir ./cmd/api/,./internal/api/ --output ./internal/api/docs
docker run --rm -v ".:/code" ghcr.io/swaggo/swag:latest
```

### Static Check

```bash
staticcheck --version
staticcheck ./...
```

### Install new dependencies

```bash
go mod tidy
```

## References

* [Using the flag library to create a CLI program](https://www.digitalocean.com/community/tutorials/how-to-use-the-flag-package-in-go)
* [flag docs #1](https://cli.urfave.org/v2/examples/flags/)
* [flag docs #2](https://pkg.go.dev/flag)
* [Optimizing Docker image](https://www.youtube.com/watch?v=cGYfMIKHC30)
* [Fake generator](https://www.fakenamegenerator.com/gen-random-br-br.php)
* [CNPJ generator](https://www.4devs.com.br/gerador_de_cnpj)
* [HTTP Yac](https://httpyac.github.io/)
* [Github - Gopportunities](https://github.com/arthur404dev/gopportunities)

## Scripts

```js
// https://www.fakenamegenerator.com/gen-random-br-br.php
console.log([
    "Command: create customer",
    document.querySelector(".content .address").childNodes[1].textContent,
    document.querySelectorAll(".content .extra .dl-horizontal")[8].childNodes[3].textContent.split(' ', 1),
    document.querySelectorAll(".content .extra .dl-horizontal")[3].childNodes[3].textContent,
    document.querySelectorAll(".content .extra .dl-horizontal")[5].childNodes[3].textContent,
    document.querySelectorAll(".content .extra .dl-horizontal")[1].childNodes[1].textContent,
].join('\n'))
```