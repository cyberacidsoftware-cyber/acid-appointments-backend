# Initial Setup Script for Acid Appointments Backend
# This script sets up the complete development environment on Windows

param(
    [string]$AdminEmail = "admin@staff.com",
    [string]$AdminPassword = "123mudar",
    [string]$AdminName = "Admin"
)

Write-Host "================================" -ForegroundColor Cyan
Write-Host "Acid Appointments Backend Setup" -ForegroundColor Cyan
Write-Host "================================" -ForegroundColor Cyan
Write-Host ""

# Function to check if command exists
function Test-CommandExists {
    param($command)
    $oldPreference = $ErrorActionPreference
    $ErrorActionPreference = 'stop'
    try {
        if (Get-Command $command) {
            return $true
        }
    }
    catch {
        return $false
    }
    finally {
        $ErrorActionPreference = $oldPreference
    }
}

# Check prerequisites
Write-Host "Checking prerequisites..." -ForegroundColor Blue

if (-not (Test-CommandExists "docker")) {
    Write-Host "ERROR: Docker is not installed or not in PATH" -ForegroundColor Red
    Write-Host "Please install Docker Desktop and restart this script" -ForegroundColor Red
    exit 1
}

if (-not (Test-CommandExists "docker-compose")) {
    Write-Host "ERROR: Docker Compose is not installed or not in PATH" -ForegroundColor Red
    Write-Host "Please install Docker Compose and restart this script" -ForegroundColor Red
    exit 1
}

if (-not (Test-CommandExists "go")) {
    Write-Host "ERROR: Go is not installed or not in PATH" -ForegroundColor Red
    Write-Host "Please install Go and restart this script" -ForegroundColor Red
    exit 1
}

Write-Host "SUCCESS: All prerequisites found!" -ForegroundColor Green
Write-Host ""

# Check if we're in the correct directory
if (-not (Test-Path "go.mod")) {
    Write-Host "ERROR: This script must be run from the project root directory" -ForegroundColor Red
    Write-Host "Please navigate to the acid-appointments-backend directory and run this script again" -ForegroundColor Red
    exit 1
}

Write-Host "SUCCESS: Running from correct directory" -ForegroundColor Green
Write-Host ""

# Step 1: Setup environment file
Write-Host "Step 1: Setting up environment configuration..." -ForegroundColor Blue

if (-not (Test-Path ".env")) {
    if (Test-Path ".env.dev") {
        Copy-Item ".env.dev" ".env"
        Write-Host "SUCCESS: Created .env file from .env.dev" -ForegroundColor Green
    } else {
        Write-Host "ERROR: .env.dev file not found!" -ForegroundColor Red
        exit 1
    }
} else {
    Write-Host "SUCCESS: .env file already exists" -ForegroundColor Green
}
Write-Host ""

# Step 2: Start database
Write-Host "Step 2: Starting PostgreSQL database..." -ForegroundColor Blue

try {
    # Stop any existing containers first
    docker-compose --project-name appointments down 2>$null
    
    # Start the database
    docker-compose --project-name appointments up db --detach
    if ($LASTEXITCODE -eq 0) {
        Write-Host "SUCCESS: Database container started successfully" -ForegroundColor Green
    } else {
        throw "Failed to start database container"
    }
} catch {
    Write-Host "ERROR: Failed to start database: $_" -ForegroundColor Red
    exit 1
}

# Wait for database to be ready
Write-Host "Waiting for database to be ready..." -ForegroundColor Yellow
Start-Sleep -Seconds 5

# Test database connection
$maxRetries = 10
$retryCount = 0
$dbReady = $false

while ($retryCount -lt $maxRetries -and -not $dbReady) {
    try {
        $result = docker exec appointments-db-1 pg_isready -U postgres -d appointments 2>$null
        if ($LASTEXITCODE -eq 0) {
            $dbReady = $true
            Write-Host "SUCCESS: Database is ready!" -ForegroundColor Green
        } else {
            $retryCount++
            Write-Host "Database not ready yet, retrying... ($retryCount/$maxRetries)" -ForegroundColor Yellow
            Start-Sleep -Seconds 3
        }
    } catch {
        $retryCount++
        Write-Host "Database not ready yet, retrying... ($retryCount/$maxRetries)" -ForegroundColor Yellow
        Start-Sleep -Seconds 3
    }
}

if (-not $dbReady) {
    Write-Host "ERROR: Database failed to start within expected time" -ForegroundColor Red
    exit 1
}
Write-Host ""

# Step 3: Create database tables
Write-Host "Step 3: Creating database tables..." -ForegroundColor Blue

# Check if setup_database.sql exists, if not create it
if (-not (Test-Path "setup_database.sql")) {
    Write-Host "Creating database setup script..." -ForegroundColor Yellow
    
    $setupScript = @"
-- Create extension first
CREATE EXTENSION IF NOT EXISTS unaccent;

-- 1. Create admins table
CREATE TABLE IF NOT EXISTS "admins" (
    "id"                UUID         PRIMARY KEY  NOT NULL   DEFAULT gen_random_uuid(),
    "name"              VARCHAR(255)              NOT NULL,
    "email"             VARCHAR(255)              NOT NULL,
    "password"          VARCHAR(255)              NOT NULL
);
CREATE UNIQUE INDEX IF NOT EXISTS "idx_admins_email" ON admins("email");

-- 2. Create secretaries table
CREATE TABLE IF NOT EXISTS secretaries (
    "id"                UUID         PRIMARY KEY  NOT NULL   DEFAULT gen_random_uuid(),
    "name"              VARCHAR(255)              NOT NULL,
    "email"             VARCHAR(255)              NOT NULL,
    "password"          VARCHAR(255)              NOT NULL,
    "phone"             VARCHAR(255)              NOT NULL,
    "birthdate"         DATE                      NULL,
    "cpf"               VARCHAR(11)               NOT NULL,
    "cnpj"              VARCHAR(14)               NULL
);
CREATE UNIQUE INDEX IF NOT EXISTS secretaries_email_idx ON secretaries("email");

-- 3. Create customers table
CREATE TABLE IF NOT EXISTS customers (
    "id"                UUID         PRIMARY KEY  NOT NULL   DEFAULT gen_random_uuid(),
    "name"              VARCHAR(255)              NOT NULL,
    "email"             VARCHAR(255)              NULL,
    "phone"             VARCHAR(255)              NOT NULL,
    "birthdate"         DATE                      NULL,
    "cpf"               VARCHAR(11)               NOT NULL
);
CREATE UNIQUE INDEX IF NOT EXISTS customers_cpf_idx ON customers("cpf");

-- 4. Create specialists table
CREATE TABLE IF NOT EXISTS specialists (
    "id"                UUID         PRIMARY KEY  NOT NULL   DEFAULT gen_random_uuid(),
    "name"              VARCHAR(255)              NOT NULL,
    "email"             VARCHAR(255)              NOT NULL,
    "phone"             VARCHAR(255)              NOT NULL,
    "birthdate"         DATE                      NULL,
    "cpf"               VARCHAR(11)               NOT NULL,
    "cnpj"              VARCHAR(14)               NULL
);
CREATE UNIQUE INDEX IF NOT EXISTS specialists_email_idx ON specialists("email");

-- 5. Create specializations table
CREATE TABLE IF NOT EXISTS specializations (
    "id"                UUID         PRIMARY KEY  NOT NULL   DEFAULT gen_random_uuid(),
    "name"              VARCHAR(255)              NOT NULL
);
CREATE UNIQUE INDEX IF NOT EXISTS specializations_name_idx ON specializations("name");

-- 6. Create service_names table
CREATE TABLE IF NOT EXISTS service_names (
    "id"                UUID         PRIMARY KEY  NOT NULL   DEFAULT gen_random_uuid(),
    "name"              VARCHAR(255)              NOT NULL,
    "specialization_id" UUID                      NOT NULL,
    FOREIGN KEY ("specialization_id") REFERENCES specializations("id") ON DELETE CASCADE
);
CREATE UNIQUE INDEX IF NOT EXISTS service_names_name_idx ON service_names("name");
CREATE INDEX IF NOT EXISTS service_names_specialization_id_idx ON service_names("specialization_id");

-- 7. Create services table
CREATE TABLE IF NOT EXISTS services (
    "id"                UUID         PRIMARY KEY  NOT NULL   DEFAULT gen_random_uuid(),
    "service_name_id"   UUID                      NOT NULL,
    "specialist_id"     UUID                      NOT NULL,
    "price"             INTEGER                   NOT NULL,
    "duration"          INTERVAL                  NOT NULL,
    FOREIGN KEY ("service_name_id") REFERENCES service_names("id") ON DELETE CASCADE,
    FOREIGN KEY ("specialist_id") REFERENCES specialists("id") ON DELETE CASCADE
);
CREATE UNIQUE INDEX IF NOT EXISTS services_service_unique_idx ON services("service_name_id", "specialist_id");
CREATE INDEX IF NOT EXISTS services_specialist_id_idx ON services("specialist_id");
CREATE INDEX IF NOT EXISTS services_service_name_id_idx ON services("service_name_id");

-- 8. Create specialist_hours table
CREATE TABLE IF NOT EXISTS specialist_hours (
    "id"                UUID         PRIMARY KEY  NOT NULL   DEFAULT gen_random_uuid(),
    "specialist_id"     UUID                      NOT NULL,
    "weekday"           INTEGER                   NOT NULL,
    "start_time"        TIME                      NOT NULL,
    "end_time"          TIME                      NOT NULL,
    FOREIGN KEY ("specialist_id") REFERENCES specialists("id") ON DELETE CASCADE
);

-- 9. Create appointments table
CREATE TABLE IF NOT EXISTS appointments (
    "id"                UUID         PRIMARY KEY  NOT NULL   DEFAULT gen_random_uuid(),
    "customer_id"       UUID                      NOT NULL,
    "specialist_id"     UUID                      NOT NULL,
    "service_name_id"   UUID                      NOT NULL,
    "price"             INTEGER                   NOT NULL,
    "duration"          INTERVAL                  NOT NULL,
    "date"              DATE                      NOT NULL,
    "time"              TIME                      NOT NULL,
    "status"            INTEGER                   NOT NULL   DEFAULT 0,
    "notified_at"       TIMESTAMPTZ               NULL,
    "notified_by"       UUID                      NULL,
    FOREIGN KEY ("customer_id") REFERENCES customers("id") ON DELETE CASCADE,
    FOREIGN KEY ("specialist_id") REFERENCES specialists("id") ON DELETE CASCADE,
    FOREIGN KEY ("service_name_id") REFERENCES service_names("id") ON DELETE CASCADE
);
CREATE INDEX IF NOT EXISTS appointments_date_idx ON appointments("date");
"@
    
    $setupScript | Out-File -FilePath "setup_database.sql" -Encoding UTF8
    Write-Host "SUCCESS: Database setup script created" -ForegroundColor Green
}

try {
    Get-Content "setup_database.sql" | docker exec -i appointments-db-1 psql -U postgres -d appointments
    if ($LASTEXITCODE -eq 0) {
        Write-Host "SUCCESS: Database tables created successfully" -ForegroundColor Green
    } else {
        throw "Failed to create database tables"
    }
} catch {
    Write-Host "ERROR: Failed to create database tables: $_" -ForegroundColor Red
    exit 1
}
Write-Host ""

# Step 4: Create admin user
Write-Host "Step 4: Creating admin user..." -ForegroundColor Blue

try {
    $output = go run ./cmd/cli/main.go create-admin --name $AdminName --email $AdminEmail --password $AdminPassword 2>&1
    if ($LASTEXITCODE -eq 0) {
        Write-Host "SUCCESS: Admin user created successfully!" -ForegroundColor Green
        Write-Host "Email: $AdminEmail" -ForegroundColor White
        Write-Host "Password: $AdminPassword" -ForegroundColor White
    } else {
        # Check if user already exists
        if ($output -match "already exists" -or $output -match "duplicate") {
            Write-Host "INFO: Admin user already exists" -ForegroundColor Yellow
            Write-Host "Email: $AdminEmail" -ForegroundColor White
            Write-Host "Password: $AdminPassword" -ForegroundColor White
        } else {
            throw "Failed to create admin user: $output"
        }
    }
} catch {
    Write-Host "ERROR: Failed to create admin user: $_" -ForegroundColor Red
    exit 1
}
Write-Host ""

# Step 5: Verify setup
Write-Host "Step 5: Verifying setup..." -ForegroundColor Blue

try {
    $tables = docker exec appointments-db-1 psql -U postgres -d appointments -t -c "\dt" 2>$null
    if ($LASTEXITCODE -eq 0 -and $tables -match "admins") {
        Write-Host "SUCCESS: Database verification successful" -ForegroundColor Green
    } else {
        throw "Database verification failed"
    }
} catch {
    Write-Host "ERROR: Database verification failed: $_" -ForegroundColor Red
    exit 1
}
Write-Host ""

# Success message
Write-Host "Setup completed successfully!" -ForegroundColor Green
Write-Host ""
Write-Host "================================" -ForegroundColor Cyan
Write-Host "SETUP SUMMARY" -ForegroundColor Cyan
Write-Host "================================" -ForegroundColor Cyan
Write-Host "Environment file: .env" -ForegroundColor White
Write-Host "Database: PostgreSQL running on port 5432" -ForegroundColor White
Write-Host "Database tables: All created" -ForegroundColor White
Write-Host "Admin user: $AdminEmail" -ForegroundColor White
Write-Host ""
Write-Host "LOGIN CREDENTIALS:" -ForegroundColor Yellow
Write-Host "   Email: $AdminEmail" -ForegroundColor White
Write-Host "   Password: $AdminPassword" -ForegroundColor White
Write-Host ""
Write-Host "NEXT STEPS:" -ForegroundColor Yellow
Write-Host "   1. Import 'Acid_Appointments_API.json' into Postman/Insomnia" -ForegroundColor White
Write-Host "   2. Start the API server using the command below" -ForegroundColor White
Write-Host "   3. Test the /health endpoint" -ForegroundColor White
Write-Host "   4. Login using the credentials above" -ForegroundColor White
Write-Host ""

# Ask user to start the server
do {
    $response = Read-Host "Would you like to start the API server now? (y/n)"
    $response = $response.ToLower()
} while ($response -ne "y" -and $response -ne "n" -and $response -ne "yes" -and $response -ne "no")

if ($response -eq "y" -or $response -eq "yes") {
    Write-Host ""
    Write-Host "Starting API server..." -ForegroundColor Blue
    Write-Host "The server will start on http://localhost:8080" -ForegroundColor Yellow
    Write-Host "Press Ctrl+C to stop the server" -ForegroundColor Yellow
    Write-Host ""
    
    try {
        go run ./cmd/api/main.go
    } catch {
        Write-Host "Server stopped or interrupted" -ForegroundColor Yellow
    }
} else {
    Write-Host ""
    Write-Host "To start the API server manually, run:" -ForegroundColor Yellow
    Write-Host "   go run ./cmd/api/main.go" -ForegroundColor Cyan
    Write-Host ""
    Write-Host "The server will be available at: http://localhost:8080" -ForegroundColor White
}

Write-Host ""
Write-Host "Thank you for using Acid Appointments Backend!" -ForegroundColor Cyan