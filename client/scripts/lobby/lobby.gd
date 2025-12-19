# File: lobby.gd
# Purpose: Controller for lobby screen shown to users with a faction. Displays
# current ship in slot 1 and 4 empty slots for additional ships. Ship display
# shows type, slot number, and faction-styled background. Includes dev reset
# button to restart onboarding flow for testing.
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


func _ready() -> void:
	API.profile_loaded.connect(_on_profile_loaded)
	API.profile_failed.connect(_on_profile_failed)
	API.reset_done.connect(_on_reset_done)
	reset_button.pressed.connect(_on_reset_pressed)
	reset_button.visible = Config.DEV_MODE
	API.get_profile()
	status_label.text = "Loading..."


func _on_profile_loaded(profile: Dictionary) -> void:
	var faction: String = profile.get("faction", "")
	faction_label.text = "Faction: " + faction.capitalize()
	
	var color: Color = FACTION_COLORS.get(faction, Color.WHITE)
	faction_label.add_theme_color_override("font_color", color)
	
	_display_ships(profile.get("ships", []))
	status_label.text = ""


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
		var slot: Panel = _create_ship_slot(i + 1, ships)
		ship_container.add_child(slot)


func _create_ship_slot(slot_num: int, ships: Array) -> Panel:
	var panel: Panel = Panel.new()
	panel.custom_minimum_size = Vector2(120, 150)
	
	var label: Label = Label.new()
	label.horizontal_alignment = HORIZONTAL_ALIGNMENT_CENTER
	label.vertical_alignment = VERTICAL_ALIGNMENT_CENTER
	label.anchors_preset = Control.PRESET_FULL_RECT
	
	var ship_data: Dictionary = _find_ship_in_slot(slot_num, ships)
	if ship_data.is_empty():
		label.text = "Slot " + str(slot_num) + "\n[Empty]"
	else:
		label.text = "Slot " + str(slot_num) + "\n" + ship_data.get("ship_type", "")
	
	panel.add_child(label)
	return panel


func _find_ship_in_slot(slot: int, ships: Array) -> Dictionary:
	for s in ships:
		if s.get("slot", 0) == slot:
			return s
	return {}
