# API

Base: `http://localhost:8080/api/v1`

## POST /auth/login

**Request:**
```json
{"email": "admin@dev.local", "password": "admin123"}
```

**Response 200:**
```json
{"access_token": "...", "refresh_token": "...", "expires_in": 900}
```

**Response 401:** `invalid credentials`

## Tokens

| Token | Duración | Header |
|-------|----------|--------|
| access | 15 min | `Authorization: Bearer <token>` |
| refresh | 7 días | Body en refresh endpoint |

## OpenAPI

→ [server/api/openapi.yaml](../server/api/openapi.yaml)
