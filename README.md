# go-url-shortener <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" height="48px" />

#### Microservice de encurtador de URL baseado no desafio do livro Programando em GO escrito por Caio Filipini.

## .ENV

#### Para conseguir rodar a aplicação sem erros, você deve criar o arquivo .env com as seguintes variáveis

```
# Porta onde a API vai rodar como mostra o exemplo abaixo
PORT=3000
```

## Rodando a aplicação

```bash
# baixar todas as dependencias do microservice
$ go mod tidy

# start
$ go run cmd/app.go

```

## JSON

#### Para gerar um URL encurtada você precisa enviar um arivo JSON no body com a seguinte informação.

```bash
{
  "original": "https://sitequevocequerencurtar.com.br"
}
```

#### Como resposta ele te retorna o site encurtado
