# Get Started

#### Docker Compose

```bash
docker-compose up -d
```

#### Instalar dependências

```bash
go mod tidy
```

#### Executar o projeto

```bash
go run cmd/main.go
```

#### Executar Testes

```bash 
go test ./...
```

Obs: Os testes devem estar em um arquivo com o nome `*_test.go` e a função de teste deve começar com `Test`.