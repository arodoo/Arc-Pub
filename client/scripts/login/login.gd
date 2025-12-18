# File: login.gd
# Purpose: Controller script for login scene UI. Handles user input from
# email and password fields, validates non-empty input, triggers API login
# request, and displays success/error feedback. Connects to API signals for
# async response handling. Entry point for user authentication flow.
# Path: client/scripts/login/login.gd
# All Rights Reserved. Arc-Pub.

extends Control

@onready var email_input: LineEdit = $VBox/EmailInput
@onready var password_input: LineEdit = $VBox/PasswordInput
@onready var login_button: Button = $VBox/LoginButton
@onready var status_label: Label = $VBox/StatusLabel


func _ready() -> void:
	login_button.pressed.connect(_on_login_pressed)
	API.login_success.connect(_on_login_success)
	API.login_failed.connect(_on_login_failed)
	
	# Dev defaults
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
	status_label.text = "Login successful!"
	login_button.disabled = false
	print("Access Token: ", tokens.get("access_token", ""))
	print("Expires in: ", tokens.get("expires_in", 0), " seconds")


func _on_login_failed(error: String) -> void:
	status_label.text = "Error: " + error
	login_button.disabled = false
