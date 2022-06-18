# go-api

RestApi with golang MongoDB

### 📋 Tecnologias utilizadas

As tecnologias abaixo foram utilizadas:

- [Golang](https://www.golang.org)
- [MongoDB](https://www.mongodb.com/try/download/community)

### 🔧 Instalação

```bash
# Clone este repositório
$ git clone https://github.com/matheuscaet/go-api.git

# Acesse a pasta do projeto no terminal/cmd
$ cd go-api

Crie um arquivo também na raiz, com nome **.env** e cole as variáveis dentro.

O preenchimento dessas variáveis são obrigatórias, então segue um exemplo de uso:

| VARIÁVEL  | VALOR                                                  |
| --------- | ------------------------------------------------------ |
| API_PORT  | 3001                                                   |
| MONGOURI  | mongodb+srv://(USER):(PASS)@cluster0.h5mtb.mongodb.net |

### 🎲 Rodando o Back End (servidor)

```bash
# Execute a aplicação
$ go run main.go

```