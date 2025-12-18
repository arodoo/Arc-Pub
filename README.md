# Arc-Pub

Metaverso 2D MMO Social.

## Estructura

```
Arc-Pub/
├── docs/           # Documentación centralizada
├── server/         # Backend Go (Chi + PostgreSQL)
├── client/         # Cliente Godot (futuro)
└── instructions/   # Reglas del proyecto
```

## Documentación

→ [docs/](docs/README.md)

## Quick Start

```bash
cd server
cp .env.example .env
go run ./cmd/api
```

## Stack

| Capa | Tecnología |
|------|------------|
| Server | Go + Chi |
| Client | Godot 4 + GDScript |
| DB | PostgreSQL |
| Protocol | WebTransport / WebSockets |
| Format | Protocol Buffers |
