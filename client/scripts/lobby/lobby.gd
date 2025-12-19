# File: lobby.gd
# Purpose: Controller for lobby screen. Displays ships as clickable slots.
# Click on occupied slot → go to game. Supports faction colors and dev reset.
# Stores selected ship in GameData for game scene to read.
# Path: client/scripts/lobby/lobby.gd
# All Rights Reserved. Arc-Pub.

extends Control

@onready var faction_label: Label = $VBox/FactionLabel
@onready var ship_container: HBoxContainer = $VBox/ShipContainer
@onready var status_label: Label = $VBox/StatusLabel
@onready var reset_button: Button = $VBox/ResetButton

const FACTION_COLORS: Dictionary = {
	"red": Color(0.8, 0.2, 0.2),
	"blue": Color(0.2, 0.4, 0.8),
	"green": Color(0.2, 0.7, 0.3)
}

var _current_faction: String = ""


func _ready() -> void:
	API.profile_loaded.connect(_on_profile_loaded)
	API.profile_failed.connect(_on_profile_failed)
	API.reset_done.connect(_on_reset_done)
	reset_button.pressed.connect(_on_reset_pressed)
	reset_button.visible = Config.DEV_MODE
	API.get_profile()
	status_label.text = "Loading..."


func _on_profile_loaded(profile: Dictionary) -> void:
	_current_faction = profile.get("faction", "")
	faction_label.text = "Faction: " + _current_faction.capitalize()
	
	var color: Color = FACTION_COLORS.get(_current_faction, Color.WHITE)
	faction_label.add_theme_color_override("font_color", color)
	
	_display_ships(profile.get("ships", []))
	status_label.text = "Click a ship to enter game"


func _on_profile_failed(error: String) -> void:
	status_label.text = "Error: " + error


func _on_reset_pressed() -> void:
	status_label.text = "Resetting progress..."
	reset_button.disabled = true
	API.reset_progress()


func _on_reset_done() -> void:
	get_tree().change_scene_to_file("res://scenes/login/login.tscn")


func _display_ships(ships: Array) -> void:
	for child in ship_container.get_children():
		child.queue_free()
	
	for i in range(5):
		var slot: Button = _create_ship_slot(i + 1, ships)
		ship_container.add_child(slot)


func _create_ship_slot(slot_num: int, ships: Array) -> Button:
	var btn: Button = Button.new()
	btn.custom_minimum_size = Vector2(120, 150)
	
	var ship_data: Dictionary = _find_ship_in_slot(slot_num, ships)
	if ship_data.is_empty():
		btn.text = "Slot " + str(slot_num) + "\n[Empty]"
		btn.disabled = true
	else:
		btn.text = "Slot " + str(slot_num) + "\n" + ship_data.get("ship_type", "")
		btn.pressed.connect(_on_ship_clicked.bind(ship_data))
	
	return btn


func _on_ship_clicked(ship_data: Dictionary) -> void:
	GameData.set_ship(ship_data, _current_faction)
	get_tree().change_scene_to_file("res://scenes/game/game.tscn")


func _find_ship_in_slot(slot: int, ships: Array) -> Dictionary:
	for s in ships:
		if s.get("slot", 0) == slot:
			return s
	return {}
