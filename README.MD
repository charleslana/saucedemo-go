# Saucedemo Go

## Inicializar

```bash
go mod init example.com/v2
```

## Instalar dependências

```bash
go mod tidy
```

## Instalar o chromedp

```bash
go get -u github.com/chromedp/chromedp
```

## Executar os testes

```bash
go run main.go
```

## Executar os testes com a pasta especifica

```bash
go test ./tests/specs
```

## Executar os testes com verbosidade

```bash
go test -v ./tests/specs
```
