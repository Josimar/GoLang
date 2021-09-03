[![codecov](https://codecov.io/gh/Josimar/GoLang/branch/main/graph/badge.svg?token=YCBEFM82RQ)](https://codecov.io/gh/Josimar/GoLang)

[![CircleCI](https://circleci.com/gh/Josimar/GoLang/tree/main.svg?style=svg)](https://circleci.com/gh/Josimar/GoLang/tree/main)

# GoLang

First application with GoLang

## Start Project ##

git clone https://github.com/Josimar/GoLang.git GoLang

## class 17 ##
- [ ] Create directory 17 - Command Line 
- [ ] enter inside it
```
go mod init commandline
```

*get the package*
```
go get github.com/urfave/cli
```

- [ ] Create file main.go
- [ ] Create directory app
- [ ] Create file app.go inside directory app



# 23 - Banco de Dados
## Acessa o diretório criado

## cria um novo go.mod
go mod init banco-de-dados

## adicionar driver do mysql
go get github.com/go-sql-driver/mysql

## rodar o script
go run banco-de-dados.go

# 24 - CRUD Basico
## Acessa o diretório criado

## cria um novo arquivo
go mod init crud

create file: main.go

## Package gorilla/mux implements a request router and dispatcher for matching incoming requests to their respective handler.
go get github.com/gorilla/mux

## rodar o script
go run main.go