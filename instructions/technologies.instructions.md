[PROJECT]
Name = Metaverso 2D
Type = MMO Social
License_Model = 100% Royalty Free

[CLIENT_SIDE]
Engine = Godot Engine 4 Standard
Language = GDScript
Build_Target = WebAssembly
Export_Tool = Emscripten
License = MIT

[SERVER_SIDE]
Language = Go (Golang)
Architecture = Authoritative Server
Protocol = WebTransport
Fallback_Protocol = WebSockets
Data_Format = Protocol Buffers
License = BSD-3-Clause

[DATABASE]
Persistence = PostgreSQL
Cache = Valkey
Model = Relational + Key/Value
License = PostgreSQL License / BSD

[DEV_ENVIRONMENT]
Editor = VS Code
Extension = godot-tools
License = MIT / GPL