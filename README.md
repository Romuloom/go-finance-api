# go-finance-api

API REST em Go para controle financeiro pessoal. Projeto de estudo, em desenvolvimento, focado em organizar um serviço Go em camadas desde o início.

## Arquitetura

```
cmd/server/          ponto de entrada da aplicação
internal/config/     configuração do servidor HTTP
internal/database/   conexão com PostgreSQL via GORM
internal/handlers/   camada HTTP (requisição e resposta)
internal/services/   regras de negócio e validação
internal/repositories/ acesso a dados
internal/models/     entidades
internal/routes/     registro das rotas
```

Cada camada depende apenas da seguinte (handler → service → repository), o que facilita testar e trocar a persistência.

## Stack

Go · net/http · GORM · PostgreSQL 16 · Docker Compose

## Endpoints

| Método | Rota      | Descrição          |
|--------|-----------|--------------------|
| GET    | `/health` | Verificação de saúde |
| GET    | `/users`  | Lista usuários     |
| POST   | `/users`  | Cria usuário (com validação) |

Erros são retornados em JSON com formato padronizado.

## Como rodar

```bash
docker compose up -d      # sobe o PostgreSQL
go run ./cmd/server       # API em http://localhost:8081
```

## Próximos passos

- Entidades de contas, categorias e transações
- Autenticação com JWT
- Testes de serviço e de repositório
