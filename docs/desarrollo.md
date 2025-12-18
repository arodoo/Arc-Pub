# Desarrollo

## Setup

```bash
cd server
cp .env.example .env
go run ./cmd/api
```

## Requisitos

- Go 1.21+
- PostgreSQL 15+
- Task (opcional): `go install github.com/go-task/task/v3/cmd/task@latest`

## Comandos

```bash
task dev           # Servidor
task db:migrate    # Migraciones
task sqlc:generate # Regenerar SQL
task check:all     # Quality checks
```

## Convenciones

| Tipo | Formato |
|------|---------|
| JSON | snake_case |
| Go | CamelCase |
| Max líneas | 120 |
| Max chars | 120 |

## Referencias

- Taskfile → [server/Taskfile.yml](../server/Taskfile.yml)
- Config → [server/.env.example](../server/.env.example)
