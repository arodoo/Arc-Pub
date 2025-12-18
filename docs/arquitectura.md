# Arquitectura

## Patrón: DDD + Hexagonal

```
HTTP (Chi) → Application (Use Cases) → Domain → Infrastructure
```

## Capas

| Capa | Ubicación | Dependencias |
|------|-----------|--------------|
| domain | `server/internal/domain/` | Ninguna |
| application | `server/internal/application/` | domain |
| infra | `server/internal/infra/` | application, domain |

## Flujo Login

```
POST /api/v1/auth/login
  → Handler.Login()
  → LoginUseCase.Execute()
    → UserRepository.FindByEmail()
    → PasswordHasher.Compare()
    → TokenService.GeneratePair()
  → { access_token, refresh_token }
```

## Referencias

- [server/internal/domain/user/user.go](../server/internal/domain/user/user.go)
- [server/internal/application/auth/](../server/internal/application/auth/)
- [server/internal/infra/postgres/](../server/internal/infra/postgres/)
