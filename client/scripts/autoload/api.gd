# File: api.gd
# Purpose: Autoload singleton providing HTTP client for backend communication.
# Handles REST API calls with JSON serialization, token storage, and signals
# for async response handling. Centralizes all backend interactions for
# maintainability. Includes login, profile, and faction selection methods.
# Path: client/scripts/autoload/api.gd
# All Rights Reserved. Arc-Pub.

extends Node

signal login_success(tokens: Dictionary)
signal login_failed(error: String)
signal profile_loaded(profile: Dictionary)
signal profile_failed(error: String)
signal faction_selected(result: Dictionary)
signal faction_failed(error: String)

const BASE_URL: String = "http://localhost:8080/api/v1"

var _http: HTTPRequest
var _access_token: String = ""
var _user_id: String = ""
var _pending_action: String = ""


func _ready() -> void:
	_http = HTTPRequest.new()
	add_child(_http)
	_http.request_completed.connect(_on_request_completed)


func login(email: String, password: String) -> void:
	_pending_action = "login"
	_post("/auth/login", {"email": email, "password": password})


func get_profile() -> void:
	_pending_action = "profile"
	_http_get("/user/profile")


func select_faction(faction: String) -> void:
	_pending_action = "faction"
	_post("/user/faction", {"faction": faction})


func _http_get(endpoint: String) -> void:
	var headers: PackedStringArray = _get_headers()
	_http.request(BASE_URL + endpoint, headers, HTTPClient.METHOD_GET)


func _post(endpoint: String, body: Dictionary) -> void:
	var headers: PackedStringArray = _get_headers()
	var json: String = JSON.stringify(body)
	_http.request(BASE_URL + endpoint, headers, HTTPClient.METHOD_POST, json)


func _get_headers() -> PackedStringArray:
	var headers: PackedStringArray = ["Content-Type: application/json"]
	if _user_id != "":
		headers.append("X-User-ID: " + _user_id)
	return headers


func _on_request_completed(
	result: int,
	response_code: int,
	_headers: PackedStringArray,
	body: PackedByteArray
) -> void:
	if result != HTTPRequest.RESULT_SUCCESS:
		_emit_error("Connection error")
		return
	
	var json: Variant = JSON.parse_string(body.get_string_from_utf8())
	
	match _pending_action:
		"login":
			_handle_login(response_code, json)
		"profile":
			_handle_profile(response_code, json)
		"faction":
			_handle_faction(response_code, json)


func _handle_login(code: int, data: Variant) -> void:
	if code == 200 and data is Dictionary:
		_access_token = data.get("access_token", "")
		login_success.emit(data)
	else:
		login_failed.emit("Invalid credentials")


func _handle_profile(code: int, data: Variant) -> void:
	if code == 200 and data is Dictionary:
		profile_loaded.emit(data)
	else:
		profile_failed.emit("Failed to load profile")


func _handle_faction(code: int, data: Variant) -> void:
	if code == 200 and data is Dictionary:
		faction_selected.emit(data)
	else:
		faction_failed.emit("Failed to select faction")


func _emit_error(msg: String) -> void:
	match _pending_action:
		"login": login_failed.emit(msg)
		"profile": profile_failed.emit(msg)
		"faction": faction_failed.emit(msg)


func set_user_id(id: String) -> void:
	_user_id = id
