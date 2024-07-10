#!/bin/bash

# Set the project name
PROJECT_NAME="panchangapi-gateway"

# Create the main project directory
mkdir -p $PROJECT_NAME
cd $PROJECT_NAME

# Create the directory structure
mkdir -p cmd/server
mkdir -p internal/{api/{handlers,routes},auth,config,database/migrations,models,utils}
mkdir -p pkg/validator
mkdir -p tests/{integration,unit}

# Create main.go
cat << EOF > cmd/server/main.go
package main

import "fmt"

func main() {
    fmt.Println("PanchangAPI Gateway")
}
EOF

# Create other Go files
touch internal/api/handlers/user_handlers.go
touch internal/api/routes.go
touch internal/auth/auth.go
touch internal/config/config.go
touch internal/database/database.go
touch internal/models/user.go
touch internal/utils/{apikey.go,password.go}
touch pkg/validator/validator.go
touch tests/integration/registration_test.go
touch tests/unit/user_test.go

# Create SQL migration file
touch internal/database/migrations/001_create_users_table.sql

# Create .env file
cat << EOF > .env
# Environment variables
DATABASE_URL=postgres://username:password@localhost:5432/database_name
EOF

# Create .gitignore
cat << EOF > .gitignore
# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, built with 'go test -c'
*.test

# Output of the go coverage tool, specifically when used with LiteIDE
*.out

# Dependency directories (remove the comment below to include it)
# vendor/

# Environment variables
.env

# IDE specific files
.vscode/
.idea/

# OS generated files
.DS_Store
.DS_Store?
._*
.Spotlight-V100
.Trashes
ehthumbs.db
Thumbs.db
EOF

# Create README.md
cat << EOF > README.md
# PanchangAPI Gateway

This project serves as a gateway for the PanchangAPI, providing user registration, authentication, and data proxying services.

## Getting Started

1. Clone the repository
2. Set up your .env file with the necessary environment variables
3. Run \`go mod init $PROJECT_NAME\` to initialize the Go module
4. Run \`go mod tidy\` to download dependencies
5. Run \`go run cmd/server/main.go\` to start the server

## Project Structure

[Add structure description here]

## API Endpoints

[Add API endpoint descriptions here]

## Contributing

[Add contribution guidelines here]

## License

[Add license information here]
EOF

# Initialize Go module
go mod init $PROJECT_NAME

echo "Project structure for $PROJECT_NAME has been set up!"