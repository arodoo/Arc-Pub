# File: faction_select.gd
# Purpose: Controller for faction selection screen shown to new users after
# login. Displays three faction buttons (red, blue, green) each with unique
# visual style. On selection, calls API to set faction and create initial
# Betha 1 ship. Transitions to lobby scene on success.
# Path: client/scripts/faction_select/faction_select.gd
# All Rights Reserved. Arc-Pub.

extends Control

@onready var red_btn: Button = $VBox/RedButton
@onready var blue_btn: Button = $VBox/BlueButton
@onready var green_btn: Button = $VBox/GreenButton
@onready var status_label: Label = $VBox/StatusLabel


func _ready() -> void:
	red_btn.pressed.connect(_on_faction_pressed.bind("red"))
	blue_btn.pressed.connect(_on_faction_pressed.bind("blue"))
	green_btn.pressed.connect(_on_faction_pressed.bind("green"))
	API.faction_selected.connect(_on_faction_selected)
	API.faction_failed.connect(_on_faction_failed)


func _on_faction_pressed(faction: String) -> void:
	status_label.text = "Selecting faction..."
	_set_buttons_disabled(true)
	API.select_faction(faction)


func _on_faction_selected(_result: Dictionary) -> void:
	status_label.text = "Welcome to your faction!"
	await get_tree().create_timer(1.0).timeout
	get_tree().change_scene_to_file("res://scenes/lobby/lobby.tscn")


func _on_faction_failed(error: String) -> void:
	status_label.text = "Error: " + error
	_set_buttons_disabled(false)


func _set_buttons_disabled(disabled: bool) -> void:
	red_btn.disabled = disabled
	blue_btn.disabled = disabled
	green_btn.disabled = disabled
