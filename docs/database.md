# Base de Datos

## Auto-Setup

El servidor ejecuta automáticamente:
1. `EnsureDatabase()` - Crea BD si no existe
2. `Runner.Run()` - Aplica migraciones pendientes
3. `Seeder.SeedAdmin()` - Crea usuario admin

## Schema

| Tabla | Columnas |
|-------|----------|
| users | id, email, hashed_password, role, created_at |
| schema_migrations | version, applied_at |

## SQLC

Queries → [server/db/queries.sql](../server/db/queries.sql)

Generado → [server/internal/infra/postgres/sqlc/](../server/internal/infra/postgres/sqlc/)

```bash
cd server && task sqlc:generate
```

## Agregar Migración

1. Crear `server/internal/infra/postgres/migrate/002_xxx.up.sql`
2. Reiniciar servidor (aplica automáticamente)
