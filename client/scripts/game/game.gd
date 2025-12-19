# File: game.gd
# Purpose: Main game controller for space view. Displays map ID based on
# faction (red→1.1, blue→2.1, green→3.1). Spawns ship with faction color.
# Handles camera zoom with scroll wheel. Uses component architecture.
# Path: client/scripts/game/game.gd
# All Rights Reserved. Arc-Pub.

extends Node2D

@onready var camera: Camera2D = $Camera2D
@onready var map_label: Label = $UI/MapLabel

const ZOOM_LEVELS: Array[float] = [0.5, 1.0, 2.0]
const SHIP_SCENE: PackedScene = preload("res://scenes/components/ship.tscn")

var _zoom_index: int = 1
var _ship: Node


func _ready() -> void:
	_show_map_id()
	_spawn_ship()
	_apply_zoom()


func _input(event: InputEvent) -> void:
	if event is InputEventMouseButton:
		if event.button_index == MOUSE_BUTTON_WHEEL_UP and event.pressed:
			_zoom_in()
		elif event.button_index == MOUSE_BUTTON_WHEEL_DOWN and event.pressed:
			_zoom_out()


func _show_map_id() -> void:
	var map_id: String = GameData.get_map_id()
	map_label.text = "Map: " + map_id


func _spawn_ship() -> void:
	_ship = SHIP_SCENE.instantiate()
	var ship_type: String = GameData.current_ship.get("ship_type", "betha_1")
	var faction: String = GameData.faction
	_ship.setup(ship_type, faction)
	add_child(_ship)


func _apply_zoom() -> void:
	var level: float = ZOOM_LEVELS[_zoom_index]
	camera.zoom = Vector2(level, level)


func _zoom_in() -> void:
	if _zoom_index < ZOOM_LEVELS.size() - 1:
		_zoom_index += 1
		_apply_zoom()


func _zoom_out() -> void:
	if _zoom_index > 0:
		_zoom_index -= 1
		_apply_zoom()
