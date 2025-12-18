---
description: Validar sintaxis GDScript y ver errores
---

# Validate Godot GDScript

// turbo-all

## 1. Validar sintaxis GDScript (como tsc --noEmit)
```bash
cd d:\zProyectos\Godot\Arc-Pub
gdparse client/scripts/**/*.gd
```

## 2. Validar proyecto completo con Godot headless
```bash
"C:\Program Files\godot\Godot_v4.5.1-stable_win64.exe" --headless --path ./client --quit 2>&1
```

## 3. Si hay errores, se mostrarán como:
```
ERROR: Failed to load script "res://path/file.gd" with error...
  at: load (modules/gdscript/gdscript.cpp:3041)
```

## Notas
- `gdparse` = validación rápida de sintaxis
- `godot --headless` = validación completa (scenes, autoloads, etc)
- Instalar gdtoolkit: `pip install gdtoolkit`
