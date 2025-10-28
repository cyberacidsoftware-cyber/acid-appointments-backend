# Appointments API in Go

## 🚀 Quick Start (Windows)

For Windows users, we provide an automated setup script that handles everything:

### **Option 1: Automated Setup (Recommended)**

```powershell
# Run the setup script (handles execution policy automatically)
powershell -ExecutionPolicy Bypass -File .\setup.ps1

# Or double-click setup.bat for automatic execution
```

**What the script does:**
- ✅ Checks prerequisites (Docker, Go, Docker Compose)
- ✅ Creates environment configuration (.env)
- ✅ Starts PostgreSQL database container
- ✅ Creates all database tables
- ✅ Creates admin user (admin@staff.com / 123mudar)
- ✅ Verifies setup completion
- ✅ Optionally starts the API server

**If you get execution policy errors:**
```powershell
# Option 1: Run with bypass (recommended)
powershell -ExecutionPolicy Bypass -File .\setup.ps1

# Option 2: Change policy temporarily (as Administrator)
Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser
.\setup.ps1
Set-ExecutionPolicy -ExecutionPolicy Restricted -Scope CurrentUser
```

### **Option 2: Manual Setup**

If you prefer manual setup or are on Linux/Mac:

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

## 📋 API Testing

After setup completion, you can test the API using the provided collection:

### **Import API Collection**
1. Import `Acid_Appointments_API.json` into **Postman** or **Insomnia**
2. Set the base URL to `http://localhost:8080`
3. Use the login endpoint with default credentials:
   - **Email:** `admin@staff.com`
   - **Password:** `123mudar`

### **Testing Workflow**
1. **Health Check** - Verify API is running
2. **Login** - Get access token
3. **Create Data** - Add specializations, services, specialists, customers
4. **Manage Appointments** - Schedule and manage appointments

The collection includes **50+ endpoints** with example requests for all functionality.

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

## 🛠️ Troubleshooting

### **Common Issues**

#### PowerShell Execution Policy
```powershell
# Error: "execution of scripts is disabled on this system"
# Solution: Use bypass policy
powershell -ExecutionPolicy Bypass -File .\setup.ps1
```

#### Docker Issues
```bash
# Error: "docker: command not found"
# Solution: Install Docker Desktop for Windows

# Error: "port 5432 already in use"
# Solution: Stop existing PostgreSQL service or change port in .env
```

#### Database Connection Issues
```bash
# Error: "relation does not exist"
# Solution: Run the setup script or manually create tables
Get-Content setup_database.sql | docker exec -i appointments-db-1 psql -U postgres -d appointments
```

#### Go Module Issues
```bash
# Error: Go module issues
# Solution: Clean and reinstall dependencies
go clean -modcache
go mod download
go mod tidy
```

### **Manual Database Reset**
```bash
# Stop containers
docker-compose --project-name appointments down

# Remove volumes (WARNING: This deletes all data)
docker volume rm appointments_db

# Restart setup
.\setup.ps1
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