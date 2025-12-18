# Quality Tools

## Checkers

| Comando | Regla |
|---------|-------|
| check:limits | 120L/120chars |
| check:headers | License header |
| check:density | Max 10 files/folder |
| check:arch | Layer imports |
| check:sqlc | Schema sync |
| check:i18n | tr() usage |
| check:visuals | Palette usage |
| check:assets | Max 2MB |
| check:scenes | Valid refs |
| check:orphans | No dead code |

## Uso

```bash
cd server
task check:all    # Checks Go
task check:godot  # Checks Godot
```

## Agregar Checker

1. `server/tools/checkers/xxx.go`
2. `server/tools/cmd/check-xxx/main.go`
3. Agregar en `Taskfile.yml`

## Referencias

→ [server/tools/](../server/tools/)
