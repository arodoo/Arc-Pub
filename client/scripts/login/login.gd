# File: login.gd
# Purpose: Controller script for login scene UI. Handles user input from
# email and password fields, validates non-empty input, triggers API login
# request, and routes to server_select, faction_select or lobby based on
# profile state. Includes dev reset button for testing fresh flow.
# Path: client/scripts/login/login.gd
# All Rights Reserved. Arc-Pub.

extends Control

@onready var email_input: LineEdit = $VBox/EmailInput
@onready var password_input: LineEdit = $VBox/PasswordInput
@onready var login_button: Button = $VBox/LoginButton
@onready var status_label: Label = $VBox/StatusLabel
@onready var reset_button: Button = $VBox/ResetButton


func _ready() -> void:
	login_button.pressed.connect(_on_login_pressed)
	reset_button.pressed.connect(_on_reset_pressed)
	reset_button.visible = Config.DEV_MODE
	API.login_success.connect(_on_login_success)
	API.login_failed.connect(_on_login_failed)
	API.profile_loaded.connect(_on_profile_loaded)
	API.profile_failed.connect(_on_profile_failed)
	API.reset_done.connect(_on_reset_done)
	
	# Dev defaults (only in dev mode)
	if Config.DEV_MODE:
		email_input.text = "admin@dev.local"
		password_input.text = "admin123"


func _on_login_pressed() -> void:
	var email: String = email_input.text.strip_edges()
	var password: String = password_input.text
	
	if email.is_empty() or password.is_empty():
		status_label.text = "Please fill all fields"
		return
	
	status_label.text = "Logging in..."
	login_button.disabled = true
	API.login(email, password)


func _on_login_success(tokens: Dictionary) -> void:
	status_label.text = "Loading profile..."
	API.set_user_id(_get_user_id_from_token(tokens.get("access_token", "")))
	API.get_profile()


func _on_login_failed(error: String) -> void:
	status_label.text = "Error: " + error
	login_button.disabled = false


func _on_profile_loaded(profile: Dictionary) -> void:
	var server_id: Variant = profile.get("server_id")
	var faction: Variant = profile.get("faction")
	
	if server_id == null or server_id == "":
		get_tree().change_scene_to_file("res://scenes/server_select/server_select.tscn")
	elif faction == null or faction == "":
		get_tree().change_scene_to_file("res://scenes/faction_select/faction_select.tscn")
	else:
		get_tree().change_scene_to_file("res://scenes/lobby/lobby.tscn")


func _on_profile_failed(_error: String) -> void:
	get_tree().change_scene_to_file("res://scenes/server_select/server_select.tscn")


func _on_reset_pressed() -> void:
	status_label.text = "Resetting... (login first)"
	API.reset_progress()


func _on_reset_done() -> void:
	status_label.text = "Progress reset! Login again."
	login_button.disabled = false


func _get_user_id_from_token(token: String) -> String:
	var parts: PackedStringArray = token.split(".")
	if parts.size() < 2:
		return ""
	var payload: String = parts[1]
	while payload.length() % 4 != 0:
		payload += "="
	var decoded: PackedByteArray = Marshalls.base64_to_raw(payload)
	var json: Variant = JSON.parse_string(decoded.get_string_from_utf8())
	if json is Dictionary:
		return json.get("sub", "")
	return ""
