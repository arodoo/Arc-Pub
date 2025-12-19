# File: server_select.gd
# Purpose: Controller for server selection screen shown to new users after
# login but before faction selection. Fetches available servers from backend
# and displays them as buttons. On selection, assigns server to user and
# transitions to faction_select. Server selection is one-time only.
# Path: client/scripts/server_select/server_select.gd
# All Rights Reserved. Arc-Pub.

extends Control

@onready var server_list: VBoxContainer = $VBox/ServerList
@onready var status_label: Label = $VBox/StatusLabel


func _ready() -> void:
	API.servers_loaded.connect(_on_servers_loaded)
	API.servers_failed.connect(_on_servers_failed)
	API.server_selected.connect(_on_server_selected)
	API.server_failed.connect(_on_server_failed)
	
	status_label.text = "Loading servers..."
	API.get_servers()


func _on_servers_loaded(servers: Array) -> void:
	status_label.text = ""
	_display_servers(servers)


func _on_servers_failed(error: String) -> void:
	status_label.text = "Error: " + error


func _on_server_selected(_result: Dictionary) -> void:
	status_label.text = "Server selected!"
	await get_tree().create_timer(0.5).timeout
	get_tree().change_scene_to_file("res://scenes/faction_select/faction_select.tscn")


func _on_server_failed(error: String) -> void:
	status_label.text = "Error: " + error
	_enable_buttons(true)


func _display_servers(servers: Array) -> void:
	for child in server_list.get_children():
		child.queue_free()
	
	for server in servers:
		var btn: Button = Button.new()
		btn.text = server.get("name", "") + " (" + server.get("region", "") + ")"
		btn.pressed.connect(_on_server_pressed.bind(server.get("id", "")))
		server_list.add_child(btn)


func _on_server_pressed(server_id: String) -> void:
	status_label.text = "Connecting..."
	_enable_buttons(false)
	API.select_server(server_id)


func _enable_buttons(enabled: bool) -> void:
	for child in server_list.get_children():
		if child is Button:
			child.disabled = not enabled
