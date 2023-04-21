# Get Started

#### Docker Compose

```bash
docker-compose up -d
```

#### Instalar dependências

```bash
go mod tidy
```

#### Executar um arquivo `.go`

```bash
go run main.go
```

#### Compilar um arquivo `.go`

```bash
go build main.go
```

#### Executar um arquivo `.go` compilado

```bash
./main
```

#### Executar Testes

```bash 
cd tests  
go test
```

Obs: Os testes devem estar em um arquivo com o nome `*_test.go` e a função de teste deve começar com `Test`.