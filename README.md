# Golang Rest API Skeleton

## Main Features
- Clean code with clean architecture
- SOLID
- Isolated repository and domains
- Postgres DB or Mysql
- Simple CRUD of Users with soft delete
- DDD
- Strategy
- Docker for Postgres
- Models for each context
- Auth With Passeto Token
- Postman Collection Exported
- Feature Tests (TODO: Unit tests)

## How to use
- Create `.env` from `.env.example` (cp .env.example .env)
- Create `.env` for tests inside `tests/feature` 
- Change boths `.env` to use your postgres configuration
- run `docker-compose up -d`
- run `go run main.go` (server application)
- run `go test ./tests/...` (testing)

## Endpoints
# Users
- FindAll Users GET `localhost:8000/users`
- Store User POST `localhost:8000/users`
- FindBYId User GET `localhost:8000/users/{id}`
- DeleteById User DELETE `localhost:8000/users/{id}`
- Update User PATCH `localhost:8000/users/{id}`

# Auth
- Login Users POST `localhost:8000/auth/login`

Body:
```json
{
    "email": "example@test.com",
    "password": "12345678"
}
```

Response:
```json
{
    "data": "AUTH_TOKEN"
}
```

- Me User GET `localhost:8000/auth/me` <br/>
(Send Auth bearer token)
<br/>
<br/>

# Thanks
- Diogo Gutierre