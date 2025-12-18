# File: api.gd
# Purpose: Autoload singleton providing HTTP client for backend communication.
# Handles REST API calls with JSON serialization, token storage, and signals
# for async response handling. Centralizes all backend interactions for
# maintainability. Used by all scenes needing server communication.
# Path: client/scripts/autoload/api.gd
# All Rights Reserved. Arc-Pub.

extends Node

## Emitted on successful login with token data.
signal login_success(tokens: Dictionary)
## Emitted on login failure with error message.
signal login_failed(error: String)

const BASE_URL: String = "http://localhost:8080/api/v1"

var _http: HTTPRequest
var _access_token: String = ""


func _ready() -> void:
	_http = HTTPRequest.new()
	add_child(_http)
	_http.request_completed.connect(_on_request_completed)


## Initiates login request with email and password.
func login(email: String, password: String) -> void:
	var body: Dictionary = {"email": email, "password": password}
	var json: String = JSON.stringify(body)
	var headers: PackedStringArray = ["Content-Type: application/json"]
	
	var error: int = _http.request(
		BASE_URL + "/auth/login",
		headers,
		HTTPClient.METHOD_POST,
		json
	)
	
	if error != OK:
		login_failed.emit("Request failed: " + str(error))


func _on_request_completed(
	result: int,
	response_code: int,
	_headers: PackedStringArray,
	body: PackedByteArray
) -> void:
	if result != HTTPRequest.RESULT_SUCCESS:
		login_failed.emit("Connection error")
		return
	
	var json: Variant = JSON.parse_string(body.get_string_from_utf8())
	
	if response_code == 200 and json is Dictionary:
		_access_token = json.get("access_token", "")
		login_success.emit(json)
	else:
		login_failed.emit("Invalid credentials")
