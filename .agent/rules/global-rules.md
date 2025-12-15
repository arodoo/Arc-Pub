---
trigger: always_on
---

---
applyTo: '**'
---

### MUSTS
- **Quality over speed**
- **Single Source of Truth**: Never duplicate logic or data definition.
- **Strict Typing**: No dynamic types allowed unless strictly necessary for external APIs.

## Global Considerations
- **Persistence**: Database is NEVER dropped, only migrated.
- **Protocol**: 
    - HTTP (Login/Market) -> Defined via OpenAPI.
    - Sockets (Game) -> Defined via Protobufs.
- **Security**: Server is Authoritative. Client is just a viewer.

## Architecture Standards

### Backend (Go - Chi)
- **Pattern**: Standard Go Project Layout.
- **Naming**: `snake_case` for JSON/DB, `CamelCase` for Structs.
- **Error Handling**: Explicit check `if err != nil`. No underscore `_` to ignore errors.

### Frontend (Godot - GDScript)
- **Composition**: Use Nodes/Components over deep class hierarchies.
- **Typing**: Static Typing (`var health: int`) is MANDATORY.
- **Structure**: Feature-based folders (e.g., `res://actors/player/`).

### Tools and philosophies
- **sqlc** 
## Automated Quality Gates (Strict)

### 1. Localization (i18n) Rules
*Ambiguity removed: Separated UI strings from System strings.*

- **Rule: User_Facing_Text**
    - Context: Argument in UI functions (`Button.text`, `Label.text`, `dialogue`).
    - Validator: **MUST** use `tr("KEY_NAME")`.
    - Fail Condition: String literal detected in UI assignment.

- **Rule: System_Identifiers**
    - Context: Dictionary keys, Signal names, Group names.
    - Validator: **MUST** use `const` variables defined at top of file.
    - Fail Condition: String literal detected in logic flow.

### 2. Visual Consistency Rules
*Ambiguity removed: Separated Editor styling from Code styling.*

- **Rule: Editor_Styling**
    - Context: `.tscn` files (Inspector properties).
    - Validator: **MUST** reference a `Theme` Resource (`.tres`).
    - Fail Condition: Local style override detected in node.

- **Rule: Code_Styling**
    - Context: `.gd` files (Script logic).
    - Validator: **MUST** reference the `Palette` Singleton (e.g., `Palette.danger_red`).
    - Fail Condition: `Color()` constructor usage detected.

### 3. Scene Integrity
- **Rule: Script_Attachment**
    - Context: Root nodes of functional scenes.
    - Validator: **MUST** have a script attached.

### 4. Contract Sync
- **Rule: Protobuf_Freshness**
    - Context: Build Pipeline.
    - Validator: Generated Go/GDScript files **MUST** be newer than `.proto` source files.

## Commit Rules
- Format: `type: subject` (e.g., `feat: login system`).
- Types: `feat`, `fix`, `chore`, `refactor`, `docs`.
- Forbidden: `wip`, `temp`, `...`