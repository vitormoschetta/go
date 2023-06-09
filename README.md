# Get Started

#### Create Database

```bash
docker-compose up -d
```

#### Install dependencies

```bash
go mod tidy
```

#### Run Project

```bash
go run cmd/main.go
```

#### Run Tests

```bash 
go test ./...
```

Test coverage:
```bash
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

#### Run Linter

Check for linting errors (code style, code smells, etc.)
```bash
golangci-lint run
```